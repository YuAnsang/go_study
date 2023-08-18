package main

import "fmt"

func main() {
	//
	var slice1 = make([]int, 3, 5)
	var slice2 = append(slice1, 4, 5)
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))

	slice1[1] = 100 // slice2까지 바뀌어버림
	fmt.Println("After change second element")
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))
}
