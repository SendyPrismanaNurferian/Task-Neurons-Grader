package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (j Junior) GetBonus() float64 {
	prorata := float64(j.WorkingMonth) / 12.0
	if prorata > 1 {
		prorata = 1
	}
	return float64(j.BaseSalary) * prorata
}

func (s Senior) GetBonus() float64 {
	prorata := float64(s.WorkingMonth) / 12.0
	if prorata > 1 {
		prorata = 1
	}
	return float64(s.BaseSalary)*2*prorata + s.PerformanceRate*float64(s.BaseSalary)
}

func (m Manager) GetBonus() float64 {
	prorata := float64(m.WorkingMonth) / 12.0
	if prorata > 1 {
		prorata = 1
	}
	return float64(m.BaseSalary)*2*prorata + m.PerformanceRate*float64(m.BaseSalary) + m.BonusManagerRate*float64(m.BaseSalary)
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	totalBonus := 0.0
	for _, emp := range employees {
		totalBonus += emp.GetBonus()
	}
	return totalBonus
}

func main() {
	// Test EmployeeBonus
	junior := Junior{Name: "Junior 1", BaseSalary: 100000, WorkingMonth: 12}
	fmt.Println(EmployeeBonus(junior)) // Output: 100000

	senior := Senior{Name: "Senior 1", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	fmt.Println(EmployeeBonus(senior)) // Output: 250000

	manager := Manager{Name: "Manager 1", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5}
	fmt.Println(EmployeeBonus(manager)) // Output: 300000

	// Test TotalEmployeeBonus
	juniorA := Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12}
	juniorB := Junior{Name: "Junior B", BaseSalary: 100000, WorkingMonth: 12}
	juniorC := Junior{Name: "Junior C", BaseSalary: 100000, WorkingMonth: 12}
	fmt.Println(TotalEmployeeBonus([]Employee{juniorA, juniorB, juniorC})) // Output: 300000

	seniorA := Senior{Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	seniorB := Senior{Name: "Senior B", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	fmt.Println(TotalEmployeeBonus([]Employee{seniorA, seniorB})) // Output: 250000

	juniorA = Junior{Name: "Junior A", BaseSalary: 100000, WorkingMonth: 12}
	seniorA = Senior{Name: "Senior A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5}
	managerA := Manager{Name: "Manager A", BaseSalary: 100000, WorkingMonth: 12, PerformanceRate: 0.5, BonusManagerRate: 0.5}
	fmt.Println(TotalEmployeeBonus([]Employee{juniorA, seniorA, managerA})) // Output: 650000
}
