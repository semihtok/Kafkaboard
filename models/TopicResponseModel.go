package models

//TopicResponseModel defines Apache Kafka topic model for HTTP responses
type TopicResponseModel struct {
	Result []string `json:"result"`
}
