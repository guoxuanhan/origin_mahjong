package dbservice

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"origin_mahjong/common/proto/inner"
	"runtime"
	"time"

	"github.com/duanhf2012/origin/v2/log"
	"github.com/duanhf2012/origin/v2/node"
	"github.com/duanhf2012/origin/v2/rpc"
	"github.com/duanhf2012/origin/v2/service"
	"github.com/duanhf2012/origin/v2/sysmodule/mongodbmodule"
	"github.com/duanhf2012/origin/v2/sysmodule/redismodule"
	"github.com/duanhf2012/origin/v2/util/typ"
)

func init() {
	node.Setup(&DBService{})
}

var emptyResult [][]byte

// DBRequest DB请求
type DBRequest struct {
	dbRequest    *inner.DBCtrlRequest
	redisRequest *inner.DBRedisCtrlRequest
	responder    rpc.Responder
}

// DBService DB服务
type DBService struct {
	service.Service
	mongoModule      mongodbmodule.MongoModule
	redisModule      redismodule.RedisModule
	channelDBRequest []chan DBRequest
	url              string
	dbName           string
	goroutineNum     uint32
	channelNum       uint32
}

// OnInit 初始化
func (dbs *DBService) OnInit() error {
	err := dbs.ReadConfig()
	if err != nil {
		return err
	}

	// 初始化mongodb模块
	err = dbs.mongoModule.Init(dbs.url, time.Second*15)
	if err != nil {
		log.SError("Init DBService[", dbs.dbName, "], url[", dbs.url, "] init error:", err.Error())
		return err
	}

	// 启动mongodb模块
	err = dbs.mongoModule.Start()
	if err != nil {
		log.SError("Start DBService[", dbs.dbName, "], url[", dbs.url, "] init error:", err.Error())
		return err
	}

	// 初始化redis服务
	redisConfig := &redismodule.ConfigRedis{
		IP:          "127.0.0.1",
		Port:        6379,
		Password:    "123456",
		DbIndex:     0,
		MaxIdle:     3,
		MaxActive:   10,
		IdleTimeout: 30,
	}

	dbs.redisModule.Init(redisConfig)

	// 启动并发协程
	dbs.channelDBRequest = make([]chan DBRequest, dbs.goroutineNum)
	for i := uint32(0); i < dbs.goroutineNum; i++ {
		dbs.channelDBRequest[i] = make(chan DBRequest, dbs.channelNum)
		go dbs.execute(dbs.channelDBRequest[i])
	}

	// 性能监控
	dbs.OpenProfiler()
	dbs.GetProfiler().SetOverTime(time.Millisecond * 500)
	dbs.GetProfiler().SetMaxOverTime(time.Second * 10)

	return nil
}

// ReadConfig 读取配置
func (dbs *DBService) ReadConfig() error {
	mapDBServerConfig, ok := dbs.GetServiceCfg().(map[string]interface{})
	if !ok {
		return fmt.Errorf("DBService config is error")
	}

	// 解析url地址
	url, ok := mapDBServerConfig["url"]
	if !ok {
		return fmt.Errorf("DBService config is error")
	}
	dbs.url = url.(string)

	// 解析数据库名称
	dbName, ok := mapDBServerConfig["dbName"]
	if !ok {
		return fmt.Errorf("DBService config is error")
	}
	dbs.dbName = dbName.(string)

	// 解析并发数量
	goroutineNum, ok := mapDBServerConfig["GoroutineNum"]
	if !ok {
		return fmt.Errorf("DBService config is error")
	}
	dbs.goroutineNum, _ = typ.ConvertToNumber[uint32](goroutineNum)

	// 解析channel容量
	channelNum, ok := mapDBServerConfig["ChannelNum"]
	if !ok {
		return fmt.Errorf("DBService config is error")
	}
	dbs.channelNum, _ = typ.ConvertToNumber[uint32](channelNum)

	return nil
}

// execute 并发执行DB操作
func (dbs *DBService) execute(dbChannel chan DBRequest) {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			l := runtime.Stack(buf, false)
			err := fmt.Errorf("%v: %s", r, buf[:l])
			log.SError("core dump info:", err.Error(), "\n")
			dbs.execute(dbChannel)
		}
	}()

	for {
		select {
		case dbData := <-dbChannel:
			switch dbData.dbRequest.GetOPType() {
			case inner.DBOperationType_Delete:
				err := dbs.doDelete(dbData)
				if err != nil {
					log.SError("DBOperationType_Delete doDelete error: ", err.Error())
				}
			case inner.DBOperationType_DeleteMany:
				err := dbs.doDeleteMany(dbData)
				if err != nil {
					log.SError("DBOperationType_Delete doDeleteMany error: ", err.Error())
				}
			case inner.DBOperationType_Find:
				err := dbs.doFind(dbData)
				if err != nil {
					log.SError("DBOperationType_Find doFind error: ", err.Error())
				}
			case inner.DBOperationType_FindMany:
				err := dbs.doFindMany(dbData)
				if err != nil {
					log.SError("DBOperationType_FindMany doFindMany error: ", err.Error())
				}
			case inner.DBOperationType_Update:
				err := dbs.doUpdate(dbData)
				if err != nil {
					log.SError("DBOperationType_Update doUpdate error: ", err.Error())
				}
			case inner.DBOperationType_UpdateMany:
				err := dbs.doUpdateMany(dbData)
				if err != nil {
					log.SError("DBOperationType_UpdateMany doUpdateMany error: ", err.Error())
				}
			case inner.DBOperationType_Redis_SetKey:
			case inner.DBOperationType_Redis_GetKey:
			case inner.DBOperationType_Redis_DelKey:
			}
		}
	}

}

// doDelete 执行删除一条数据
func (dbs *DBService) doDelete(dbReq DBRequest) error {
	// 选择数据库与表
	s := dbs.mongoModule.TakeSession()

	// 定义返回数据集
	var dbRes inner.DBCtrlResponse
	dbRes.Result = emptyResult

	// 设置条件
	var condition any
	unmarshalErr := bson.Unmarshal(dbReq.dbRequest.GetCondition(), &condition)
	if unmarshalErr != nil {
		dbReq.responder(&dbRes, rpc.RpcError(unmarshalErr.Error()))
		return unmarshalErr
	}

	ctx, cancel := s.GetDefaultContext()
	defer cancel()
	result, err := s.Collection(dbs.dbName, dbReq.dbRequest.GetCollectName()).DeleteOne(ctx, condition)
	if err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	dbRes.DeletedCount = result.DeletedCount
	dbReq.responder(&dbRes, rpc.NilError)
	return nil
}

// doDeleteMany 执行删除匹配条件的所有数据
func (dbs *DBService) doDeleteMany(dbReq DBRequest) error {
	// 选择数据库与表
	s := dbs.mongoModule.TakeSession()

	// 定义返回数据集
	var dbRes inner.DBCtrlResponse
	dbRes.Result = emptyResult

	// 设置条件
	var condition any
	unmarshalErr := bson.Unmarshal(dbReq.dbRequest.GetCondition(), &condition)
	if unmarshalErr != nil {
		dbReq.responder(&dbRes, rpc.RpcError(unmarshalErr.Error()))
		return unmarshalErr
	}

	ctx, cancel := s.GetDefaultContext()
	defer cancel()
	result, err := s.Collection(dbs.dbName, dbReq.dbRequest.GetCollectName()).DeleteMany(ctx, condition)
	if err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	dbRes.DeletedCount = result.DeletedCount
	dbReq.responder(&dbRes, rpc.NilError)
	return nil
}

// doFind 执行查询一条数据
func (dbs *DBService) doFind(dbReq DBRequest) error {
	// 选择数据库与表
	s := dbs.mongoModule.TakeSession()

	// 定义返回数据集
	var dbRes inner.DBCtrlResponse
	dbRes.Result = emptyResult

	// 设置条件
	var condition any
	unmarshalErr := bson.Unmarshal(dbReq.dbRequest.GetCondition(), &condition)
	if unmarshalErr != nil {
		dbReq.responder(&dbRes, rpc.RpcError(unmarshalErr.Error()))
		return unmarshalErr
	}

	ctx, cancel := s.GetDefaultContext()
	defer cancel()
	result := s.Collection(dbs.dbName, dbReq.dbRequest.GetCollectName()).FindOne(ctx, condition)
	if result.Err() != nil {
		dbReq.responder(&dbRes, rpc.RpcError(result.Err().Error()))
		return result.Err()
	}

	var rawResult bson.M
	if err := result.Decode(&rawResult); err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(result.Err().Error()))
		return err
	}

	// 序列化为 []byte
	data, err := bson.Marshal(rawResult)
	if err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	dbRes.OPType = dbReq.dbRequest.GetOPType()
	dbRes.Result = [][]byte{data}

	dbRes.MatchedCount = 1
	dbReq.responder(&dbRes, rpc.NilError)
	return nil
}

// doFindMany 执行查询匹配的所有数据
func (dbs *DBService) doFindMany(dbReq DBRequest) error {
	// 选择数据库与表
	s := dbs.mongoModule.TakeSession()

	// 定义返回数据集
	var dbRes inner.DBCtrlResponse
	dbRes.Result = emptyResult

	// 设置条件
	var condition any
	unmarshalErr := bson.Unmarshal(dbReq.dbRequest.GetCondition(), &condition)
	if unmarshalErr != nil {
		dbReq.responder(&dbRes, rpc.RpcError(unmarshalErr.Error()))
		return unmarshalErr
	}

	ctx, cancel := s.GetDefaultContext()
	defer cancel()
	cursor, err := s.Collection(dbs.dbName, dbReq.dbRequest.GetCollectName()).Find(ctx, condition)
	if err != nil || cursor.Err() != nil {
		if err == nil {
			err = cursor.Err()
		}
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	var result []interface{}
	ctxAll, cancelAll := s.GetDefaultContext()
	defer cancelAll()
	err = cursor.All(ctxAll, &result)
	if err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	// 获取结果集
	var rpcErr rpc.RpcError
	dbRes.OPType = dbReq.dbRequest.GetOPType()
	dbRes.Result = make([][]byte, len(result))

	// 反序列化结果集
	for i := 0; i < len(result); i++ {
		dbRes.Result[i], err = bson.Marshal(result[i])
		if err != nil {
			rpcErr = rpc.RpcError(err.Error())
			dbRes.Result = emptyResult
			break
		}
	}

	dbRes.MatchedCount = int64(len(result))

	dbReq.responder(&dbRes, rpcErr)
	return nil
}

// doUpdate 执行更新数据，若不存在则插入
func (dbs *DBService) doUpdate(dbReq DBRequest) error {
	// 选择数据库与表
	s := dbs.mongoModule.TakeSession()

	// 定义返回数据集
	var dbRes inner.DBCtrlResponse
	dbRes.Result = emptyResult

	// 设置数据
	if len(dbReq.dbRequest.Data) == 0 {
		err := fmt.Errorf("%s doUpdate data len is error %d.", dbReq.dbRequest.CollectName, len(dbReq.dbRequest.Data))
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	// 设置条件
	var condition any
	unmarshalErr := bson.Unmarshal(dbReq.dbRequest.GetCondition(), &condition)
	if unmarshalErr != nil {
		dbReq.responder(&dbRes, rpc.RpcError(unmarshalErr.Error()))
		return unmarshalErr
	}

	// 序列化更新数据
	var updateData any
	var data any
	err := bson.Unmarshal(dbReq.dbRequest.Data[0], &data)
	if err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}
	updateData = bson.M{"$set": data}

	ctx, cancel := s.GetDefaultContext()
	defer cancel()

	// 如果不存在则插入
	opts := options.Update().SetUpsert(true)
	result, err := s.Collection(dbs.dbName, dbReq.dbRequest.GetCollectName()).UpdateOne(ctx, condition, updateData, opts)
	if err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	dbRes.MatchedCount = result.MatchedCount
	dbRes.ModifiedCount = result.ModifiedCount
	dbRes.UpsertedCount = result.UpsertedCount

	dbReq.responder(&dbRes, rpc.NilError)
	return nil
}

// doUpdateMany 执行更新一批数据，若不存在则插入
func (dbs *DBService) doUpdateMany(dbReq DBRequest) error {
	// 选择数据库与表
	s := dbs.mongoModule.TakeSession()

	// 定义返回数据集
	var dbRes inner.DBCtrlResponse
	dbRes.Result = emptyResult

	// 设置数据
	if len(dbReq.dbRequest.Data) == 0 {
		err := fmt.Errorf("%s doUpdateMany data len is error %d.", dbReq.dbRequest.CollectName, len(dbReq.dbRequest.Data))
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	// 设置条件
	var condition any
	unmarshalErr := bson.Unmarshal(dbReq.dbRequest.GetCondition(), &condition)
	if unmarshalErr != nil {
		dbReq.responder(&dbRes, rpc.RpcError(unmarshalErr.Error()))
		return unmarshalErr
	}

	// 序列化更新数据
	var updateData any
	var data any
	err := bson.Unmarshal(dbReq.dbRequest.Data[0], &data)
	if err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}
	updateData = bson.M{"$set": data}

	ctx, cancel := s.GetDefaultContext()
	defer cancel()

	// 如果不存在则插入
	opts := options.Update().SetUpsert(true)
	result, err := s.Collection(dbs.dbName, dbReq.dbRequest.GetCollectName()).UpdateMany(ctx, condition, updateData, opts)
	if err != nil {
		dbReq.responder(&dbRes, rpc.RpcError(err.Error()))
		return err
	}

	dbRes.MatchedCount = result.MatchedCount
	dbRes.ModifiedCount = result.ModifiedCount
	dbRes.UpsertedCount = result.UpsertedCount

	dbReq.responder(&dbRes, rpc.NilError)
	return nil
}

//// RPC_Query 对外提供RPC方法，接收DB操作数据后写入管道，交给协程执行
//func (dbs *DBService) RPC_Query(responder rpc.Responder) error {
//	dbs <- nil
//	return nil
//}
