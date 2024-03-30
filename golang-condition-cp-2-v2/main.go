package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var percentage float64
	if gender == "laki-laki" {
		percentage = 10
	} else if gender == "perempuan" {
		percentage = 15
	} else {
		return 0.0 // Mengembalikan 0 jika jenis kelamin tidak valid
	}

	bmi := float64(height-100) - (float64(height-100) * (percentage / 100))
	return bmi
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 170))
	fmt.Println(BMICalculator("perempuan", 165))
}
