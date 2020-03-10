package main

import "fmt"

func main() {
	testList := []int{1, 2, 3, 4, 5}
	fmt.Println(sumfunc(testList))
}
func sumfunc(test []int) int {
	sum := 0
	for _, val := range test {
		sum += val

	}
	return sum
}
