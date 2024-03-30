package main

import (
	"fmt"
	"math"
)

type School struct {
	Name    string
	Address string
	Grades  []int
}

func (s *School) AddGrade(grades ...int) {
	s.Grades = append(s.Grades, grades...)
}

func Analysis(s School) (float64, int, int) {
	if len(s.Grades) == 0 {
		return 0, 0, 0
	}

	min, max := s.Grades[0], s.Grades[0]
	sum := 0

	for _, grade := range s.Grades {
		sum += grade
		if grade < min {
			min = grade
		}
		if grade > max {
			max = grade
		}
	}

	avg := float64(sum) / float64(len(s.Grades))
	return math.Round(avg*100) / 100, min, max
}

func main() {
	avg, min, max := Analysis(School{
		Name:    "Imam Assidiqi School",
		Address: "Jl. Imam Assidiqi",
		Grades:  []int{100, 90, 80, 70, 60},
	})

	fmt.Println(avg, min, max)
}
