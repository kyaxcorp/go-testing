package main

import (
	"log"
	"time"
)

// so, looking at this example, there is no queue in a channel!
// if you want to more data through a channel, you should create an array of channels?!

func main() {
	chan1 := make(chan bool)

	// 2 goroutines
	// 1 will read from the channel
	// another will push

	go func() {
		for {

			time.Sleep(time.Second)
			// let's print the value
			log.Println(<-chan1)
		}
	}()
	go func() {
		for {
			//time.Sleep(time.Millisecond * 500)
			time.Sleep(time.Millisecond)
			chan1 <- false
			log.Println("pushing val")
		}
	}()
	select {}
}
