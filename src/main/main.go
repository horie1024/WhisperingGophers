// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//  Hello!
// Produces this output:
//  {"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {

	// Create a new bufio.Scanner reading from the standard input.
	s := bufio.NewScanner(os.Stdin)

	// Create a new json.Encoder writing into the standard output.
	enc := json.NewEncoder(os.Stdout)

	// Iterate over every line in the scanner
	for s.Scan() {

		// Create a new message with the read text.
		msg := Message{
			Body: s.Text(),
		}

		// Encode the message, and check for errors!
		err := enc.Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Check for a scan error.
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
