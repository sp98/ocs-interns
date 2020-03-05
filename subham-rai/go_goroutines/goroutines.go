package main

import ("fmt"
	"time"
)

func number() {
	for i:=0; i<=5;i++{
		time.Sleep(250*time.Millisecond)
		fmt.Println(i)
	}
}


func alpha() {
	for i:='a'; i<='e';i++{
		time.Sleep(500*time.Millisecond)
		fmt.Printf("%c\n",i)
	}
}

func main() {
	go number()
	go alpha()
	time.Sleep(3000*time.Millisecond)
	fmt.Println("main done")
}








