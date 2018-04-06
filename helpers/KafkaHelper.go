package helpers

import (
	"context"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// ReadFromKafka is getting messages from specified Apache Kafka topic
func ReadFromKafka(topic string) []string {

	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message

	var messageList []string

	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		log.Println(string(b))
		messageList = append(messageList, string(b))
	}

	log.Println(messageList)

	batch.Close()
	conn.Close()
	return messageList
}

// WriteToKafka is writing messages into specified Apache Kafka topic
func WriteToKafka(topic string, partition int, message []byte) {

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: message},
	)
	conn.Close()
}
