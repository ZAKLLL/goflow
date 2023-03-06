package flowregistry

import (
	"fmt"

	flow "github.com/s8sg/goflow/flow/v1"
)

// Workload function
func f1(data []byte, option map[string][]string) ([]byte, error) {
	return []byte(fmt.Sprintf("f1 you said \"%s\"", string(data))), nil
}

// Workload function
func f2(data []byte, option map[string][]string) ([]byte, error) {
	return []byte(fmt.Sprintf("f2 you said \"%s\"", string(data))), nil
}

// Define provide definition of the workflow
func MyWorkFlow(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	dag.Node("test", f1)
	dag.Node("test", f2)
	return nil
}

// 执行注册操作
func init() {
	registerNewFlow("myFlow", MyWorkFlow)
}
