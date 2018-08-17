package service

import "fmt"

type Item struct {
	Value string
}

func Add(item Item) {
	println("Sorry, not implemented yet")
	fmt.Sprintln("item value: ", item.Value)
}
