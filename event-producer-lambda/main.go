package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	c "github.com/azam-akram/aws-lambda-sns-events-go/common/constant"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/model"
)

func handleRequest() (model.Libraries, error) {

	jsonByte := readJsonFile(c.LibrariesJSONPath)

	libraryJSON := bytesToJSONObject(jsonByte)
	fmt.Printf("Libray data: %s", libraryJSON)

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

	/*for i := 0; i < len(libraries.Libraries); i++ {
		fmt.Println("Id: ", libraries.Libraries[i].ID)
		fmt.Println("Name: ", libraries.Libraries[i].Name)
		fmt.Println("City: ", libraries.Libraries[i].City)
	}*/

	return libraries
}

func main() {
	//lambda.Start(HandleLambdaEvent)

	handleRequest()
}
