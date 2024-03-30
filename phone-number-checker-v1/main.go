package main

import (
	"fmt"
	"strings"
)

func PhoneNumberChecker(number string, result *string) {
	// Fungsi untuk mengubah 08 menjadi 62
	// 0811 -> 62811
	if number[0:2] == "08" {
		number = "62" + number[1:]
	}
	// 08 -> 628
	// 08x -> 628x

	// Fungsi untuk cek valid dan invalid nomor
	// 623 -> 623 => invalid
	// 028 -> 028 => invalid
	if number[0:3] != "628" {
		*result = "invalid"
		return
	}

	if len(number) < 11 {
		*result = "invalid"
		return
	}

	// Fungsi untuk yang valid
	if strings.HasPrefix(number, "62811") || strings.HasPrefix(number, "62812") || strings.HasPrefix(number, "62813") || strings.HasPrefix(number, "62814") || strings.HasPrefix(number, "62815") {
		*result = "Telkomsel"
	} else if strings.HasPrefix(number, "62816") || strings.HasPrefix(number, "62817") || strings.HasPrefix(number, "62818") || strings.HasPrefix(number, "62819") {
		*result = "Indosat"
	} else if strings.HasPrefix(number, "62821") || strings.HasPrefix(number, "62822") || strings.HasPrefix(number, "62823") || strings.HasPrefix(number, "62824") || strings.HasPrefix(number, "62825") {
		*result = "XL"
	} else if strings.HasPrefix(number, "62826") || strings.HasPrefix(number, "62827") || strings.HasPrefix(number, "62828") || strings.HasPrefix(number, "62829") {
		*result = "Tri"
	} else if strings.HasPrefix(number, "62852") || strings.HasPrefix(number, "62853") {
		*result = "AS"
	} else if strings.HasPrefix(number, "62881") || strings.HasPrefix(number, "62882") || strings.HasPrefix(number, "62883") || strings.HasPrefix(number, "62884") || strings.HasPrefix(number, "62885") || strings.HasPrefix(number, "62886") || strings.HasPrefix(number, "62887") || strings.HasPrefix(number, "62888") {
		*result = "Smartfren"
	} else {
		*result = "invalid"
	}

}

func main() {
	var number1 = "081211111111"
	var result1 string
	fmt.Println("input:", number1)
	PhoneNumberChecker(number1, &result1)
	fmt.Println(result1)
	fmt.Println()

	var number2 = "08193456123"
	var result2 string
	fmt.Println("input:", number2)
	PhoneNumberChecker(number2, &result2)
	fmt.Println(result2)
	fmt.Println()

	var number3 = "628523456789"
	var result3 string
	fmt.Println("input:", number3)
	PhoneNumberChecker(number3, &result3)
	fmt.Println(result3)
	fmt.Println()

	var number4 = "081234567"
	var result4 string
	fmt.Println("input:", number4)
	PhoneNumberChecker(number4, &result4)
	fmt.Println(result4)
	fmt.Println()

	var number5 = "08222"
	var result5 string
	fmt.Println("input:", number5)
	PhoneNumberChecker(number5, &result5)
	fmt.Println(result5)
	fmt.Println()

	var number6 = "62811-63815 / 0811-0815"
	var result6 string
	fmt.Println("input:", number6)
	PhoneNumberChecker(number6, &result6)
	fmt.Println(result6)
	fmt.Println()

	var number7 = "62816-63819 / 0816-0819"
	var result7 string
	fmt.Println("input:", number7)
	PhoneNumberChecker(number7, &result7)
	fmt.Println(result7)
	fmt.Println()

	var number8 = "62821-63823 / 0821-0823"
	var result8 string
	fmt.Println("input:", number8)
	PhoneNumberChecker(number8, &result8)
	fmt.Println(result8)
	fmt.Println()

	var number9 = "62827-63829 / 0827-0829"
	var result9 string
	fmt.Println("input:", number9)
	PhoneNumberChecker(number9, &result9)
	fmt.Println(result9)
}
