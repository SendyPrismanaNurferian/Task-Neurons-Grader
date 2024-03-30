package main

func SchedulableDays(date1 []int, date2 []int) []int {
	emptyDates := []int{}

	// Membuat map untuk menyimpan tanggal-tanggal kosong dari date1
	emptyMap := make(map[int]bool)
	for _, date := range date1 {
		emptyMap[date] = true
	}

	// Memeriksa tanggal-tanggal kosong dari date2 yang juga kosong di date1
	for _, date := range date2 {
		if emptyMap[date] {
			emptyDates = append(emptyDates, date)
		}
	}

	return emptyDates
}
