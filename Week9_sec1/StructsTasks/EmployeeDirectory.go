package main

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
	HireDate   time.Time
}

func NewEmployee(id int, name, department string, salary float64, hireDate time.Time) (Employee, error) {
	if salary < 0 {
		return Employee{}, fmt.Errorf("invalid salary: %.2f. Salary cannot be negative", salary)
	}
	return Employee{
		ID:         id,
		Name:       name,
		Department: department,
		Salary:     salary,
		HireDate:   hireDate,
	}, nil
}

func (e Employee) YearsEmployed() int {
	now := time.Now()
	years := now.Year() - e.HireDate.Year()
	if now.YearDay() < e.HireDate.YearDay() {
		years--
	}
	return years
}

func (e *Employee) ApplyRaise(percentage float64) error {
	if percentage < 0 {
		return errors.New("raise percentage cannot be negative")
	}
	e.Salary += e.Salary * (percentage / 100.0)
	return nil
}

func main() {
	e1, _ := NewEmployee(1, "Ahmed", "Engineering", 15000, time.Date(2020, 5, 15, 0, 0, 0, 0, time.UTC))
	e2, _ := NewEmployee(2, "Sara", "HR", 12000, time.Date(2018, 2, 10, 0, 0, 0, 0, time.UTC))
	e3, _ := NewEmployee(3, "Mona", "Engineering", 18000, time.Date(2021, 8, 1, 0, 0, 0, 0, time.UTC))
	e4, _ := NewEmployee(4, "Omar", "HR", 15000, time.Date(2019, 11, 20, 0, 0, 0, 0, time.UTC))

	_, err := NewEmployee(5, "Khaled", "Sales", -5000, time.Now())
	if err != nil {
		fmt.Println("Error Creation Test:", err)
	}

	employees := []Employee{e1, e2, e3, e4}

	fmt.Println("\n--- Apply Raise Test ---")
	fmt.Printf("Salary before raise for %s: %.2f\n", employees[0].Name, employees[0].Salary)
	employees[0].ApplyRaise(10)
	fmt.Printf("Salary after 10%% raise: %.2f\n", employees[0].Salary)

	fmt.Println("\n--- Years Employed ---")
	for _, emp := range employees {
		fmt.Printf("%s has been employed for %d years.\n", emp.Name, emp.YearsEmployed())
	}

	sort.Slice(employees, func(i, j int) bool {
		if employees[i].Department == employees[j].Department {
			return employees[i].Salary < employees[j].Salary
		}
		return employees[i].Department < employees[j].Department
	})

	fmt.Println("\n--- Employees Sorted by Department & Salary ---")
	for _, emp := range employees {
		fmt.Printf("Dept: %-12s | Salary: %-8.2f | Name: %s\n", emp.Department, emp.Salary, emp.Name)
	}
}
