package main

import "fmt"

func checkVowel(ch rune) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	return false
}

func main() {
	var s string
	fmt.Printf("Enter a string: ")
	fmt.Scan(&s)

	length := len(s)
	vowel := 0
	cons := 0
	for i := 0; i < length; i++ {
		if checkVowel(rune(s[i])) {
			vowel++
		} else {
			cons++
		}
	}
	fmt.Printf("Vowels: %d, consonants: %d", vowel, cons)
}
