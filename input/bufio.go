package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your intro")
	name, _ := reader.ReadString('\n')
	fmt.Println("your intro is", name)

	/*
		fmt.Println("enter your age")
		var a int
		fmt.Scan(&a)
		fmt.Println("your age is", a)
	*/
}
