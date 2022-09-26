package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

func main() {

	router := gin.Default()
	router.GET("/", publishMessage)
	router.GET("/defer", publishDeferMessage)

	router.Run(":8080")
}

func publishDeferMessage(c *gin.Context) {
	config := nsq.NewConfig()

	producer, err := nsq.NewProducer("localhost:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	topic := "topic_golang_nsq"
	msg := Message{
		Name:      "Message golang nsq",
		Content:   "Content golang nsq deferred",
		Timestamp: time.Now().String(),
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	err = producer.DeferredPublish(topic, 10*time.Second, payload)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, "ok")
}

func publishMessage(c *gin.Context) {
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

	c.JSON(http.StatusOK, "ok")
}
