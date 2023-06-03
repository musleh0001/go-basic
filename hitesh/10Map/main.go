package main

import "fmt"

func main() {
	fmt.Println("Lesson on Map")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("Lost of all languages: ", languages)

	// delete
	delete(languages, "RB")
	fmt.Println("Lost of all languages: ", languages)

	// loop through map
	for key, value := range languages {
		fmt.Printf("%v: %v\n", key, value)
	}
}
