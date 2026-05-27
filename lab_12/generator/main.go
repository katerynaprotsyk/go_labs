package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

type Message struct {
	Value  int  `json:"value"`
	IsLast bool `json:"is_last"`
}

func main() {
	nc, err := nats.Connect("nats://nats-server:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Drain()

	time.Sleep(2 * time.Second)

	for i := 1; i <= 100; i++ {
		msg := Message{Value: i, IsLast: false}
		data, _ := json.Marshal(msg)
		nc.Publish("pipeline.numbers", data)
	}

	lastMsg := Message{Value: 0, IsLast: true}
	data, _ := json.Marshal(lastMsg)
	nc.Publish("pipeline.numbers", data)

	nc.Flush()

	log.Println("Generator finished work")
}
