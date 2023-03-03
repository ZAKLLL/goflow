package main

import (
	"github.com/s8sg/goflow/defineflows"
	goflow "github.com/s8sg/goflow/v1"
)

func main() {
	fs := &goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		OpenTraceUrl:      "localhost:5775",
		WorkerConcurrency: 5,
	}
	defineflows.DoRegister(fs)
	err := fs.Start()
	if err != nil {
		panic(err)
	}
}
