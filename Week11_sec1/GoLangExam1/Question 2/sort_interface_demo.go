package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type Rankable interface {
	GetScore() int
}

func (p Person) GetScore() int {
	return p.Age
}

type Ranker interface {
	Rank(items []Rankable)
}

type DescendingRanker struct{}

func (r DescendingRanker) Rank(items []Rankable) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].GetScore() > items[j].GetScore()
	})
}

func main() {
	people := []Person{
		{Name: "Ali", Age: 25},
		{Name: "Sara", Age: 30},
		{Name: "Omar", Age: 20},
	}

	sort.Sort(ByAge(people))
	fmt.Println("Sorted by Age (Ascending):")
	for _, p := range people {
		fmt.Printf("%s - %d\n", p.Name, p.Age)
	}

	rankableItems := []Rankable{
		Person{Name: "Ali", Age: 25},
		Person{Name: "Sara", Age: 30},
		Person{Name: "Omar", Age: 20},
	}

	ranker := DescendingRanker{}
	ranker.Rank(rankableItems)

	fmt.Println("\nRanked by Score (Descending):")
	for _, item := range rankableItems {
		p := item.(Person)
		fmt.Printf("%s - %d\n", p.Name, p.GetScore())
	}
}
