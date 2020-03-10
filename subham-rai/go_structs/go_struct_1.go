package main

import "fmt"

type rect struct {
	
	left int
	right int
	top int
	bottom int
}


func main(){

	var rect_val = rect{2,3,3,2}
	fmt.Println(rect_val.left," ", rect_val.right,"  ", rect_val.top,"  ", rect_val.bottom)

	fmt.Println(rect_val.area())
}

func (rectangle *rect) area() int{
	return rectangle.top*rectangle.left
}
