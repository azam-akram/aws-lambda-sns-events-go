package main

import (
	"encoding/json"
	"fmt"

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

	type Libraries struct {
		Name string
	}

	library := &Libraries{Name: "The Library"}
	b, err := json.Marshal(library)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}

func main() {

	//lambda.Start(handleRequest)
}
