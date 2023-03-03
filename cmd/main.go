package main

import (
	goflow "github.com/s8sg/goflow/v1"
)

func main() {
	fs := &goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		OpenTraceUrl:      "localhost:5775",
		WorkerConcurrency: 5,
	}
	err := fs.Register("myflow", MyWorkFlow)
	if err != nil {
		panic(err)
	}
	err = fs.Start()
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	fs := &goflow.FlowService{
// 		RedisURL: "localhost:6379",
// 	}
// 	fs.Execute("myflow", &goflow.Request{
// 		Body: []byte("hallo"),
// 	})
// }
