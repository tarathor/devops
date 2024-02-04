package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	names := []string{"Robb", "Sansa", "Ned"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(quote.Go())

	fmt.Println(message)
}
