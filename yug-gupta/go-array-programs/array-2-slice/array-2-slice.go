package main

import "fmt"

func main() {
	arr2 := []int{10, 20, 30, 40, 50, 60}
	slice_arr2 := arr2[3:]
	slice2_arr2 := arr2[:3]
	slice3_arr2 := arr2[:5]
	arr2 = append(arr2, 0, 1)
	fmt.Println(arr2)

	for _, value := range slice_arr2 {
		fmt.Println(value)
	}
	fmt.Println(slice2_arr2)
	fmt.Println(slice3_arr2)
}
