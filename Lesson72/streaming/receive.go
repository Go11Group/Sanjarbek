package main

import (
    "log"
    "fmt"
    "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/amqp"
    "github.com/rabbitmq/rabbitmq-stream-go-client/pkg/stream"
)

func main() {
    env, err := stream.NewEnvironment(stream.NewEnvironmentOptions())
    if err != nil {
        log.Fatalf("Failed to create environment: %v", err)
    }

    streamName := "hello-go-stream"
    env.DeclareStream(streamName, &stream.StreamOptions{MaxLengthBytes: stream.ByteCapacity{}.GB(2)})

    messagesHandler := func(consumerContext stream.ConsumerContext, message *amqp.Message) {
        fmt.Printf("Stream: %s - Received message: %s\n", consumerContext.Consumer.GetStreamName(), message.Data)
    }

    _, err = env.NewConsumer(streamName, messagesHandler, stream.NewConsumerOptions().SetOffset(stream.OffsetSpecification{}.First()))
    if err != nil {
        log.Fatalf("Failed to create consumer: %v", err)
    }

    log.Println("Consumer running")
    select {} // Block forever
}
