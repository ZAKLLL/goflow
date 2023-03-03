package coderunner

type Power_shell_runner struct{}

func (psr *Power_shell_runner) ExecCode(code *CodeRunner) (err error) {

	return nil
}

func init() {
	CodeExecutorMap[POWER_SHELL] = &Power_shell_runner{}
}
