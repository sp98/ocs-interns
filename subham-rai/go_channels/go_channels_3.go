package main

import ("fmt"
	"time"
)

func main() {
	c1:=make(chan string)
	c2:=make(chan string)

	go func(){
	for{
		c1<-"1 second"
		time.Sleep(time.Second*1)
	}
	}()

	go func(){
        for{
                c2<-"2 second"
                time.Sleep(time.Second*2)
        }
        }()

	for i:=0; i<=5;i++{
		select {		//using select will help in printing goroutines data which comes first no need to wait for other to print
		case m1 :=<-c1:
			fmt.Println(m1)
		
		case m2:=<-c2:
			fmt.Println(m2)
		}
	}

}
