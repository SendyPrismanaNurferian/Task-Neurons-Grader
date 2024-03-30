package main

import (
	"fmt"
	"strings"
)

const (
	JKT = "JKT"
	BDG = "BDG"
	BKS = "BKS"
	DPK = "DPK"
)

func DeliveryOrder(data []string, day string) map[string]float32 {
	deliveryData := make(map[string]float32)

	// Hitung biaya admin berdasarkan hari
	adminFee := 0.0
	if day == "senin" || day == "rabu" || day == "jumat" {
		adminFee = 0.1 // 10%
	} else {
		adminFee = 0.05 // 5%
	}

	for _, entry := range data {
		parts := strings.Split(entry, ":")
		firstName := parts[0]
		lastName := parts[1]
		price := parseFloat(parts[2])
		location := parts[3]

		// Cek apakah lokasi dapat dikirimkan pada hari tersebut
		if isLocationAvailable(location, day) {
			// Hitung total biaya pengiriman
			totalPrice := price + (price * float32(adminFee))
			// Buat kunci unik dari nama pengirim
			key := firstName + "-" + lastName
			// Tambahkan ke map
			deliveryData[key] = totalPrice
		}
	}

	return deliveryData
}

func isLocationAvailable(location, day string) bool {
	switch location {
	case JKT:
		return day == "senin" || day == "selasa" || day == "rabu" || day == "kamis" || day == "jumat" || day == "sabtu"
	case BDG:
		return day == "rabu" || day == "kamis" || day == "sabtu"
	case BKS:
		return day == "selasa" || day == "kamis" || day == "jumat"
	case DPK:
		return day == "senin" || day == "selasa"
	default:
		return false
	}
}

func parseFloat(s string) float32 {
	var f float32
	fmt.Sscanf(s, "%f", &f)
	return f
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
