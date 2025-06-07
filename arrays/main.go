package main

import "fmt"

func main() {
	//in go, uninitialised values in array takes zeroed values, like 0 in int, "" in string
	//and false in bool
	var arr [2]float64
	arr[0] = 6.0
	fmt.Println("array size is", len(arr), "and array is", arr)

	array := [2]int{6, 7}
	fmt.Println(array)

	twoDArray := [2][2]int{{1, 2}, {6, 7}}
	fmt.Println("size is", len(twoDArray), "and array is", twoDArray)
}
