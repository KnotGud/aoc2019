package main

import (
	"fmt"
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

	var memory []int
	for n := range c {
		memory = append(memory, n)
	}
	// modify to 1202
	memory[1] = 12
	memory[2] = 2
	execute(memory)
}

// run computer with program mem
func execute(mem []int) {
	// instruction pointer
	done := false
	for ip := 0; !done; ip += 4 {
		op := mem[ip]
		if op == 99 {
			// print final command
			printCmd(mem, ip)
			done = true
			continue
		}
		printCmd(mem, ip)
		arg1 := mem[ip+1]
		arg2 := mem[ip+2]
		dst := mem[ip+3]
		switch op {
		case 1:
			mem[dst] = mem[arg1] + mem[arg2]
		case 2:
			mem[dst] = mem[arg1] * mem[arg2]
		default:
			log.Panicf("Unknown OpCode <%d> at position <%d>", op, ip)
		}
	}
	//printStack(mem)
	log.Print("Final Value of Index 0: ", mem[0])
}

func printStack(mem []int) {
	for i := 0; i < len(mem); i += 4 {
		printCmd(mem, i)
	}
}

func printCmd(mem []int, i int) {
	switch len(mem) - i {
	case 3:
		fmt.Printf("%3v: %4v %4v  %4v\n", i, mem[i], mem[i+1], mem[i+2])
	case 2:
		fmt.Printf("%3v: %4v %4v\n", i, mem[i], mem[i+1])
	case 1:
		fmt.Printf("%3v: %4v\n", i, mem[i])
	default:
		fmt.Printf("%3v: %4v %4v  %4v %4v\n", i, mem[i], mem[i+1], mem[i+2], mem[i+3])
	}
}
