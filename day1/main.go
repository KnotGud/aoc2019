package main

import (
	"bufio"
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
	c := make(chan string, 10)
	go func() {
		for s.Scan() {
			c <- s.Text()
		}
		close(c)
	}()

	part2(c)
}
