package main

import (
	"log"
	"reflect"
)

type Notification struct {
	Name string
}

func main() {
	item := Notification{}
	item2 := &Notification{}
	log.Println(reflect.TypeOf(item))
	log.Println(reflect.TypeOf(item2))
}
