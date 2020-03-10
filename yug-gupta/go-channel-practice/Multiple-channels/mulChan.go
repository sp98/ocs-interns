package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)
	go Black("Black", c1)
	go White("White", c2)
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}

func Black(b string, c1 chan string) {
	for i := 0; i < 5; i++ {
		c1 <- b
		time.Sleep(time.Millisecond * 500)
	}
	close(c1)
}

func White(w string, c2 chan string) {
	for j := 0; j < 5; j++ {
		c2 <- w
		time.Sleep(time.Millisecond * 1000)
	}
	close(c2)
}
