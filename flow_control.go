package main

import "fmt"

func main() {
	height := 5.8

	if height > 6 {
		fmt.Println("You are very tall")
	} else if height > 5 {
		fmt.Println("You are not tall enough")
	} else {
		fmt.Println("You are very short")
	}

	// initializing variable inside if scope
	if name := "sumit"; name == "sumit" {
		fmt.Println("You are right")
	}
}
