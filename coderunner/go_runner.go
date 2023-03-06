package coderunner

import (
	"errors"
	"log"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// 代码执行输入输出目录
const InputFileName = "__goFlowInput"
const OutputFileName = "__goFlowOutput"

type Go_runner struct{}

func (go_runner *Go_runner) ExecCode(code *CodeRunner) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("函数解析异常")
		}
	}()

	intp := interp.New(interp.Options{}) // 初始化一个 yaegi 解释器
	intp.Use(stdlib.Symbols)             // 允许脚本调用（几乎）所有的 Go 官方 package 代码
	intp.Eval(code.SourceCode)           // src 就是上面的 Go 代码字符串
	v, _ := intp.Eval("Solution.Run")
	fu := v.Interface().(func(string, string))
	doExec(code, fu)
	return err
}

func doExec(code *CodeRunner, fu func(string, string)) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("函数执行异常")
		}
	}()
	fu(code.GetInputPath(), code.GetOutputPath())
}

func init() {
	CodeExecutorMap[GOLANG] = &Go_runner{}
}
