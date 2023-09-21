package main

import "fmt"

func main() {
	var username string
	var password string
	var isLogged bool
	var age float64
	fmt.Printf("%q %q %v %f\n", username, password, isLogged, age)

	// shorthand syntax of initializing variables
	congrats := "happy birthday!"
	fmt.Println(congrats)
	fmt.Printf("Type: %T\n", congrats)

	// multiple assignments
	avgOpenRate, displayMessage := .23, " is the average open rate of your messages"
	fmt.Println(avgOpenRate, displayMessage)

	// type casting
	apxAvgOpenRate := int(avgOpenRate)
	fmt.Printf("Your approximate average open rate is %v\n", apxAvgOpenRate)

	// constant declarations
	const pi = 3.14
	fmt.Printf("PI: %f\n", pi)

	// constant declarations using operations
	const secsInMin = 60
	const minsInHr = 60
	const secsInHr = secsInMin * minsInHr

	// storing formatted string
	msg := fmt.Sprintf("Seconds in a Hour: %v\n", secsInHr)
	fmt.Println(msg)
}
