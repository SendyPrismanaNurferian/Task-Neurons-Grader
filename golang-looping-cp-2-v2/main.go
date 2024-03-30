package main

import (
	"fmt"
)

// hello World => dlr_o_W o_l_l_e_H
func ReverseString(str string) string {
	ran := []rune(str)
	revers := make([]rune, 0, len(ran))

	for i := len(ran) - 1; i >= 0; i-- {
		if ran[i] != ' ' {
			revers = append(revers, ran[i])
		} else {
			// Jika ditemukan spasi, tambahkan spasi tanpa underscore
			revers = append(revers, ' ')
			continue
		}

		// Tambahkan underscore jika bukan huruf terakhir dan bukan spasi
		if i != 0 && ran[i-1] != ' ' {
			revers = append(revers, '_')
		}
	}

	return string(revers)
}

// gunakan untuk melakukan debug
func main() {
	// Test Case 1
	str1 := "Hello World"
	fmt.Println("Test Case 1:")
	fmt.Println("Input:", str1)
	fmt.Println("Output:", ReverseString(str1))
	fmt.Println()

	// Test Case 2
	str2 := "I am a student"
	fmt.Println("Test Case 2:")
	fmt.Println("Input:", str2)
	fmt.Println("Output:", ReverseString(str2))
}
