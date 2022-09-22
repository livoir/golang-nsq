package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"time"
)

type Message struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

func main() {
	config := nsq.NewConfig()

	producer, err := nsq.NewProducer("localhost:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	topic := "topic_golang_nsq"
	msg := Message{
		Name:      "Message golang nsq",
		Content:   "Content golang nsq",
		Timestamp: time.Now().String(),
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	err = producer.Publish(topic, payload)
	if err != nil {
		log.Println(err)
	}
}
