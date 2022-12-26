package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/model"
	"github.com/azam-akram/aws-lambda-sns-events-go/common/utils"
)

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) error {
	log.Println("Context: ", ctx)
	log.Println("SNS event received: ", snsEvent)

	var event model.Event

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS
		log.Printf("[%s %s] Message = %s \n", record.EventSource, snsRecord.Timestamp, snsRecord.Message)

		err := json.Unmarshal([]byte(snsRecord.Message), &event)
		if err != nil {
			return err
		}

		if event.Name == "SumCompleted" {
			// Conclude message exchange here
			if event.Source == "Calculator" {
				log.Println("Answer received: ", event.Payload.Answer)
				return nil
			}

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
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
