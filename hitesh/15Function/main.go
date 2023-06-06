package main

import "fmt"

func main() {
	fmt.Println(greeter("Welcome to golang"))
	fmt.Println("Lesson on function")
	fmt.Println(adder(12, 43))
	value, msg := proAdder(12, 23, 34, 54, 56)
	fmt.Printf("%v: %v\n", msg, value)
}

func greeter(msg string) string {
	return msg
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	var total int

	for _, val := range values {
		total += val
	}

	return total, "Sum"
}
