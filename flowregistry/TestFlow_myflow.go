package flowregistry

import (
	"fmt"

	flow "github.com/s8sg/goflow/flow/v1"
)

// Workload function
func f1(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Printf("f1 you said \"%s\" \n", string(data))
	return []byte(fmt.Sprintf("f1 you said \"%s\"", string(data))), nil
}

// Workload function
func f2(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Printf("f2 you said \"%s\" \n", string(data))
	return []byte(fmt.Sprintf("f2 you said \"%s\"", string(data))), nil
}

// Workload function
func f3(data []byte, option map[string][]string) ([]byte, error) {
	fmt.Printf("f3 you said \"%s\" \n", string(data))
	return []byte(fmt.Sprintf("f3 you said \"%s\"", string(data))), nil
}

// Define provide definition of the workflow
func MyWorkFlow(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()
	dag.Node("n1", f1)
	dag.Node("n2", f2)
	dag.Node("n3", f3)
	dag.Edge("n1", "n2")
	dag.Edge("n2", "n3")
	return nil
}

// 执行注册操作
func init() {
	registerNewFlow("myflow", MyWorkFlow)
}
