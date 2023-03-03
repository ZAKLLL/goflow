package coderunner

import (
	"testing"
)

func TestCodeGo(t *testing.T) {
	const src = `
	package Solution
	import (
		"fmt"
	)
	func Run()  {
		panic()
		fmt.Println("This is FuckInt Run")
	}
	`

	ci := CodeInfo{
		WodeType:   0,
		WorkSpace:  "",
		IsAsync:    false,
		SourceCode: src,
	}
	err := ci.RunGo()
	if err != nil {
		panic(err)
	}

}
