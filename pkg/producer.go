package pkg

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"time"

	"producer/test/config"
)

type TopicStruct struct {
	Header interface{} `json:"header"`
	Body   string     `json:"body"`
}

func Produce() {
	var NotificationPublishingKafkaTopicStruct = TopicStruct{
		Header: nil,
		Body:   "Hello! Pingign!!!",
	}
	partition := 0
	conn, _ := kafka.DialLeader(context.Background(), "tcp", config.KafkaBroker, config.Topic, partition)
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	out, _ := json.Marshal(NotificationPublishingKafkaTopicStruct)
	conn.WriteMessages(
		kafka.Message{Value: []byte(out)},
	)
	conn.Close()
}