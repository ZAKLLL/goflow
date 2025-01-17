package operation

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/s8sg/goflow/coderunner"
)

// FuncErrorHandler the error handler for OnFailure() options
type FuncErrorHandler func(error) error

// Modifier definition for Modify() call
type Modifier func([]byte, map[string][]string) ([]byte, error)

type GoFlowOperation struct {
	Id      string              // ID
	Mod     Modifier            // Modifier
	Options map[string][]string // The option as a input to workload

	FailureHandler FuncErrorHandler // The Failure handler of the operation

	IsCodeExec bool

	CodeRunner *coderunner.CodeRunner
}

func (operation *GoFlowOperation) addOptions(key string, value string) {
	array, ok := operation.Options[key]
	if !ok {
		operation.Options[key] = make([]string, 1)
		operation.Options[key][0] = value
	} else {
		operation.Options[key] = append(array, value)
	}
}

func (operation *GoFlowOperation) AddFailureHandler(handler FuncErrorHandler) {
	operation.FailureHandler = handler
}

func (operation *GoFlowOperation) GetOptions() map[string][]string {
	return operation.Options
}

func (operation *GoFlowOperation) GetId() string {
	return operation.Id
}

func (operation *GoFlowOperation) Encode() []byte {
	return []byte("")
}

// executeWorkload executes a function call
func executeWorkload(operation *GoFlowOperation, data []byte) ([]byte, error) {
	var err error
	var result []byte

	if !operation.IsCodeExec {
		options := operation.GetOptions()
		result, err = operation.Mod(data, options)
	} else {
		if f, _ := hasDir(operation.CodeRunner.WorkSpace); !f {
			_ = os.Mkdir(operation.CodeRunner.WorkSpace, os.ModeDir)
		}

		// 将data 写入 coderunner.InputFileName
		err = ioutil.WriteFile(operation.CodeRunner.GetInputPath(), data, 0644)
		if err != nil {
			return result, err
		}
		outPutPath := operation.CodeRunner.GetOutputPath()
		//创建输出数据文件
		_, err = os.Create(outPutPath)
		if err != nil {
			return result, err
		}
		//执行代码逻辑
		err = operation.CodeRunner.Exec()
		if err != nil {
			return result, err
		}

		//将结果写入到 coderunner.outputFileName
		result, err = ioutil.ReadFile(outPutPath)

	}
	return result, err
}

func (operation *GoFlowOperation) Execute(data []byte, _ map[string]interface{}) ([]byte, error) {
	var result []byte
	var err error

	if operation.Mod != nil || (operation.IsCodeExec && operation.CodeRunner != nil) {
		result, err = executeWorkload(operation, data)
		if err != nil {
			err = fmt.Errorf("function(%s), error: function execution failed, %v",
				operation.Id, err)
			if operation.FailureHandler != nil {
				err = operation.FailureHandler(err)
			}
			if err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

func (operation *GoFlowOperation) GetProperties() map[string][]string {

	result := make(map[string][]string)

	isMod := "false"
	isFunction := "false"
	isHttpRequest := "false"
	hasFailureHandler := "false"

	if operation.Mod != nil {
		isFunction = "true"
	}
	if operation.FailureHandler != nil {
		hasFailureHandler = "true"
	}

	result["isMod"] = []string{isMod}
	result["isFunction"] = []string{isFunction}
	result["isHttpRequest"] = []string{isHttpRequest}
	result["hasFailureHandler"] = []string{hasFailureHandler}

	return result
}
func hasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}
