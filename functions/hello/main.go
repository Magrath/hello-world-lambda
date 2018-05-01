package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

type inputMessage struct {
	Name string `json:"name"`
}

type outputMessage struct {
	Hello string `json:"hello"`
}

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest handles the lambda request
func HandleRequest(ctx context.Context, input inputMessage) (string, error) {
	output := outputMessage{Hello: input.Name}
	json, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
