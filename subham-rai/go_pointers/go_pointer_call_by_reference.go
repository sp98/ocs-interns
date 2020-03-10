package main

import "fmt"

func main(){
	x :=2
	change_x_by_ref(&x)
	
	//we create pointer by using keyword using new syntax x = new(int)


	fmt.Println(x)
}

func change_x_by_ref(x *int){

	*x =10
}
