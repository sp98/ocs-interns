package main

import "fmt"

func main() {
	num := 2
	result := Increment(num)
	fmt.Println(result)
}

func Increment(n int) int {
	n += 2
	return n
}
