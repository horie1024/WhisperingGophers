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
	_ "compress/gzip"
	_ "encoding/base64"
	_ "encoding/json"
	"fmt"
	_ "io"
	"log"
	_ "os"
	"strings"
)

type Message struct {
	Body string
}

func main() {

	s := bufio.NewScanner(strings.NewReader(input))

	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	// TODO: Create a new bufio.Scanner reading from the standard input.
	// TODO: Create a new json.Encoder writing into the standard output.
	for /* TODO: Iterate over every line in the scanner */ {
		// TODO: Create a new message with the read text.
		// TODO: Encode the message, and check for errors!
	}
	// TODO: Check for a scan error.
}

const input = `A haiku is more
Than just a collection of
Well-formed syllables`
