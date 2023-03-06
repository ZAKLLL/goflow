package main

import (
	"github.com/s8sg/goflow/flowregistry"
	goflow "github.com/s8sg/goflow/v1"
)

func main() {
	fs := &goflow.FlowService{
		Port:              8080,
		RedisURL:          "localhost:6379",
		OpenTraceUrl:      "localhost:5775",
		WorkerConcurrency: 5,
	}
	err := flowregistry.RegisterDefineFlows(fs)
	if err != nil {
		panic(err)
	}
	err = fs.Start()
	if err != nil {
		panic(err)
	}
}
