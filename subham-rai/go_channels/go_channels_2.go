package main

import (
	"fmt"
	"time"
)

func main() {

	c:=make(chan string)
	go count("run",c)
	for msg :=range c{
		fmt.Println(msg)
	}
	d:=make(chan string)
	go count("stop",d)
        for msg :=range d{
                fmt.Println(msg)
        }
}

func count(m string, c chan string){

	for i:=1;i<=5;i++{
	c <-m
	time.Sleep(time.Millisecond*500)
	}

	close(c)
}
