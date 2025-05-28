package main

import "fmt"

func main() {
	var number int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&number)
	copy := number
	var reverse int
	for copy > 0 {
		z := copy % 10
		copy /= 10
		reverse = reverse*10 + z
	}
	if reverse == number {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
