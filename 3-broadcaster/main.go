package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	var messages []any

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any

		err := json.Unmarshal(msg.Body, &body)
		if err != nil {
			return err
		}

		messages = append(messages, body["message"])
		log.Println("Message is not of type int")

		var reply = make(map[string]any)
		reply["type"] = "broadcast_ok"

		return n.Reply(msg, reply)
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		var reply = make(map[string]any)
		reply["type"] = "read_ok"
		reply["messages"] = messages

		return n.Reply(msg, reply)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		var reply = make(map[string]any)
		reply["type"] = "topology_ok"

		return n.Reply(msg, reply)
	})

	n.Run()
}

