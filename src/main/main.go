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
	_ "bufio"
	_ "encoding/json"
	"flag"
	"fmt"
	_ "log"
	_ "os"
	"time"
)

var (
	message = flag.String("message", "Hello!", "what to say")
	delay   = flag.Duration("delay", 2*time.Second, "how long to wait")
)

type Message struct {
	Body string
}

func main() {

	flag.Parse()
	fmt.Println(*message)
	time.Sleep(*delay)
}
