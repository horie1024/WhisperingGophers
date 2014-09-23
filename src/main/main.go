// Solution to part 2 of the Whispering Gophers code lab.
//
// This program extends part 1.
//
// It makes a connection the host and port specified by the -dial flag, reads
// lines from standard input and writes JSON-encoded messages to the network
// connection.
//
// You can test this program by installing and running the dump program:
//  $ go get code.google.com/p/whispering-gophers/util/dump
//  $ dump -listen=localhost:8000
// And in another terminal session, run this program:
//  $ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
)

var dialAddr = flag.String("dial", "localhost:8000", "host:port to dial")

type Message struct {
	Body string
}

func main() {

	// Parse the flags.
	flag.Parse()

	// Open a new connection using the value of the "dial" flag.
	c, err := net.Dial("tcp", *dialAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new bufio.Scanner reading from the standard input.
	s := bufio.NewScanner(os.Stdin)

	// Create a json.Encoder writing into the connection you created before.
	enc := json.NewEncoder(c)

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
