package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	c "github.com/azam-akram/aws-lambda-sns-events-go/common/constant"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/model"
)

func handleRequest(libraries model.Libraries) {

	JSONToString(libraries)

}

func JSONToString(libraries model.Libraries) {

	marshalledJSON, err := json.Marshal(libraries)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(marshalledJSON))
}

func main() {

	byteValue, _ := ioutil.ReadFile(c.LibrariesJSONPath)
	var libraries model.Libraries
	json.Unmarshal(byteValue, &libraries)

	handleRequest(libraries)
	//lambda.Start(handleRequest)
}
