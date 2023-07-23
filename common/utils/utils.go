package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/azam-akram/aws-lambda-external-sns-topic-go/common/model"
)

func getAWSRegion() string {
	region := os.Getenv("AWS_REGION")
	fmt.Println("AWS region: ", region)
	return region
}

func getSNSTopicARN() string {
	topicArn := os.Getenv("SNS_TOPIC_ARN")
	fmt.Println("SNS Topic ARN: ", topicArn)
	return topicArn
}

func PublishEvent(_ context.Context, event model.Event) (msgId string, err error) {
	region := getAWSRegion()
	awsConfig := &aws.Config{
		Region: &region,
	}

	snsSession, err := session.NewSession(awsConfig)
	if err != nil {
		return "", err
	}
	snsClient := sns.New(snsSession)

	eventBytes, err := json.Marshal(event)
	if nil != err {
		return "", err
	}
	payload := string(eventBytes)

	snsInput := &sns.PublishInput{
		Message:  aws.String(payload),
		TopicArn: aws.String(getSNSTopicARN()),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"name": {
				DataType:    aws.String("String"),
				StringValue: aws.String(event.Name),
			},
		},
	}

	snsMsg, err := snsClient.Publish(snsInput)
	if err != nil {
		return "", err
	}

	fmt.Println("Published event: ", snsMsg)

	return *snsMsg.MessageId, nil
}

func GetSumCompletedEvent(event *model.Event) error {
	event.Name = "SumCompleted"
	event.Source = "Calculation Service"
	event.EventTime = time.Now().Format(time.RFC3339)

	sum := 0
	for _, num := range event.Payload.Numbers {
		sum += num
	}
	event.Payload.Sum = sum

	fmt.Println("Event to publish: ", event)

	return nil
}
