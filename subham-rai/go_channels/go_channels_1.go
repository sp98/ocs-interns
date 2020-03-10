package main

import ("fmt"
	"time"
)

func main() {
	go count("run")
	count("stop")                 // if we write go in front of both function or make both the fucntion  goroutines than it will print nothing because main goroutine will exit`:wq


}

func count(str string){
	for i:=0; i<=5;i++{
		fmt.Println(i,str)
		time.Sleep(time.Second*2)
	}

}
