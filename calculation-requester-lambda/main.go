package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/azam-akram/aws-lambda-external-sns-topic-go/common/model"
	"github.com/azam-akram/aws-lambda-external-sns-topic-go/common/utils"
)

const (
	eventNameSumCompleted    = "SumCompleted"
	eventNameStartingEvent   = "StartingEvent"
	eventSourceCalculator    = "Calculator"
	eventSourceCalcRequester = "Calculation Requester"
)

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) error {
	log.Println("SNS event received:", snsEvent)

	for _, record := range snsEvent.Records {
		var event model.Event
		err := json.Unmarshal([]byte(record.SNS.Message), &event)
		if err != nil {
			log.Println("Error unmarshaling JSON:", err)
			return err
		}

		switch event.Name {
		case eventNameSumCompleted:
			if event.Source == eventSourceCalculator {
				log.Println("Answer received:", event.Payload.Sum)
				return nil
			}

		case eventNameStartingEvent:
			event.Name = "SumRequested"
			event.Source = eventSourceCalcRequester
			event.EventTime = time.Now().Format(time.RFC3339)

			log.Println("Event to publish: ", event)

			if _, err := utils.PublishEvent(context.Background(), event); err != nil {
				return fmt.Errorf("error publishing event: %w", err)
			}

			return nil

		default:
			log.Println("Unknown event, ignoring this..")
		}
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
