package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	count := 0
	unreadableLetters := map[rune]bool{'R': true, 'S': true, 'T': true, 'Z': true, 'r': true, 's': true, 't': true, 'z': true}

	for _, char := range text {
		if unreadableLetters[char] {
			count++
		}
	}

	return count

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
