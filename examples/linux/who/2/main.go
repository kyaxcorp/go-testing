package main

import (
	"context"
	"log"
	"os/exec"
)

func main() {
	// -u --users  list users logged in
	// -H --heading  print line of column headings
	out, err := exec.CommandContext(context.Background(), "who", "-u", "-H").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("who is logged in")
	log.Println(string(out))

	// we need to see the current active users...
}
