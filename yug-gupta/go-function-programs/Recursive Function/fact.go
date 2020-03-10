package main

import "fmt"

func main() {
	fmt.Println(fact(3))
}
func fact(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * fact(num-1)
	}
}
