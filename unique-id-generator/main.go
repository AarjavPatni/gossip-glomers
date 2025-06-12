package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	uuid "github.com/google/uuid"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func (msg maelstrom.Message) error {
		// generate a number
		// send as a reply to the client

		var body = make(map[string]any)

		body["type"] = "generate_ok"
		body["id"] = uuid.NewString()

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
