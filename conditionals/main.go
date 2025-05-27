package main

import "fmt"

//there is no ternary operator in golang
func main() {
	var n int
	fmt.Println("Enter your age")
	fmt.Scan(&n)

	if n >= 18 {
		fmt.Println("you are an adult")
	} else if n >= 12 { //else or else if cant go to next line as it treats it as another statement
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a child")
	}
}
