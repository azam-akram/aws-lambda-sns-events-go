package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/azam-akram/aws-lambda-external-sns-topic-go/common/model"
	"github.com/azam-akram/aws-lambda-external-sns-topic-go/common/utils"
)

func processEvent(event *model.Event) error {
	if event.Name != "SumRequested" {
		// Ignore events other than SumRequested
		return nil
	}

	event.Name = "SumCompleted"
	event.Source = "Calculation Service"
	event.EventTime = time.Now().Format(time.RFC3339)

	sum := 0
	for _, num := range event.Payload.Numbers {
		sum += num
	}
	event.Payload.Sum = sum

	fmt.Println("Event to publish: ", event)

	if _, err := utils.PublishEvent(context.Background(), *event); err != nil {
		return fmt.Errorf("error publishing event: %w", err)
	}

	return nil
}

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) error {
	fmt.Println("SNS event received:", snsEvent)

	for _, record := range snsEvent.Records {
		var event model.Event
		if err := json.Unmarshal([]byte(record.SNS.Message), &event); err != nil {
			return fmt.Errorf("error in unmarshaling JSON: %w", err)
		}

		fmt.Println("Unmarshalled Event:", event)

		if err := processEvent(&event); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
