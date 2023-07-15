package main

import (
	"encoding/json"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var dataStore = struct {
	values []int
	*sync.RWMutex
}{
	values:  []int{},
	RWMutex: &sync.RWMutex{},
}

func main() {
	n := maelstrom.NewNode()

	n.Handle("broadcast", func(msg maelstrom.Message) error {
		var body struct {
			Message int `json:"message"`
		}

		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		dataStore.Lock()
		dataStore.values = append(dataStore.values, body.Message)
		dataStore.Unlock()

		return n.Reply(msg, map[string]string{
			"type": "broadcast_ok",
		})
	})

	n.Handle("read", func(msg maelstrom.Message) error {
		dataStore.RLock()
		defer dataStore.RUnlock()

		return n.Reply(msg, map[string]any{
			"type":     "read_ok",
			"messages": dataStore.values,
		})
	})

	n.Handle("topology", func(msg maelstrom.Message) error {
		return n.Reply(msg, map[string]any{
			"type": "topology_ok",
		})
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
