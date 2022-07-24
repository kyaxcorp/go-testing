package main

import (
	"context"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.CommandContext(context.Background(), "who")
	_err := cmd.Run()
	if _err != nil {
		log.Fatalln("failed to run", _err)
	}

	outputBuffer, _err := cmd.Output()
	if _err != nil {
		log.Fatalln("failed output", _err)
	}
	log.Println(string(outputBuffer))
}
