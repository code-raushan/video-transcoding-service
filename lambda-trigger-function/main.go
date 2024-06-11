package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type VideoDetail struct {
    Bucket string `json:"bucket"`
    Key    string `json:"key"`
    Size   int64  `json:"size"`
}

func handleS3Event(ctx context.Context, s3Event events.S3Event) error {
    for _, record := range s3Event.Records {
        s3 := record.S3
        videoDetail := VideoDetail{
            Bucket: s3.Bucket.Name,
            Key:    s3.Object.Key,
            Size:   s3.Object.Size,
        }
        message, err := json.Marshal(videoDetail)
        if err != nil {
            log.Printf("Failed to marshal video detail: %v", err)
            continue
        }

        // Send message to Kafka
        if err := sendMessageToKafka(message); err != nil {
            log.Printf("Failed to send message to Kafka: %v", err)
            continue
        }
        fmt.Printf("Successfully processed and sent message for %s\n", s3.Object.Key)
    }
    return nil
}

func sendMessageToKafka(message []byte) error {
    brokerAddress := os.Getenv("KAFKA_BROKER")
    username := os.Getenv("KAFKA_USERNAME")
    password := os.Getenv("KAFKA_PASSWORD")
    topic := os.Getenv("KAFKA_TOPIC")

    mechanism, err := scram.Mechanism(scram.SHA256, username, password)
    if err != nil {
        return fmt.Errorf("failed to create SCRAM mechanism: %v", err)
    }

    writer := kafka.Writer{
        Addr: kafka.TCP(brokerAddress),
        Topic: topic,
        Transport: &kafka.Transport{
            SASL: mechanism,
            TLS:  &tls.Config{},
        },
    }

    err = writer.WriteMessages(context.Background(), kafka.Message{Value: message})
    if err != nil {
        return fmt.Errorf("could not write message: %v", err)
    }

    return nil
}

func main() {
    lambda.Start(handleS3Event)
}