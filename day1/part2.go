package main

import (
	"log"
	"strconv"
)

func part2(in <-chan string) {
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
		total += recurseFuelCalc(num)
	}
	log.Print("Total: ", total)
}

func recurseFuelCalc(mass int) int {
	n := int(mass/3) - 2
	if n <= 0 {
		return 0
	} else {
		return n + recurseFuelCalc(n)
	}
}
