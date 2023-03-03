package coderunner

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
)

// 代码执行输入输出目录
const InputFileName = "__goFlowInput"
const OutputFileName = "__goFlowOutput"

type CodeType int

const (
	LINUX_SHELL CodeType = iota
	POWER_SHELL
	PYTHON
	GOLANG
)

type CodeRunner struct {
	CodeType   CodeType
	WorkSpace  string
	IsAsync    bool
	SourceCode string
}

type CodeExecute interface {
	ExecCoderRunner(CodeRunner *CodeRunner) error
}

// 代码执行器
var CodeExecutorMap = make(map[CodeType]CodeExecute)

// 获取当前codeRunner 数据输入文件位置
func (code *CodeRunner) GetInputPath() string {
	return filepath.Join(code.WorkSpace, InputFileName)
}

// 获取当前codeRunner 数据输出文件位置
func (code *CodeRunner) GetOutputPath() string {
	return filepath.Join(code.WorkSpace, OutputFileName)
}

func (code *CodeRunner) Exec() (err error) {
	defer func() {
		if r := recover(); r != nil {
			jsonBs, _ := json.Marshal(code)
			errInfo := fmt.Sprintf("codeRunner exec failed, codeInfo :%s", string(jsonBs))
			err = errors.New(errInfo)
		}
	}()
	switch code.CodeType {
	case LINUX_SHELL:
		{
			err = code.runLinux_Shell()
		}
	case POWER_SHELL:
		{
			err = code.runPower_Shell()
		}
	case PYTHON:
		{
			err = code.runPython()
		}
	case GOLANG:
		{
			err = code.runGo()
		}
	default:
		{
			//todo
		}
	}
	return
}

func (code *CodeRunner) runLinux_Shell() error {
	return nil
}
func (code *CodeRunner) runPower_Shell() error {
	return nil
}
func (code *CodeRunner) runPython() error {
	return nil
}
