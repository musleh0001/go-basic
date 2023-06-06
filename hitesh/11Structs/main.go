package main

import "fmt"

func main() {
	fmt.Println("Lesson on Struct")

	musleh := User{"Musleh", "musleh@go.dev", true, 23}
	fmt.Println(musleh)
	fmt.Printf("Name: %v\nEmail: %v\nStatus: %v\nAge: %v\n", musleh.Name, musleh.Email, musleh.Status, musleh.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
