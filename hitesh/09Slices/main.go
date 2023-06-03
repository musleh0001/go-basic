package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Lesson on Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Type of fruit list is: ", fruitList)

	// update slices
	fruitList = append(fruitList, "Mango")
	fmt.Println("Updated fruit list is: ", fruitList)

	highScores := make([]int, 4)

	highScores[0] = 251
	highScores[1] = 951
	highScores[2] = 451
	highScores[3] = 851

	highScores = append(highScores, 555, 666, 321)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// remove element
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
