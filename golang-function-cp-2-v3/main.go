package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowels := "aiueoAIUEO"
	consonants := "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
	vowelCount := 0
	consonantCount := 0

	// Iterasi setiap karakter dalam string
	for _, char := range str {
		// Jika karakter adalah huruf vokal, tambahkan ke jumlah vokal
		if strings.ContainsRune(vowels, char) {
			vowelCount++
		}
		// Jika karakter adalah huruf konsonan, tambahkan ke jumlah konsonan
		if strings.ContainsRune(consonants, char) {
			consonantCount++
		}
	}

	// Pengecekan apakah terdapat huruf vokal atau konsonan
	noVowelOrConsonant := vowelCount == 0 || consonantCount == 0

	return vowelCount, consonantCount, noVowelOrConsonant
}

func main() {
	vowelCount, consonantCount, noVowelOrConsonant := CountVowelConsonant("Hidup Itu Indah")
	fmt.Printf("%d, %d, %t\n", vowelCount, consonantCount, noVowelOrConsonant)
}
