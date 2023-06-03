package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Lesson User Input\n")

	// create input buffer
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the rating for our Pizza: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of this rating is %T\n", input)
}
