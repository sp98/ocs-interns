package main

import "fmt"

func main() {
	//m:=make(map[string]int)
	//m["one"]=1
	//fmt.Println(m["one"])

	var n =map[string]int{
	"one": 1,
	"two":2,
	}
	
	fmt.Println(n,"\n")
	
	for k,v := range n{
		fmt.Println(k," ",v)
		
	}

}
