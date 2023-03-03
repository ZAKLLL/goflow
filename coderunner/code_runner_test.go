package coderunner

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

func t1() {
	f, _ := os.Create(fmt.Sprintf("C:\\Users\\张家魁\\Desktop\\tt%d.txt", rand.Intn(10000)))
	ioutil.WriteFile(f.Name(), []byte("HelloWorls"), 0644)
	defer f.Close()
}
func TestT1(t *testing.T) {
	t1()
}

func TestCodeGo(t *testing.T) {
	const src = `
	package Solution
	import (
		"fmt"
		"io/ioutil"
		"math/rand"
		"os"
	)
	func Run()  {
		if(1==2){
			panic(1)
		}

		fmt.Println("This is FuckInt Run")
	}
	`

	ci := CodeRunner{
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
