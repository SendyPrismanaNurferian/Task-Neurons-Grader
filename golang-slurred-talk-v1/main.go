package main

import (
	"fmt"
)

func SlurredTalk(words *string) {
	if words == nil || len(*words) == 0 {
		return
	}

	replaceMap := map[rune]rune{'S': 'L', 's': 'l', 'R': 'L', 'r': 'l', 'Z': 'L', 'z': 'l'}

	result := ""
	for _, char := range *words {
		if replaceChar, ok := replaceMap[char]; ok {
			result += string(replaceChar)
		} else {
			result += string(char)
		}
	}

	*words = result
}

func main() {
	var words1 string = "Steven"
	SlurredTalk(&words1)
	fmt.Println(words1)

	var words2 string = "Saya Steven"
	SlurredTalk(&words2)
	fmt.Println(words2)

	var words3 string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words3)
	fmt.Println(words3)

	var words4 string = ""
	SlurredTalk(&words4)
	fmt.Println(words4)
}
