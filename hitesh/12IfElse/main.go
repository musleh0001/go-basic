package main

import "fmt"

func main() {
	fmt.Println("Lesson on If Else")

	loginCount := 22
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	if loginCount%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	fmt.Println(result)
}
