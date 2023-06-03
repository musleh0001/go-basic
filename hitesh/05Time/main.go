package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Lesson on Time")

	presentTime := time.Now()
	fmt.Println("Present Time: ", presentTime)
	fmt.Println(presentTime.Format("01/02/2006 15:04:05 Monday"))
	fmt.Println(presentTime.Format("01/02/2006 03:04:05 PM Monday"))
	fmt.Println(presentTime.Format("Jan 01, 2006 03:04:05 PM Monday"))

	createdDate := time.Date(1999, time.March, 10, 0, 0, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("Jan 01, 2006 03:04:05 PM Monday"))
}
