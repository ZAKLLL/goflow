package coderunner

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
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
func (code *CodeRunner) runGo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("函数解析异常")
		}
	}()

	intp := interp.New(interp.Options{}) // 初始化一个 yaegi 解释器
	intp.Use(stdlib.Symbols)             // 允许脚本调用（几乎）所有的 Go 官方 package 代码
	intp.Eval(code.SourceCode)           // src 就是上面的 Go 代码字符串
	v, _ := intp.Eval("Solution.Run")
	fu := v.Interface().(func())
	code.doRun(fu)
	return err
}

func (code *CodeRunner) doRun(fu func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("函数执行异常")
		}
	}()
	fu()

}
