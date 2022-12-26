package utils

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/model"
)

func PublishEvent(ctx context.Context, event *model.Event) (msgId string, err error) {
	eventBytes, err := json.Marshal(event)
	if nil != err {
		return "", err
	}

	payload := string(eventBytes)
	region := "eu-west-1"
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
		TopicArn: aws.String("arn:aws:sns:eu-west-1:107118238565:demo-event-sns-topic"),
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

	return *snsMsg.MessageId, nil
}
