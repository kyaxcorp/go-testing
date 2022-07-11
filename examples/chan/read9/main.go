package main

import (
	"log"
)

// in this case we cannot read a value from a channel that hasn't any value pushed to it
func main() {
	chan1 := make(chan bool)
	go func() {
		chan1 <- true
	}()

	log.Println("everything set!")
	go func() {
		select {
		case val := <-chan1:
			// let's print the value
			log.Println(val)
		}
	}()
	//select {}
	for {
	}
}
