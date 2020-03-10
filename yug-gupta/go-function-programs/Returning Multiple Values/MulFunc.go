package main

import "fmt"

func main() {
	test := 5
	fmt.Println(nextTwo(test))
}
func nextTwo(testnum int) (int, int) {
	next1 := testnum + 1
	next2 := testnum + 2
	return next1, next2
}
