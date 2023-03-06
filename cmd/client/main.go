package main

import goflow "github.com/s8sg/goflow/v1"

func main() {
	fs := &goflow.FlowService{
		RedisURL: "localhost:6379",
	}
	err := fs.Execute("myCode", &goflow.Request{
		Body: []byte("hallo"),
	})
	if err != nil {
		panic(err)
	}
}
