package main

import (
	"log"
	"time"
)

func main() {

	buffer := make(chan bool, 5)
	go func() {
		for {
			log.Println("sending message")
			buffer <- true
		}
	}()

	time.Sleep(time.Second * 5)
	for {

		select {
		case <-buffer:
			log.Println("new message")
		}
		time.Sleep(time.Second)
	}
}
