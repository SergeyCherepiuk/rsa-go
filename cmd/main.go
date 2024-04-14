package main

import (
	"log"
	"os"
)

func main() {
	message, err := os.ReadFile("message.txt")
	if err != nil {
		log.Fatal(err)
	}

	// blocks := splitter.Split(message, 10)
}
