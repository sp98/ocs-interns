package main

import "fmt"

func main() {
	fptr := new(int)
	changeVal(fptr)
	fmt.Println(*fptr)
}
func changeVal(ptr *int) {
	*ptr = 100

}
