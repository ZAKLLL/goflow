package defineflows

import (
	"github.com/s8sg/goflow/coderunner"
	flow "github.com/s8sg/goflow/flow/v1"
)

// Define provide definition of the workflow
func WorkFlow_codeRunner(workflow *flow.Workflow, context *flow.Context) error {
	dag := workflow.Dag()

	src :=
		`package Solution
		import (
			"fmt"
			"io/ioutil"
			"time")

		func Run(intputPath, outputPath string) {
			data, _ := ioutil.ReadFile(intputPath)
			fmt.Printf("Hi this codeRunner! input Data :%s \n", string(data))
			tStr := time.Now().String()
			ioutil.WriteFile(outputPath, []byte("Hello this is CodeRunner ! "+tStr), 0644)
		}
		`

	codeRunner := &coderunner.CodeRunner{
		WorkSpace:  "C:\\Users\\张家魁\\Desktop\\codeRunner",
		CodeType:   coderunner.GOLANG,
		IsAsync:    true,
		SourceCode: src,
	}
	dag.NodeWithCode("test", codeRunner)
	return nil
}

// 执行注册操作
func init() {
	registerNewFlow("mycode", WorkFlow_codeRunner)
}
