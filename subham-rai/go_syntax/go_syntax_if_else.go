package main

import "fmt"

func main (){

	var a int 
	fmt.Println("enter a number")
	fmt.Scanln(&a) // it is ideal to use err,_ := syntax for input
	// there are others ways also to get terminal input 
	if a <=9 {
		fmt.Println("If condition satisfied")
		a--
	} else if a>=20 {
		fmt.Println("else if condition satisfied")
	} else {
		fmt.Println("Oops !no condition satisfied")
	}

}
