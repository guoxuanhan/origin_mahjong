package loginservice

import (
	"github.com/duanhf2012/origin/v2/node"
	"github.com/duanhf2012/origin/v2/service"
	"github.com/duanhf2012/origin/v2/sysmodule/netmodule/ginmodule"
	"github.com/gin-gonic/gin"
	"time"
)

func init() {
	node.Setup(&LoginService{})
}

// LoginService 给客户端提供登录服务
type LoginService struct {
	service.Service
	loginModule *LoginModule
	ginModule   ginmodule.GinModule
}

func (ls *LoginService) OnInit() error {
	// 解析配置文件
	var config struct {
		ListenPort        string
		MinGoroutineNum   int32
		MaxGoroutineNum   int32
		MaxTaskChannelNum int
	}

	err := ls.ParseServiceCfg(&config)
	if err != nil {
		return err
	}

	// 打开并发协程模式
	ls.OpenConcurrent(config.MinGoroutineNum, config.MaxGoroutineNum, config.MaxTaskChannelNum)

	// 添加gin模块
	ls.ginModule.Init(config.ListenPort, time.Second*15, nil)
	ls.ginModule.AppendDataProcessor(ls)
	ls.AddModule(&ls.ginModule)
	ls.ginModule.Start()

	// 添加login模块
	ls.loginModule = &LoginModule{}
	ls.AddModule(ls.loginModule)

	// 启动性能监控
	ls.OpenProfiler()
	ls.GetProfiler().SetOverTime(time.Millisecond * 100)
	ls.GetProfiler().SetMaxOverTime(time.Second * 10)

	// 注册登录接口 请求url：http://127.0.0.1:9000/api/login
	ls.ginModule.SafePOST("/api/login", ls.loginModule.Login)

	return nil
}

func (ls *LoginService) Process(c *gin.Context) (*gin.Context, error) {
	return c, nil
}
