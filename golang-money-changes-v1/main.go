package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	total := 0
	for _, product := range products {
		total += product.Price + product.Tax
	}

	change := amount - total
	if change < 0 {
		return nil
	}

	// Pecahan uang yang tersedia
	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	result := make([]int, 0)

	// Menghitung kembalian dengan menggunakan pecahan uang yang tersedia
	for _, coin := range coins {
		for change >= coin {
			result = append(result, coin)
			change -= coin
		}
	}

	// Mengurutkan kembalian dari terbesar ke terkecil
	sort.Sort(sort.Reverse(sort.IntSlice(result)))

	return result
}

func main() {
	// Test case
	change1 := MoneyChanges(10000, []Product{{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300}})
	fmt.Println(change1) // Output: [1000 200]

	change2 := MoneyChanges(30000, []Product{{Name: "baju 1", Price: 10000, Tax: 1000}, {Name: "Sepatu", Price: 15550, Tax: 1555}})
	fmt.Println(change2) // Output: [1000 500 200 100 50 20 20 5]

	change3 := MoneyChanges(5500, []Product{{Name: "Baju", Price: 5000, Tax: 500}})
	fmt.Println(change3) // Output: []
}
