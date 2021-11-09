package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)
const (
	BrokerAddress ="localhost:9093"
	Topic = "TODO_OPERATION"
)
var writer *kafka.Writer
var Ctx = context.Background()

func getWriter() *kafka.Writer{
	if writer == nil {
		l := log.New(os.Stdout, "kafka writer: ", 0)
		writer = kafka.NewWriter(kafka.WriterConfig{
			Brokers: []string{BrokerAddress},
			Topic:   Topic,
			// assign the logger to the writer
			Logger: l,
		})
	}
	return writer
}
func Produce(ctx context.Context, key, value interface{}) {

	// intialize the writer with the broker addresses, and the topic
	w := getWriter()
	byteKey,_ := json.Marshal(key)
	byteValue,_  := json.Marshal(value)
	err := w.WriteMessages(ctx, kafka.Message{
		Key: byteKey,
		// create an arbitrary message payload for the value
		Value: byteValue,
	})
	if err != nil {
		log.Fatal(err.Error())
	}

}

