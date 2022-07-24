package main

import (
	"context"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.CommandContext(context.Background(), "last").Output()
	// TODO: check by still logged in
	if err != nil {
		log.Fatal(err)
	}
	log.Println("who is logged in", string(out))
	// we need to see the current active users...
}
