package flowregistry

import (
	"fmt"
	"github.com/s8sg/goflow/runtimeRegistry"
	goflow "github.com/s8sg/goflow/v1"
	"gopkg.in/redis.v5"
	"log"

	"github.com/s8sg/goflow/runtime"
)

// 预定义的flow,通过init()函数调用registerNewFlow 进行注册
var registryFlowMap = make(map[string]runtime.FlowDefinitionHandler)

var redisClient *redis.Client

// call by init()
func registerNewFlow(flowName string, flowDefine runtime.FlowDefinitionHandler) {
	if _, ok := registryFlowMap[flowName]; ok {
		log.Panicf("exist duplicated flowName %s \n", flowName)
	}
	registryFlowMap[flowName] = flowDefine
}

func RegisterDefineFlows(fs *goflow.FlowService) error {
	for k, v := range registryFlowMap {
		err := fs.Register(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// todo test
func RegisterAtRuntime(fs *goflow.FlowService, flowJson string) error {
	return DoRegisterAtRuntime(fs, flowJson, true)
}

func DoRegisterAtRuntime(fs *goflow.FlowService, flowJson string, deliver bool) error {
	flowName, dag, err := runtimeRegistry.ConstructDag(flowJson)
	if err != nil {
		return err
	}
	fs.Logger.Log(fmt.Sprintf("runtime registing, flowInfo :%s \n", flowJson))
	err = fs.Register(flowName, dag)
	if err != nil {
		return err
	}

	if deliver {
		//提交到各个节点注册dag
		err = distributeDagToOtherService(fs, flowName, flowJson)
		if err != nil {
			return err
		}
	}
	return nil
}

func distributeDagToOtherService(fs *goflow.FlowService, flowName, json string) error {
	if redisClient != nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fs.RedisURL,
			DB:   0,
		})
	}
	set := redisClient.Set(runtimeRegistry.GenFlowRedisKey(flowName), json, runtimeRegistry.ExpireLastTime)
	_, err := set.Result()
	if err != nil {
		return err
	}
	return nil
}

func Consume(Fs *goflow.FlowService, flowJson string) {

	err := DoRegisterAtRuntime(Fs, flowJson, false)
	if err != nil {
		Fs.Logger.Log("[goflow RuntimeRegister] failed to register new Flow, error " + err.Error())
		return
	}

}

//// 运行时提交dag代码
//func registerNewFlow(flowName string, flowDefine runtime.FlowDefinitionHandler) {
//
//}
//
//func registerNewFlowRuntime() {
//
//}
