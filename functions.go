package main 

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

// shorthand syntax for same variable type declaration
func add(x, y int) int {
	return x - y
}

func getNames() (string, string) {
	return "Sumit", " Modak"
}

func main() {
	fmt.Println(add(3, 8))

	fname, lname := getNames()
	fmt.Println(concat(fname, lname))

	// ignoring return values
	_, _ := getNames()
}
