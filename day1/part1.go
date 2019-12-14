package main

import (
	"log"
	"strconv"
)

func part1(in <-chan string) {
	c := make(chan int, 10)
	go func() {
		for s := range in {
			if n, err := strconv.Atoi(s); err == nil {
				c <- n
			}
		}

		close(c)
	}()

	total := 0
	for num := range c {
		total += int(num/3) - 2
	}
	log.Print("Total: ", total)
}
