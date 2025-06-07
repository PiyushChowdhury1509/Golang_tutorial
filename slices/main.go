package main

import "fmt"

func main() {
	//unitialised slice
	var s1 []int
	fmt.Println("string: ", s1, "size: ", len(s1), "capacity: ", cap(s1))
	//unitialised slice is equal to nil
	fmt.Println(s1 == nil)

	//slice of a particular size
	s2 := make([]int, 3) //give another number if want another capacity
	s2[0] = 1
	s2[1] = 2
	s2[2] = 3
	fmt.Println("slice2:", s2, "length:", len(s2), "capacity:", cap(s2))

	//appending in a slice
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	s1 = append(s1, 4)
	fmt.Println("slice1:", s1, "length:", len(s1), "capacity:", cap(s1))

	//copying a slice
	s3 := make([]int, 3)
	copy(s3, s2)
	fmt.Println("slice3:", s3)

	//initializing a slice with default values
	s4 := []int{1, 2, 3, 4, 5, 6}

	//slice operator, this creates a slice with the syntax [low:high]. high isnt included
	s5 := s4[2:5]
	fmt.Println("slice5 after first operation:", s5)
	s5 = s4[:5]
	fmt.Println("slice5 after second operation:", s5)
	s5 = s4[2:]
	fmt.Println("slice5 after third operation:", s5)

	//2d slices
	s6 := make([][]int, 3)
	count := 1
	for i := 0; i < len(s6); i++ {
		innerLength := i + 1
		s6[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			s6[i][j] = count
			count++
		}
	}
	for i := range len(s6) {
		for j := range len(s6[i]) {
			fmt.Printf("%d ", s6[i][j])
		}
		fmt.Printf("\n")
	}
}
