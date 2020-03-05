package main

import "fmt"

func main() {
	defPanic()
}

func defPanic() {
	defer func() {
		fmt.Println(recover())
	}()
	panic("This is a drill")
}
