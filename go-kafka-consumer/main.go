package main

import (
	"context"
	"github.com/bhochhi/kafka-guide/go-kafka-consumer/consumer"
)
func main(){
	consumer.Consume(context.Background())
}
