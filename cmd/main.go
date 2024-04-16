package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/SergeyCherepiuk/rsa-go/internal/rsa"
	"github.com/SergeyCherepiuk/rsa-go/internal/splitter"
)

func main() {
	message, err := os.ReadFile("message.txt")
	if err != nil {
		log.Fatal(err)
	}

	messageParts := splitter.Split(message, 5)

	privateKey := rsa.GeneratePrivateKey()
	publicKey := privateKey.PublicKey()

	fmt.Printf("Key bit length: %d\n", privateKey.N.BitLen())

	encodedParts := make([]string, 0)
	for _, part := range messageParts {
		encodedPart := rsa.Encode(part, publicKey).String()
		encodedParts = append(encodedParts, encodedPart)
	}
	fmt.Printf("Encoded: %s\n", strings.Join(encodedParts, ""))

	decodedParts := make([]string, 0)
	for _, encodedPart := range encodedParts {
		e, _ := new(big.Int).SetString(encodedPart, 10)
		decodedPart := string(rsa.Decode(e, privateKey))
		decodedParts = append(decodedParts, decodedPart)
	}
	fmt.Printf("Decoded: %s\n", strings.Join(decodedParts, ""))
}
