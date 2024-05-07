package main

import (
	"fmt"
	"time"
)

func main() {

	// for loops, break + continue as usual
	for i := range 3 {
		fmt.Println(i)
	}

	// if else
	a := 7
	b := 8
	if a%2 == 0 || b%2 == 0 {
		output := fmt.Sprintf("%s, %d, %s, %d, %s", "Either", a, "or", b, "is even")
		fmt.Println(output)
	}

	//switch case
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}
}
