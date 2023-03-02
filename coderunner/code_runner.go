package coderunner

import (
	"log"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type CodeType int

const (
	LINUX_SHELL CodeType = iota
	POWER_SHELL
	PYTHON
	GOLANG
)

type CodeInfo struct {
	WodeType   CodeType
	WorkSpace  string
	IsAsync    bool
	SourceCode string
}

func (code *CodeInfo) RunLinux_Shell() error {
	return nil
}
func (code *CodeInfo) RunPower_Shell() error {
	return nil
}
func (code *CodeInfo) RunPython() error {
	return nil
}
func (code *CodeInfo) RunGo() error {
	intp := interp.New(interp.Options{}) // 初始化一个 yaegi 解释器
	intp.Use(stdlib.Symbols)             // 允许脚本调用（几乎）所有的 Go 官方 package 代码
	intp.Eval(code.SourceCode)           // src 就是上面的 Go 代码字符串
	v, _ := intp.Eval("Solution.Run")
	fu := v.Interface().(func())
	code.doRun(fu)
	return nil
}

func (code *CodeInfo) doRun(fu func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	fu()

}
