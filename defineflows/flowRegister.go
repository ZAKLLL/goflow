package defineflows

import (
	"log"

	"github.com/s8sg/goflow/runtime"
	goflow "github.com/s8sg/goflow/v1"
)

type FlowRunTimeRegistryInfo struct {
	FlowName    string `json:"FlowName"`
	Description string `json:"Description"`
	Dag         struct {
		Nodes []struct {
			NodeName    string `json:"NodeName"`
			NodeCodeTag string `json:"NodeCodeTag"`
		} `json:"Nodes"`
		Edges []struct {
			From string `json:"From"`
			To   string `json:"To"`
		} `json:"Edges"`
	} `json:"Dag"`
	SourceCode []struct {
		CodeTag   string `json:"CodeTag"`
		CodeSrc   string `json:"CodeSrc"`
		CodeType  int    `json:"CodeType"`
		WorkSpace string `json:"WorkSpace"`
	} `json:"SourceCode"`
}

// 预定义的flow,通过init()函数调用registerNewFlow 进行注册
var registryFlowMap = make(map[string]runtime.FlowDefinitionHandler)

func DoRegister(fs *goflow.FlowService) error {
	for k, v := range registryFlowMap {
		err := fs.Register(k, v)
		if err != nil {
			return err
		}
		log.Printf("registry flow :[%s] successfully \n", k)
	}
	return nil
}

func registerNewFlow(flowName string, flowDefine runtime.FlowDefinitionHandler) {
	if _, ok := registryFlowMap[flowName]; ok {
		log.Panicf("exist duplicated flowName %s \n", flowName)
	}
	registryFlowMap[flowName] = flowDefine
}

//// 运行时提交dag代码
//func registerNewFlow(flowName string, flowDefine runtime.FlowDefinitionHandler) {
//
//}
//
//func registerNewFlowRuntime(flowJson string) {
//
//}
