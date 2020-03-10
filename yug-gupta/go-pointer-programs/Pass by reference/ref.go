package main

import "fmt"

func main() {
	x := 2
	changeValNow(&x)
	fmt.Println(x)
	fmt.Println("Memory address=", &x)

}
func changeValNow(val *int) {
	*val++
	fmt.Println("Memory address of *val", val)
}
