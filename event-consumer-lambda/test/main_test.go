package main_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/azam-akram/aws-lambda-sns-events-go/common/constant"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_Main_Success(t *testing.T) {

	assertThat := assert.New(t)

	byteValue, errReadFile := ioutil.ReadFile(constant.LibrariesJSONPathConsumer)
	if errReadFile != nil {
		log.Fatal(errReadFile)
	}

	var libraries model.Libraries

	json.Unmarshal(byteValue, &libraries)
	assertThat.NotNil(libraries)
}
