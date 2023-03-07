package flowregistry

import (
	goflow "github.com/s8sg/goflow/v1"
	"testing"
)

func TestRegisterAtRuntime(t *testing.T) {

	type args struct {
		fs       *goflow.FlowService
		flowJson string
	}

	fs := &goflow.FlowService{
		RedisURL: "localhost:6379",
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "myCode2",
			args: args{
				fs:       fs,
				flowJson: "{\n  \"FlowName\": \"testFlow\",\n  \"Description\": \"描述信息\",\n  \"Dag\": {\n    \"Nodes\": [\n      {\n        \"NodeName\": \"node1\",\n        \"NodeCodeTag\": \"node1Code\"\n      },\n      {\n        \"NodeName\": \"node2\",\n        \"NodeCodeTag\": \"node2Code\"\n      }\n    ],\n    \"Edges\": [\n      {\n        \"From\": \"node1\",\n        \"To\": \"node2\"\n      }\n    ]\n  },\n  \"SourceCode\": [\n    {\n      \"CodeTag\": \"node1Code\",\n      \"CodeSrc\": \"package Solution\\n\\t\\timport (\\n\\t\\t\\t\\\"fmt\\\"\\n\\t\\t\\t\\\"io/ioutil\\\"\\n\\t\\t\\t\\\"time\\\")\\n\\n\\t\\tfunc Run(intputPath, outputPath string) {\\n\\t\\t\\tdata, _ := ioutil.ReadFile(intputPath)\\n\\t\\t\\tfmt.Printf(\\\"Hi this codeRunner! input Data :%s \\\\n\\\", string(data))\\n\\t\\t\\ttStr := time.Now().String()\\n\\t\\t\\tioutil.WriteFile(outputPath, []byte(\\\"Hello this is CodeRunner ! \\\"+tStr), 0644)\\n\\t\\t}\",\n      \"CodeType\": 3,\n      \"WorkSpace\": \"C:\\\\Users\\\\张家魁\\\\Desktop\\\\codeRunner22\"\n    },\n    {\n      \"CodeTag\": \"node2Code\",\n      \"CodeSrc\": \"package Solution\\n\\t\\timport (\\n\\t\\t\\t\\\"fmt\\\"\\n\\t\\t\\t\\\"io/ioutil\\\"\\n\\t\\t\\t\\\"time\\\")\\n\\n\\t\\tfunc Run(intputPath, outputPath string) {\\n\\t\\t\\tdata, _ := ioutil.ReadFile(intputPath)\\n\\t\\t\\tfmt.Printf(\\\"Hi this codeRunner! input Data :%s \\\\n\\\", string(data))\\n\\t\\t\\ttStr := time.Now().String()\\n\\t\\t\\tioutil.WriteFile(outputPath, []byte(\\\"Hello this is CodeRunner ! \\\"+tStr), 0644)\\n\\t\\t}\",\n      \"CodeType\": 3,\n      \"WorkSpace\": \"C:\\\\Users\\\\张家魁\\\\Desktop\\\\codeRunner33\"\n    }\n  ]\n}",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RegisterAtRuntime(tt.args.fs, tt.args.flowJson); (err != nil) != tt.wantErr {
				t.Errorf("RegisterAtRuntime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
