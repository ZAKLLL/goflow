package coderunner

import (
	"errors"
)

type Linux_shell_runner struct{}

func (lsr *Linux_shell_runner) ExecCode(code *CodeRunner) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("函数解析异常")
		}
	}()

	return nil
}

func init() {
	CodeExecutorMap[LINUX_SHELL] = &Linux_shell_runner{}
}
