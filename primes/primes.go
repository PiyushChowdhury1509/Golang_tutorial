package main

import "fmt"

func main() {
	var number int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&number)

	flag := false
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			flag = true
			break
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
