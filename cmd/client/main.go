package main

import goflow "github.com/s8sg/goflow/v1"

func main() {
	fs := &goflow.FlowService{
		RedisURL: "localhost:6379",
	}
	fs.Execute("myflow", &goflow.Request{
		Body: []byte("hallo"),
	})
}
