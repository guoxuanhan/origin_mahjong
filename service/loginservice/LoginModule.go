package loginservice

import (
	"github.com/duanhf2012/origin/v2/log"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"origin_mahjong/common/proto/common"
	"origin_mahjong/common/proto/outer"
	"time"

	"github.com/duanhf2012/origin/v2/service"
	"github.com/duanhf2012/origin/v2/sysmodule/netmodule/ginmodule"
)

type GateStatus int

const (
	Offline GateStatus = 0
	Online  GateStatus = 1
)

const (
	MaxLoginCD = 3 * time.Second // 登录的冷却时间
)

// LoginModule 登录模块
type LoginModule struct {
	service.Module
	mapLoginCD    map[string]int64 // 账号登录的冷却时间
	lastResetTime int64            // 上一次重置冷却时间的时间 避免map一直增大
}

func (login *LoginModule) OnInit() error {
	login.mapLoginCD = make(map[string]int64, 1024)
	login.lastResetTime = 0
	return nil
}

func (login *LoginModule) OnRelease() {

}

// Login 客户端登录逻辑
func (login *LoginModule) Login(c *ginmodule.SafeContext) {
	var loginRequest outer.C2L_LoginRequest
	loginResponse := outer.L2C_LoginResponse{}
	err := c.ShouldBindBodyWith(&loginRequest, binding.JSON)
	if err != nil {
		loginResponse.Error = int32(common.ErrorCode_ERR_AccountOrPassword)
		c.JSONAndDone(http.StatusBadRequest, loginResponse)
		return
	}

	log.SDebug("Login Request: ", int32(loginRequest.LoginType), "Account:", loginRequest.Account)

	if loginRequest.LoginType < outer.LoginType_Guest || loginRequest.LoginType >= outer.LoginType_Max {
		loginResponse.Error = int32(common.ErrorCode_ERR_AccountOrPassword)
		c.JSONAndDone(http.StatusBadRequest, loginResponse)
		return
	}

	if len(loginRequest.Account) == 0 {
		loginResponse.Error = int32(common.ErrorCode_ERR_AccountOrPassword)
		c.JSONAndDone(http.StatusBadRequest, loginResponse)
		return
	}
}
