package main

import "fmt"

func main() {
	fmt.Println("Lesson on Loops")

	days := []string{"Friday", "Saturday", "Sunday", "Monday", "Tuesday", "Thursday"}

	//for d := 0; d < len(days); d++ {
	//	fmt.Println(days[d])
	//}

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	for index, value := range days {
		fmt.Printf("%v: %v\n", index, value)
	}
}
