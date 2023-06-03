package main

import "fmt"

func main() {
	fmt.Println("Lesson on Array")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit list length: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList)
}
