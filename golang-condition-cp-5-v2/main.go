package main

import "fmt"

func TicketPlayground(height, age int) int {
	// Mengecek apakah umur di bawah 5 tahun
	if age < 5 {
		return -1
	}

	var price int

	switch {
	case age >= 12:
		price = 100000 // Tiket remaja
	case age >= 10 || height > 150:
		price = 40000
	case age >= 8 || height > 135:
		price = 25000
	case age >= 5 || height > 120:
		price = 15000
	default:
		price = -1 // Kondisi tak terpenuhi
	}

	return price
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
