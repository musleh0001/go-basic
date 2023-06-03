package main

import "fmt"

// LoginToken global const variable
const LoginToken string = "afoijfea5.adsfoiijfe.8afafe"

func main() {
	fmt.Println("This program is about Variables")

	var name string
	name = "Khan"
	fmt.Println(name)

	username := "Musleh"
	fmt.Println("Username is: ", username)
	fmt.Printf("Username type is: %T\n", username)

	var isLoggedIn = false
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn type is: %T\n", isLoggedIn)

	var number uint8 = 128
	fmt.Printf("Number type is: %T\n", number)

	var str string
	fmt.Println("str is: ", str)
	fmt.Printf("str type is: %T\n", str)

	fmt.Println("Token is: ", LoginToken)
}
