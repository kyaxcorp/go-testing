package main

import (
	"log"
	"time"
)

// in this case we cannot read a value from a channel that hasn't any value pushed to it
func main() {
	chan1 := make(chan bool)
	chan1 <- true

	go func() {
		for {

			time.Sleep(time.Second)
			// let's print the value
			log.Println(<-chan1)
		}
	}()
	select {}
}
