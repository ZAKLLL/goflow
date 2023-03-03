package coderunner

type Python_runner struct{}

func (go_runner *Python_runner) ExecCode(code *CodeRunner) (err error) {
	return nil
}

func init() {
	CodeExecutorMap[PYTHON] = &Python_runner{}
}
