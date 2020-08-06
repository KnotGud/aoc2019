package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

const inputPath = "./input"

func main() {
	f, err := os.OpenFile(inputPath, os.O_RDONLY, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(commaSplit)
	c := make(chan string, 10)
	go func() {
		for s.Scan() {
			c <- s.Text()
		}
		close(c)
	}()

	part1(c)
}

// commaSplit is a custom bufio split function that splits on commas, stripping 2
func commaSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, bytes.TrimSpace(data[0:i]), nil
	}
	if atEOF {
		return len(data), bytes.TrimSpace(data), nil
	}
	// request more data
	return 0, nil, nil
}
