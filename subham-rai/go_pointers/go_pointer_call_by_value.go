package main

import "fmt"

func main(){
	x :=2
	change_x_by_value(x)
	
	fmt.Println(x)
}

func change_x_by_value(x int){

	x =10
}
