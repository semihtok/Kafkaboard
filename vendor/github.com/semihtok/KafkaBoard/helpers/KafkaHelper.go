package helpers

import (
	"context"
	"log"
	"os"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// ReadFromKafka is getting messages from specified Apache Kafka topic
func ReadFromKafka(topic string) []string {

	partition := 0

	kafkaURL := os.Getenv("KafkaURL")

	if kafkaURL == "" {
		log.Println("KafkaURL not defined as ENV variable (if you're using docker)")
		log.Panicln("Using default Kafka url : localhost:9092")
		kafkaURL = "localhost:9092"
	}
	conn, _ := kafka.DialLeader(context.Background(), "tcp", kafkaURL, topic, partition)

	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message

	var messageList []string

	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		messageList = append(messageList, string(b))
	}

	batch.Close()
	conn.Close()
	return messageList
}

// WriteToKafka is writing messages into specified Apache Kafka topic
func WriteToKafka(topic string, message []byte, partition int) bool {

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.WriteMessages(
		kafka.Message{Value: message},
	)
	conn.Close()

	if err != nil {
		return false
	} else {
		return true
	}
}
