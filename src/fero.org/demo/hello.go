package main

import (
	"fmt"

	"fero.org/msg"
)

func main() {
	fmt.Println(msg.GetMessage())

	var myMap map[string]int
	myMap = make(map[string]int)

	myMap["Hello"] = 1
	myMap["World"] = 2
	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}
