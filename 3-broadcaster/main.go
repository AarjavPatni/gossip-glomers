package main

import (
	"encoding/json"
	"slices"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	var messages []int

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body map[string]any

		err := json.Unmarshal(msg.Body, &body)
		if err != nil {
			return err
		}

		if !(slices.Contains(messages, int(body["message"].(float64)))) {
			incoming_msg := int(body["message"].(float64))
			messages = append(messages, incoming_msg)

			for _, node := range n.NodeIDs() {
				n.Send(node, body)
			}
		}

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
		var body map[string]any
		err := json.Unmarshal(msg.Body, &body)

		if err != nil {
			return err
		}

		rawArr := body["topology"].(map[string]any)[n.ID()].([]any)
		strArr := make([]string, len(rawArr))

		for i, node := range rawArr {
			strArr[i] = node.(string)
		}

		var reply = make(map[string]any)
		reply["type"] = "topology_ok"

		return n.Reply(msg, reply)
	})

	n.Run()
}

