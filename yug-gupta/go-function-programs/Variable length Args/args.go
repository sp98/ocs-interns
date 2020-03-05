//Program to find sum of any number of elements
package main

import "fmt"

func main() {
	fmt.Println(findSum(1, 2, 3, 4, 5, 6))

}
func findSum(args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}
