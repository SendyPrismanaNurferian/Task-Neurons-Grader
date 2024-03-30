package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	// Split string by spaces to get individual words
	words := strings.Fields(str)

	// Initialize an empty slice to store the reversed words
	reversedWords := make([]string, len(words))

	// Iterate over each word and reverse it
	for a, word := range words {
		reversedWord := reverseString(word)
		// Check if the first character is uppercase
		if unicode.IsUpper(rune(word[0])) {
			// If uppercase, title case the reversed word
			reversedWord = strings.Title(strings.ToLower(reversedWord))
		}
		// Append the reversed word to the slice
		reversedWords[a] = reversedWord
	}

	// Join the reversed words back into a string
	return strings.Join(reversedWords, " ")
}

// Function to reverse a string
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println("Input case 1:")
	str1 := "Aku Sayang Ibu"
	fmt.Printf("str = \"%s\"\n", str1)
	fmt.Println("Expected Output / Behavior:")
	fmt.Printf("\"%s\"\n", ReverseWord(str1))
	fmt.Println()

	fmt.Println("Input case 2:")
	str2 := "ini terlalu mudah"
	fmt.Printf("str = \"%s\"\n", str2)
	fmt.Println("Expected Output / Behavior:")
	fmt.Printf("\"%s\"\n", ReverseWord(str2))
	fmt.Println()

	fmt.Println("Input case 3:")
	str3 := "KITA SELALU BERSAMA"
	fmt.Printf("str = \"%s\"\n", str3)
	fmt.Println("Expected Output / Behavior:")
	fmt.Printf("\"%s\"\n", ReverseWord(str3))
}
