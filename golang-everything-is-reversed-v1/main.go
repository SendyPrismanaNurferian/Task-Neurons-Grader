package main

func ReverseData(arr [5]int) [5]int {
	reversed := [5]int{}
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		reversed[i] = reverseNumber(arr[j])
	}
	return reversed
}

func reverseNumber(num int) int {
	reversed := 0
	for num > 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}
	return reversed
}
