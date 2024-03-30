package main

func CountProfit(data [][][2]int) []int {
	// Inisialisasi slice untuk menyimpan total keuntungan
	profits := make([]int, len(data[0]))

	// Iterasi melalui setiap bulan
	for _, branch := range data {
		// Iterasi melalui setiap cabang pada bulan tertentu
		for i, sale := range branch {
			// Menghitung keuntungan dari penjualan dan pengeluaran
			profit := sale[0] - sale[1]
			// Menambahkan keuntungan ke slice hasil
			profits[i] += profit
		}
	}

	return profits
}
