package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go sheep("sheep", c)
	for msg := range c {
		fmt.Println(msg)
	}

}

func sheep(str string, c chan string) {
	for i := 0; i <= 5; i++ {
		c <- str
		time.Sleep(time.Second)
	}
	close(c)
}
