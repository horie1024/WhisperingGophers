// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//  Hello!
// Produces this output:
//  {"Body":"Hello!"}
//
package main

import (
	_ "bufio"
	"compress/gzip"
	"encoding/base64"
	_ "encoding/json"
	"io"
	_ "log"
	"os"
	"strings"
)

type Message struct {
	Body string
}

func main() {

	var r io.Reader
	r = strings.NewReader(data)
	r = base64.NewDecoder(base64.StdEncoding, r)
	r, _ = gzip.NewReader(r)
	io.Copy(os.Stdout, r)

	// TODO: Create a new bufio.Scanner reading from the standard input.
	// TODO: Create a new json.Encoder writing into the standard output.
	for /* TODO: Iterate over every line in the scanner */ {
		// TODO: Create a new message with the read text.
		// TODO: Encode the message, and check for errors!
	}
	// TODO: Check for a scan error.
}

const data = `
H4sIAAAJbogA/1SOO5KDQAxE8zlFZ5tQXGCjjfYIjoURoPKgcY0E57f4VZlQXf2e+r8yOYbMZJhoZWRxz3wkCVjeReETS0VHz5fBCzpxxg/PbfrT/gacCjbjeiRNOChaVkA9RAdR8eVEw4vxa0Dcs3Fe2ZqowpeqG79L995l3VaMBUV/02OS+B6kMWikwG51c8n5GnEPr11F2/QJAAD//z9IppsHAQAA
`
