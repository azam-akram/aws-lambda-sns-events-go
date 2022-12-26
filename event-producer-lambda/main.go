package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/model"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/utils"
)

func HandleRequest(ctx context.Context, event *model.Event) error {
	log.Println("Request received: ", event)

	outputEvent := model.Event{
		ID:        event.ID,
		Name:      "SumRequested",
		Source:    "Calculation Requester",
		EventTime: time.Now().Format(time.RFC3339),
		Payload: model.Payload{
			Number1: event.Payload.Number1,
			Number2: event.Payload.Number2,
		},
	}

	log.Println("Event to publish:", outputEvent)

	msgId, err := utils.PublishEvent(ctx, &outputEvent)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Event published to SNS, msgId = ", msgId)

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
