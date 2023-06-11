package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Lesson on JSON")
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	goCourse := []course{
		{"ReactJS Bootcamp", 299, "KhanAcademe.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 399, "KhanAcademe.com", "bcd789", nil},
		{"Django Bootcamp", 499, "KhanAcademe.com", "3io897", []string{"full-stack", "python"}},
	}

	// package this data as JSON data
	finalJson, err := json.MarshalIndent(goCourse, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
		{
                "course_name": "Django Bootcomp",
                "price": 499,
                "website": "KhanAcademe.com",
                "tags": [
                        "full-stack",
                        "python"
                ]
        }
	`)

	var Course course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		err := json.Unmarshal(jsonDataFromWeb, &Course)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", Course)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data ot key value
	var myOnlineData map[string]interface{}
	err := json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", myOnlineData)
	for key, value := range myOnlineData {
		fmt.Printf("%v: %v\n", key, value)
	}
}
