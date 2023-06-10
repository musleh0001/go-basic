package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response status: %v\n", response.StatusCode)
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
