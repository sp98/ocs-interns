package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	sampleStr := "This is Yug Gupta"
	splitStr := strings.Split(sampleStr, "")
	fmt.Println(splitStr)
	sampleList := []string{"Y", "u", "g"}
	joinedList := strings.Join(sampleList, ",")
	listOfLetters := []string{"c", "b", "a"}
	sort.Strings(listOfLetters)
	fmt.Println(listOfLetters)
	fmt.Println(joinedList)

}
