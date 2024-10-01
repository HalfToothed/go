package main

import (
	"fmt"
	"log"

	"go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	var names []string = []string{"harry", "ron", "hermione", "albus"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
