package main

import (
	"log"
	"reflect"
)

func main() {
	slice1 := []string{"jordan", "wise", "junior"}
	slice2 := []string{"jordan", "wise", "junior"}

	/*if slice1 == slice2 {

	}*/

	if reflect.DeepEqual(slice1, slice2) {
		log.Println("are equal")
	} else {
		log.Println("aren't equal")
	}

}
