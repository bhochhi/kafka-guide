package consumer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"os"
)

func Consume(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "quickstart-events",
		GroupID: "my-grp-consumer-1",
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		fmt.Printf("%s received: %s",os.Getgid(), string(msg.Value))
	}
}
