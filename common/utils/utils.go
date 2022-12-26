package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/model"
)

func getAWSRegion() string {
	region := os.Getenv("AWS_REGION")
	log.Println("AWS region: ", region)
	return region
}

func getSNSTopicARN() string {
	topicArn := os.Getenv("SNS_TOPIC_ARN")
	log.Println("SNS Topic ARN: ", topicArn)
	return topicArn
}

func PublishEvent(ctx context.Context, event *model.Event) (msgId string, err error) {
	eventBytes, err := json.Marshal(event)
	if nil != err {
		return "", err
	}

	payload := string(eventBytes)
	region := getAWSRegion()
	awsConfig := &aws.Config{
		Region: &region,
	}

	snsSession, err := session.NewSession(awsConfig)
	if err != nil {
		return "", err
	}

	snsClient := sns.New(snsSession)

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

	log.Println("Published event: ", snsMsg)

	return *snsMsg.MessageId, nil
}
