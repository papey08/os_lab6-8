package main

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type Message struct {
	NodeID int    `json:"id"`
	Cmd    string `json:"cmd"`
	Arg1   int    `json:"arg1"`
}

func main() {
	// connecting to rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// creating a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// creating message queue
	q, err := ch.QueueDeclare("TestQueue", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Enter the messages in format <command> <ID of the node>")

	// sending message into queue
	supportedCommands := map[string]struct{}{
		"get":    {},
		"insert": {},
		"start":  {},
		"pause":  {},
		"delete": {},
		"size":   {},
		"reset":  {},
	}
	for {
		var msg Message
		fmt.Scan(&msg.Cmd)
		if _, ok := supportedCommands[msg.Cmd]; ok {
			if msg.Cmd != "size" {
				fmt.Scan(&msg.NodeID)
			}
		}
		msgStr, _ := json.Marshal(msg)
		err = ch.PublishWithContext(ctx,
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        msgStr,
			})
		if err != nil {
			log.Println("Publish error: " + err.Error())
			continue
		}
	}
}
