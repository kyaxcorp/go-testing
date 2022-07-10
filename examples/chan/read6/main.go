package main

import (
	"log"
	"time"
)

// when sending, the goroutine that is sending will be blocked until the reader will select!
func main() {
	chan1 := make(chan bool)

	go func() {
		for {
			chan1 <- true
			log.Println("go 1 -> sent a message!")
			//time.Sleep(time.Second)
		}
	}()

	go func() {
		for {

			time.Sleep(time.Second * 5)
			// let's print the value

			select {
			case chanVal := <-chan1:
				log.Println("reading message", chanVal)
			}
		}
	}()

	for {

	}
	// if using select instead of for, you will receive an error
	// that all goroutines are asleep!
	//select {}
}
