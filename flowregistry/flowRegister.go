package flowregistry

import (
	"fmt"
	log2 "github.com/s8sg/goflow/log"
	"github.com/s8sg/goflow/runtime"
	goflow "github.com/s8sg/goflow/v1"
	"gopkg.in/redis.v5"
	"log"
	"time"
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

func RegisterFlows(fs *goflow.FlowService) error {
	err := registerRuntimeFlows(fs)
	if err != nil {
		return err
	}
	go StartRuntimeRegister(fs)
	return nil
}

func registerRuntimeFlows(fs *goflow.FlowService) error {
	for k, v := range registryFlowMap {
		err := fs.Register(k, v, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func StartRuntimeRegister(fs *goflow.FlowService) {
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fs.RedisURL,
			DB:   0,
		})
		_, err := redisClient.Ping().Result()
		if err != nil {
			panic("redis connection error :" + err.Error())
		}
	}
	for {
		//4s执行一次
		time.Sleep(time.Second * 4)

		keys := redisClient.Keys(runtimeRegistryFlowInitial + "*")
		runtimeregistryflowKeys, err := keys.Result()
		if err != nil {
			continue
		}
		for _, key := range runtimeregistryflowKeys {
			flowName := key[len(runtimeRegistryFlowInitial)+1:]
			if _, ok := fs.Flows[flowName]; !ok {
				flowJson, err := redisClient.Get(key).Result()
				if err != nil {
					continue
				}
				err = doRegisterAtRuntime(fs, flowJson, false)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// todo test
func RegisterAtRuntime(fs *goflow.FlowService, flowJson string) error {
	return doRegisterAtRuntime(fs, flowJson, true)
}

func doRegisterAtRuntime(fs *goflow.FlowService, flowJson string, deliver bool) error {
	flowName, dag, err := ConstructDag(flowJson)
	if err != nil {
		return err
	}
	if fs.Logger == nil {
		fs.Logger = &log2.StdErrLogger{}
	}
	fs.Logger.Log(fmt.Sprintf("runtime registing, flowInfo :%s \n", flowJson))
	//相当于在本地预注册一次
	err = fs.Register(flowName, dag, !deliver)
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
	if redisClient == nil {
		redisClient = redis.NewClient(&redis.Options{
			Addr: fs.RedisURL,
			DB:   0,
		})
	}
	set := redisClient.Set(GenFlowRedisKey(flowName), json, ExpireLastTime)
	_, err := set.Result()
	if err != nil {
		return err
	}
	return nil
}

//// 运行时提交dag代码
//func registerNewFlow(flowName string, flowDefine runtime.FlowDefinitionHandler) {
//
//}
//
//func registerNewFlowRuntime() {
//
//}
