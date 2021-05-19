package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	c "github.com/azam-akram/aws-lambda-sns-events-go/common/constant"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/model"
)

func handleRequest() (model.Libraries, error) {

	jsonByte := readJsonFile(c.LibrariesJSONPath)

	libraryJSON := bytesToJSONObject(jsonByte)
	fmt.Println(libraryJSON)

	// Publish this JSON to SNS

	return model.Libraries{Libraries: nil}, nil
}

func readJsonFile(path string) (jsonByte []byte) {
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err) // TODO: better logging
	}
	return byteValue
}

func bytesToJSONObject(byteValue []byte) model.Libraries {

	var libraries model.Libraries
	err := json.Unmarshal(byteValue, &libraries)
	if err != nil {
		log.Fatal(err)
	}

	return libraries
}

func main() {
	lambda.Start(handleRequest)
}
