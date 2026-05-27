package main

import (
	"encoding/json"
	"log"

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

	var totalSum int
	quit := make(chan struct{})

	nc.Subscribe("pipeline.squared", func(m *nats.Msg) {
		var msg Message
		json.Unmarshal(m.Data, &msg)

		if msg.IsLast {
			log.Printf("FINAL SUM: %d \n", totalSum)
			quit <- struct{}{}
			return
		}

		totalSum += msg.Value
	})

	log.Println("Sum service is running and calculating...")
	<-quit
}
