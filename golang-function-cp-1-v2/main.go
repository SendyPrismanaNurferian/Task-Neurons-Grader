package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	// Map untuk mengonversi nomor bulan menjadi nama bulan
	monthNames := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	// Format tanggal dengan menambahkan 0 di depan jika kurang dari 10
	dayStr := fmt.Sprintf("%02d", day)

	// Ambil nama bulan dari map
	monthStr := monthNames[month]

	// Format tahun
	yearStr := fmt.Sprintf("%d", year)

	// Gabungkan semua bagian dengan tanda "-" sebagai pemisah
	return fmt.Sprintf("%s-%s-%s", dayStr, monthStr, yearStr)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
	fmt.Println(DateFormat(31, 12, 2020))
}
