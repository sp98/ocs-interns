package main

import "fmt"

func main() {
	
	const pi float64 = 3.141592
	var  try_bool bool = true

	fmt.Printf("%.3f\n",pi)
	fmt.Printf("%d\n", 110)
	fmt.Printf("%t\n",try_bool)
	fmt.Printf("%b \n", 110)  //%b for binary 
	fmt.Printf("%c \n", 110) //%c for character
	fmt.Println("%x", 18)   //%x used for hexcode
	fmt.Println("%e", pi) //%e used for scientific code
	fmt.Println("\n")
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || fale= ", true || false)
	fmt.Println("!true=",!true)
	fmt.Println("\n")
}
