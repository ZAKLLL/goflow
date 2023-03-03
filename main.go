package main

import (
	"flag"

	defineflow "github.com/s8sg/goflow/defineflows"
	goflow "github.com/s8sg/goflow/v1"
)

var (
	p     int
	redis string
	trace string
)

func init() {
	flag.IntVar(&p, "port", 8080, "api server port")
	flag.StringVar(&redis, "redis", "localhost:6379", "redisAddr")
	flag.StringVar(&trace, "trace", "localhost:5775", "openTraceUrl")
}

func main() {
	flag.Parse()
	fs := &goflow.FlowService{
		Port:              p,
		RedisURL:          redis,
		OpenTraceUrl:      trace,
		WorkerConcurrency: 5,
	}
	defineflow.DoRegister(fs)
	err := fs.Start()
	if err != nil {
		panic(err)
	}
}

// package main

// import (
// 	goflow "github.com/s8sg/goflow/v1"
// )

// func main() {
// 	fs := &goflow.FlowService{
// 		RedisURL: "localhost:6379",
// 	}
// 	fs.Execute("myflow", &goflow.Request{
// 		Body: []byte("hallo"),
// 	})
// }
