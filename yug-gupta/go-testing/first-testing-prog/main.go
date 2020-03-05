package main

import (
	"fmt"
)

func main() {
	n := 2
	ans := increment(n)
	fmt.Println(ans)

}

func increment(num int) int {
	num += 2
	return num
}
