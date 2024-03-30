package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindShortestName(names string) string {
	// Pisahkan nama-nama dengan spasi, koma, dan titik koma
	nameList := strings.FieldsFunc(names, func(r rune) bool {
		return r == ' ' || r == ',' || r == ';'
	})

	// Urutkan daftar nama secara alfabetis
	sort.Strings(nameList)

	// Iterasi melalui daftar nama untuk menemukan nama terpendek
	shortestName := nameList[0]
	shortestLength := len(shortestName)

	for _, name := range nameList {
		nameLength := len(name)
		if nameLength < shortestLength {
			shortestName = name
			shortestLength = nameLength
		}
	}

	return shortestName
}

func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan"))
	fmt.Println(FindShortestName("Budi;Tia;Tio"))
}
