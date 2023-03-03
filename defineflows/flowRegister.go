package defineflows

import (
	"log"

	"github.com/s8sg/goflow/runtime"
	goflow "github.com/s8sg/goflow/v1"
)

var registryFlowMap = make(map[string]runtime.FlowDefinitionHandler)

func DoRegister(fs *goflow.FlowService) {
	for k, v := range registryFlowMap {
		fs.Register(k, v)
		log.Printf("registry flow :[%s] successfully \n", k)
	}
}

func registerNewFlow(flowName string, flowDefine runtime.FlowDefinitionHandler) {
	if _, ok := registryFlowMap[flowName]; ok {
		log.Panicf("exist duplicated flowName %s \n", flowName)
	}

	registryFlowMap[flowName] = flowDefine
}
