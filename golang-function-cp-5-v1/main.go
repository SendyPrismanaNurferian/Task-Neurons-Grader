package main

import "fmt"

func FindMin(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}

func FindMax(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func SumMinMax(nums ...int) int {
	min := FindMin(nums...)
	max := FindMax(nums...)
	return min + max
}

func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(SumMinMax(333, 456, 654, 123, 111, 1000, 1500, 2000, 3000, 1250, 1111))
}
