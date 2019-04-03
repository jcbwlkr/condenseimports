package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Llongfile)

	f, err := os.OpenFile(os.Args[1], os.O_RDWR, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var out bytes.Buffer

	var start, end bool

	s := bufio.NewScanner(f)
	for s.Scan() {
		ln := s.Bytes()

		if bytes.Equal(ln, []byte("import (")) {
			start = true
		}
		if bytes.Equal(ln, []byte(")")) {
			end = true
		}

		ln = append(ln, '\n')

		// Anything before or after imports gets written back to the file. OR If
		// we're still in the import block and it's not a blank line.
		if !start || end || (!end && len(ln) > 1) {
			if _, err := out.Write(ln); err != nil {
				log.Fatal(err)
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	f.Truncate(0)
	f.Seek(0, 0)

	if _, err := io.Copy(f, &out); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
