package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
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

	nc.Subscribe("pipeline.even", func(m *nats.Msg) {
		var msg Message
		json.Unmarshal(m.Data, &msg)

		if msg.IsLast {
			nc.Publish("pipeline.squared", m.Data)
			return
		}

		msg.Value = msg.Value * msg.Value
		data, _ := json.Marshal(msg)
		nc.Publish("pipeline.squared", data)
	})

	log.Println("Square service is running...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
