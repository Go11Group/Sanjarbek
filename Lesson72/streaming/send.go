package main

import (
    "log"
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

    producer, err := env.NewProducer(streamName, stream.NewProducerOptions())
    if err != nil {
        log.Fatalf("Failed to create producer: %v", err)
    }

    err = producer.Send(amqp.NewMessage([]byte("Hello world")))
    if err != nil {
        log.Fatalf("Failed to send message: %v", err)
    }

    log.Println("Message sent")
}
