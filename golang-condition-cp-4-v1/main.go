package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	// Menghitung total harga per tiket
	totalPrice := float32(VIP*30 + regular*20 + student*10)
	// Mengecek apakah total harga tiket minimal $100 dan hari ganjil/genap
	if totalPrice >= 100 {
		var discount float32
		// Jika hari genap
		if day%2 == 0 {
			if VIP+regular+student < 5 {
				discount = 0.1 // Diskon 10%
			} else {
				discount = 0.2 // Diskon 20%
			}
		} else {
			// Jika hari ganjil
			if VIP+regular+student < 5 {
				discount = 0.15 // Diskon 15%
			} else {
				discount = 0.25 // Diskon 25%
			}
		}
		// Menghitung total harga setelah diskon
		totalPrice -= totalPrice * discount
	}

	return totalPrice
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
