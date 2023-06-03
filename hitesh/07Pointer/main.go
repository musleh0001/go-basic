package main

import "fmt"

func main() {
	fmt.Println("Lesson on Pointer")

	//var ptr *uint8
	//fmt.Println("Value of point is: ", ptr)

	num := 23
	ptr := &num
	fmt.Println("Value of num is: ", *ptr)

	*ptr = *ptr + 4
	fmt.Println("Value of num is: ", num)
}
