package main

import (
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os_lab6-8/internal/data_struct/rbtree"
)

type Message struct {
	NodeID int    `json:"id"`
	Cmd    string `json:"cmd"`
	Arg1   int    `json:"arg1"`
}

func main() {
	// creating myDataStructure which contains nodes with timers
	myDataStructure := rbtree.NewRBTree()

	// connecting to RabbitMQ
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
	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	// getting messages from the queue
	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Consumer is ready to get messages")
	for msg := range msgs {
		var m Message
		err = json.Unmarshal(msg.Body, &m)
		if err != nil {
			log.Println("json unmarshal error: " + err.Error())
			continue
		}
		switch m.Cmd {
		case "get":
			t, GetTimeError := myDataStructure.GetTime(m.NodeID)
			if GetTimeError != nil {
				log.Println("GetTime error: " + GetTimeError.Error())
			} else {
				log.Println("Node ", m.NodeID, " OK: ", t)
			}
		case "insert":
			if InsertErr := myDataStructure.InsertNode(m.NodeID); InsertErr != nil {
				log.Println("InsertNode error: " + InsertErr.Error())
			} else {
				log.Println("Node ", m.NodeID, " OK")
			}
		case "start":
			if StartErr := myDataStructure.StartTimer(m.NodeID); StartErr != nil {
				log.Println("StartTimer error: " + StartErr.Error())
			} else {
				log.Println("Node ", m.NodeID, " OK")
			}
		case "pause":
			if PauseErr := myDataStructure.PauseTimer(m.NodeID); PauseErr != nil {
				log.Println("PauseTimer error: " + PauseErr.Error())
			} else {
				log.Println("Node ", m.NodeID, " OK")
			}
		case "delete":
			if DeleteErr := myDataStructure.DeleteNode(m.NodeID); DeleteErr != nil {
				log.Println("DeleteNode error: " + DeleteErr.Error())
			} else {
				log.Println("Node ", m.NodeID, " OK")
			}
		case "size":
			log.Println("size: ", myDataStructure.Length(), " nodes")
		case "reset":
			if ResetErr := myDataStructure.ResetTimer(m.NodeID); ResetErr != nil {
				log.Println("ResetTimer error: " + ResetErr.Error())
			} else {
				log.Println("Node ", m.NodeID, " OK")
			}
		default:
			log.Println("command " + m.Cmd + " not found")
		}
	}
}
