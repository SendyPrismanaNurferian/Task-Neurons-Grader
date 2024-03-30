package main

import "sort"

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	}

	// Menggunakan map untuk menghitung frekuensi kemunculan setiap tanggal
	dateFrequency := make(map[int]int)
	for _, dates := range villager {
		for _, date := range dates {
			dateFrequency[date]++
		}
	}

	// Membuat slice untuk menyimpan tanggal-tanggal yang muncul pada setiap jadwal
	var commonDates []int
	for date, frequency := range dateFrequency {
		if frequency == len(villager) {
			commonDates = append(commonDates, date)
		}
	}

	// Mengurutkan tanggal-tanggal yang sama
	sort.Ints(commonDates)

	return commonDates
}
