package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var similarData []string

	for _, d := range data {
		if strings.Contains(d, input) {
			similarData = append(similarData, d)
		}
	}

	return strings.Join(similarData, ",")
}

func main() {
	fmt.Println(FindSimilarData("mobil", "mobil APV", "mobil Avanza", "motor matic", "motor gede"))
	fmt.Println(FindSimilarData("motor", "mobil APV", "mobil Avanza", "motor matic", "motor gede", "iphone 14", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel"))
}
