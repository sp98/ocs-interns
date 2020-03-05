package main

import "fmt"

func main() {

	m := map[string]int{
	"one":1,
	"two":2,
	"three":3,
	}

	//fmt.Println(m)
	//fmt.Println(len(m))
	delete(m,"one")
	fmt.Println(m)
	for k,v:=range m {
	fmt.Println(k,v)
	}
}
