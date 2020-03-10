package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	for i, value := range arr {
		fmt.Println(i, value)
	}

}
