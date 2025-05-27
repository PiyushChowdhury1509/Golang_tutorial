package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter n:")
	fmt.Scan(&n)

	//for loop
	for i := 1; i < n; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")
	//there is no while loop in golang, but we can create it
	var i int = n
	for i > 0 {
		fmt.Printf("%d ", i)
		i--
	}
}
