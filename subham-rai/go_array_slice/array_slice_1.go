package main

import "fmt"

func main() {

	//var a[5] int 
	//a[0]=1
	//a[1]=6

	b :=[] int{1,2,3,4,5}

	//for _,value:=range b{
	//fmt.Println(value)
	//}

	//b1 := b[1:5]
	//fmt.Println(b1)

	//fmt.Println("\n")

	c :=make([]int,4,10)  // 4 means first 4 element will be 0 and maximum value will be 10
	for  _,v:=range c{
		fmt.Println(v)
	}
	fmt.Println("\n")
	copy(c,b)
	c = append(c,5,6)
	for  _,v:=range c{
                fmt.Println(v)
        }

}
