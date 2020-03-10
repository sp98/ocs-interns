package main

import (
	"fmt"
)

func main() {
	input := make(chan int, 100) //Channel with 100 buffer
	output := make(chan int, 100)
	go fibCal(input, output)
	for i := 0; i < 100; i++ {
		input <- i
	}
	close(input)
	for k := 0; k < 100; k++ {
		fmt.Println(<-output)
	}

}

func fibCal(input <-chan int, output chan<- int) { //We will only recieve in input channel and only send for output channel
	for n := range input {
		output <- fib(n)
	}

}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
