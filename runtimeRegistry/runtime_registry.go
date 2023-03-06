package runtimeRegistry

import (
	"encoding/json"
	"fmt"
	"github.com/s8sg/goflow/coderunner"
	flow "github.com/s8sg/goflow/flow/v1"
	"github.com/s8sg/goflow/runtime"
	"time"
)

const RuntimeRegistryFlowInitial = "goflow-runtime-registry-flow"
const ExpireLastTime = 60 * time.Minute

type RuntimeRegistryDag struct {
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
		IsAsync   bool   `json:"isAsync"`
	} `json:"SourceCode"`
}

func (rDag *RuntimeRegistryDag) getCodeRunners() ([]string, []*coderunner.CodeRunner) {
	codeRunners := make([]*coderunner.CodeRunner, 0)
	nodeNames := make([]string, 0)
	for _, rNode := range rDag.Dag.Nodes {
		cRunner := &coderunner.CodeRunner{}
		codeRunners = append(codeRunners, cRunner)
		nodeNames = append(nodeNames, rNode.NodeName)
		for _, srcCode := range rDag.SourceCode {
			if srcCode.CodeTag == rNode.NodeCodeTag {
				cRunner.CodeType = (coderunner.CodeType)(srcCode.CodeType)
				cRunner.SourceCode = srcCode.CodeSrc
				cRunner.IsAsync = srcCode.IsAsync
				cRunner.WorkSpace = srcCode.WorkSpace
				break
			}
		}
	}
	return nodeNames, nil
}

func ConstructDag(dagJson string) (string, runtime.FlowDefinitionHandler, error) {
	runtimeRegistryDag := &RuntimeRegistryDag{}
	err := json.Unmarshal([]byte(dagJson), runtimeRegistryDag)
	if err != nil {
		return "", nil, err
	}
	return runtimeRegistryDag.FlowName, func(workflow *flow.Workflow, context *flow.Context) error {
		dag := workflow.Dag()
		//构建node
		nodeNames, coderRunners := runtimeRegistryDag.getCodeRunners()
		for i, name := range nodeNames {
			runner := coderRunners[i]
			dag.NodeWithCode(name, runner)
		}
		//构建边关系
		for _, edge := range runtimeRegistryDag.Dag.Edges {
			dag.Edge(edge.From, edge.To)
		}
		return nil
	}, nil
}

func GenFlowRedisKey(flowName string) string {
	return fmt.Sprintf("%s:%s", RuntimeRegistryFlowInitial, flowName)
}
