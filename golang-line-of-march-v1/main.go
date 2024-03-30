package main

import "sort"

func Sortheight(height []int) []int {
	// Menggunakan fungsi sort.Slice untuk mengurutkan slice heights secara ascending
	sort.Slice(height, func(i, j int) bool {
		return height[i] < height[j]
	})

	return height
}
