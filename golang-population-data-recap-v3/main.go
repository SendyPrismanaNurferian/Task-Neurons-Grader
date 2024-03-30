package main

import (
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	population := make([]map[string]interface{}, 0)

	for _, item := range data {
		fields := strings.Split(item, ";")
		person := make(map[string]interface{})

		// Name
		person["name"] = fields[0]

		// Age
		age, _ := strconv.Atoi(fields[1])
		person["age"] = age

		// Address
		person["address"] = fields[2]

		// Height
		if len(fields[3]) > 0 {
			height, _ := strconv.ParseFloat(fields[3], 64)
			person["height"] = height
		}

		// IsMarried
		if len(fields) > 4 && len(fields[4]) > 0 {
			isMarried, _ := strconv.ParseBool(fields[4])
			person["isMarried"] = isMarried
		}

		population = append(population, person)
	}

	return population
}

func main() {
	// Test cases
	data1 := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	result1 := PopulationData(data1)
	for _, p := range result1 {
		println("Name:", p["name"].(string), ", Age:", p["age"].(int), ", Address:", p["address"].(string))
		if height, ok := p["height"].(float64); ok {
			println("Height:", height)
		}
		if isMarried, ok := p["isMarried"].(bool); ok {
			println("Is Married:", isMarried)
		}
	}

	data2 := []string{"Jaka;25;Jakarta;false;170.1", "Anggi;24;Bandung;;"}
	result2 := PopulationData(data2)
	for _, p := range result2 {
		println("Name:", p["name"].(string), ", Age:", p["age"].(int), ", Address:", p["address"].(string))
		if height, ok := p["height"].(float64); ok {
			println("Height:", height)
		}
		if isMarried, ok := p["isMarried"].(bool); ok {
			println("Is Married:", isMarried)
		}
	}

	data3 := []string{}
	result3 := PopulationData(data3)
	for _, p := range result3 {
		println("Name:", p["name"].(string), ", Age:", p["age"].(int), ", Address:", p["address"].(string))
		if height, ok := p["height"].(float64); ok {
			println("Height:", height)
		}
		if isMarried, ok := p["isMarried"].(bool); ok {
			println("Is Married:", isMarried)
		}
	}
}
