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

	libraryJSON := stringToJSONObject()
	fmt.Println(libraryJSON)

	// Publish this JSON to SNS

	return model.Libraries{Libraries: nil}, nil
}

func stringToJSONObject() model.Libraries {

	byteValue, errReadFile := ioutil.ReadFile(c.LibrariesJSONPath)
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}

	var libraries model.Libraries

	errJSONUnmarshall := json.Unmarshal(byteValue, &libraries)
	if errJSONUnmarshall != nil {
		log.Fatal(errReadFile)
	}

	for i := 0; i < len(libraries.Libraries); i++ {
		fmt.Println("Id: ", libraries.Libraries[i].ID)
		fmt.Println("Name: ", libraries.Libraries[i].Name)
		fmt.Println("City: ", libraries.Libraries[i].City)
	}

	return libraries
}

func main() {
	//lambda.Start(HandleLambdaEvent)

	handleRequest()
}
