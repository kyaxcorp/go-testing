package main

import (
	"context"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.CommandContext(context.Background(), "who").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("who is logged in", string(out))
}
