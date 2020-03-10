package main

import (
	"fmt"
	"time"
)

func main() {
	go count("Sheep")
	count("Dragon")
	//defer count1("Done")

}
func count(toPrint string) {
	for i := 0; true; i++ {
		fmt.Println(i, toPrint)
		time.Sleep(time.Millisecond * 500)
	}
}
