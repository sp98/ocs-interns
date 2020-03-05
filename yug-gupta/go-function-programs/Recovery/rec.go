package main

import "fmt"

func main() {
	fmt.Println(problemFreeDiv(5, 0))
	fmt.Println(problemFreeDiv(10, 5))

}
func problemFreeDiv(num1, num2 int) int {
	defer func() {
		recover() // we can print the error caused
	}()
	answer := num1 / num2
	return answer
}

// program continues even after failing the first call
