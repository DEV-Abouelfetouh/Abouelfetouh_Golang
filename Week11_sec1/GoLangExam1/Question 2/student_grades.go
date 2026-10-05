package main

import "fmt"

type Student struct {
	Name   string
	ID     int
	Grades map[string]float64
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0.0
	}
	total := 0.0
	for _, grade := range s.Grades {
		total += grade
	}
	return total / float64(len(s.Grades))
}

func TopStudent(students []Student) *Student {
	if len(students) == 0 {
		return nil
	}

	topIndex := 0
	highestAvg := students[0].Average()

	for i := 1; i < len(students); i++ {
		avg := students[i].Average()
		if avg > highestAvg {
			highestAvg = avg
			topIndex = i
		}
	}

	return &students[topIndex]
}

func main() {
	students := []Student{
		{Name: "Ali", ID: 1, Grades: map[string]float64{"Math": 90, "Science": 80}},
		{Name: "Sara", ID: 2, Grades: map[string]float64{"Math": 95, "Science": 98}},
		{Name: "Omar", ID: 3, Grades: map[string]float64{"Math": 70, "Science": 85}},
	}

	top := TopStudent(students)
	if top != nil {
		fmt.Printf("Top Student: %s, Average: %.2f\n", top.Name, top.Average())
	}
}
