package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	numStr := strconv.Itoa(numbers)
	maxSum := 0

	for i := 0; i < len(numStr)-1; i++ {
		pair, _ := strconv.Atoi(numStr[i : i+2])
		if pair > maxSum {
			maxSum = pair
		}
	}

	return maxSum
}

func main() {
	fmt.Println(BiggestPairNumber(11223344))
	fmt.Println(BiggestPairNumber(89083278))
}
