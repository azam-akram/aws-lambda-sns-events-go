package utils

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/azam-akram/aws-lambda-external-sns-topic-go/common/model"
	"github.com/stretchr/testify/assert"
)

var calStr = `{
    "id": 1,
    "name": "StartingEvent",
    "payload":
    {
        "numbers": [2, 3]
    }
}`

func Test_StringToStruct_Success(t *testing.T) {
	assertThat := assert.New(t)

	var event model.Event
	err := json.Unmarshal([]byte(calStr), &event)
	assertThat.Nil(err)

	assertThat.Equal(event.ID, 1)
	assertThat.Equal(event.Name, "StartingEvent")
	//assertThat.Equal(event.Payload.Number1, 2)
	//assertThat.Equal(event.Payload.Number2, 3)

	err = GetSumCompletedEvent(&event)
	assertThat.Nil(err)

	assertThat.Equal(event.ID, 1)
	assertThat.Equal(event.Name, "SumCompleted")
	assertThat.Equal(len(event.Payload.Numbers), 2)
	assertThat.Equal(event.Payload.Sum, 5)
}

func Test_StringChangeMarshlling_Success(t *testing.T) {
	assertThat := assert.New(t)

	var event model.Event
	err := json.Unmarshal([]byte(calStr), &event)
	assertThat.Nil(err)

	assertThat.Equal(event.ID, 1)
	assertThat.Equal(event.Name, "StartingEvent")
	//assertThat.Equal(event.Payload.Number1, 2)
	//assertThat.Equal(event.Payload.Number2, 3)

	event.Name = "SumRequested"
	event.Source = "Calculation Requester"
	event.EventTime = time.Now().Format(time.RFC3339)

	assertThat.Equal(event.ID, 1)
	assertThat.Equal(event.Name, "SumRequested")
	//assertThat.Equal(event.Payload.Number1, 2)
	//assertThat.Equal(event.Payload.Number2, 3)
}

/*
func Test_StringToStruct_EmployeeShort_Success(t *testing.T) {
	assertThat := assert.New(t)

	jh := NewJsonHandler()

	var emp EmployeeShort
	err := jh.StringToStruct(empStr, &emp)
	assertThat.Nil(err)

	assertThat.Equal(emp.ID, "The ID")
	assertThat.Equal(emp.Name, "The User")
}

func Test_StructToString_Success(t *testing.T) {
	assertThat := assert.New(t)

	employee := &Employee{
		ID:   "The ID",
		Name: "The User",
	}

	jh := NewJsonHandler()
	str, err := jh.StructToString(employee)
	assertThat.Nil(err)

	expectedRes := `{"id":"The ID","name":"The User"}`

	assertThat.Equal(expectedRes, str)
}

func Test_StringToMap_Success(t *testing.T) {
	assertThat := assert.New(t)

	jh := NewJsonHandler()
	jMap, _ := jh.StringToMap(empStr)

	id := jMap["id"].(string)
	user := jMap["name"].(string)

	assertThat.Equal(id, "The ID")
	assertThat.Equal(user, "The User")
}

func Test_MapToString_Success(t *testing.T) {
	assertThat := assert.New(t)

	expectedRes := "{\"id\":\"The ID\",\"name\":\"The User\"}"

	mapData := map[string]interface{}{
		"id":   "The ID",
		"name": "The User",
	}

	jh := NewJsonHandler()
	jsonStr, err := jh.MapToString(mapData)

	assertThat.Nil(err)
	assertThat.Equal(jsonStr, expectedRes)
}

func Test_StringToBytes_Success(t *testing.T) {
	jh := NewJsonHandler()

	jh.StringToBytes(empStr)

	assert.NotNil(t, empStr)
}

func Test_BytesToString_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := NewJsonHandler()

	inputBytes := []byte(`{"id": "The ID", "name": "The User"}`)
	outputString := jh.BytesToString(inputBytes)

	actualBytes := []byte(outputString)

	assertThat.Equal(inputBytes, actualBytes)
}

func Test_StructToBytes_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := NewJsonHandler()

	employee := &Employee{
		ID:   "The ID",
		Name: "The User",
	}
	actualBytes, err := jh.StructToBytes(employee)

	assertThat.Nil(err)
	assertThat.NotNil(actualBytes)
}

func Test_BytesToStruct_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := NewJsonHandler()

	byteValue, err := ioutil.ReadFile("testdata/employee.json")
	assertThat.Nil(err)

	var emp *Employee
	err = jh.BytesToStruct(byteValue, &emp)

	assertThat.Nil(err)
	assertThat.Equal(emp.ID, "The ID")
}

func Test_ModifyInputJson_Success(t *testing.T) {
	assertThat := assert.New(t)
	jh := NewJsonHandler()

	modifiedEmpMap, err := jh.ModifyInputJson(empStr)

	assertThat.Nil(err)
	assert.NotNil(t, modifiedEmpMap)
	assertThat.Equal(modifiedEmpMap["degree"], "phd")
	assertThat.Equal(modifiedEmpMap["name"], "The User")
}

func Test_DisplayAllJsonHandlers_Success(t *testing.T) {
	assertThat := assert.New(t)

	byteValue, err := ioutil.ReadFile("testdata/employee.json")
	assertThat.Nil(err)

	jh := NewJsonHandler()

	str := jh.BytesToString(byteValue)

	jh.DisplayAllJsonHandlers(str)
}
*/
