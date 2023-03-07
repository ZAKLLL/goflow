package main

import (
	"flag"
	"github.com/s8sg/goflow/flowregistry"
	goflow "github.com/s8sg/goflow/v1"
)

var (
	port  int
	redis string
	trace string
)

func init() {
	flag.IntVar(&port, "port", 8080, "api server port")
	flag.StringVar(&redis, "redis", "localhost:6379", "redisAddr")
	flag.StringVar(&trace, "trace", "localhost:5775", "openTraceUrl")
}

func main() {
	flag.Parse()
	fs := &goflow.FlowService{
		Port:              port,
		RedisURL:          redis,
		OpenTraceUrl:      trace,
		WorkerConcurrency: 5,
	}
	err := flowregistry.RegisterFlows(fs)
	if err != nil {
		panic(err)
	}
	err = fs.Start()
	if err != nil {
		panic(err)
	}
}
