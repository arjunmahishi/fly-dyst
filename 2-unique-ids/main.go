package main

import (
	"bytes"
	"log"
	"os/exec"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		newUUID, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Fatal(err)
		}

		return n.Reply(msg, map[string]string{
			"type": "generate_ok",
			"id":   string(bytes.TrimSpace(newUUID)),
		})
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
