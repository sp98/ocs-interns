package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	func() {
		count("Hello")
		wg.Done()
	}()
	wg.Wait()
}
func count(val string) {
	for i := 0; i <= 5; i++ {
		fmt.Println(i, val)
		time.Sleep(time.Millisecond * 500)
	}
}
