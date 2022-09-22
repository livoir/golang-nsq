package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type messageHandler struct{}

type Message struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

func main() {
	config := nsq.NewConfig()

	config.MaxAttempts = 10
	config.MaxInFlight = 5
	config.MaxRequeueDelay = time.Second * 900
	config.DefaultRequeueDelay = time.Second * 10

	topic := "topic_golang_nsq"
	channel := "channel_golang_nsq"

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.AddHandler(&messageHandler{})

	consumer.ConnectToNSQLookupd("localhost:4161")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	consumer.Stop()
}

func (h *messageHandler) HandleMessage(message *nsq.Message) error {
	var request Message
	if err := json.Unmarshal(message.Body, &request); err != nil {
		log.Println("Error when unmarshalling the message body, Err : ", err)
		return err
	}

	log.Println("Message")
	log.Println("----------------")
	log.Println("Name : ", request.Name)
	log.Println("Content : ", request.Content)
	log.Println("Timestamp : ", request.Timestamp)
	log.Println("----------------")
	log.Println("")

	message.Finish()
	return nil
}
