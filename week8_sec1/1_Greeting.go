package main

import "fmt"

func CreateGreeting(greeting string) func(string) string {
	return func(name string) string {
		return greeting + " " + name
	}
}

func main() {
	sayHello := CreateGreeting("Welcome")

	fmt.Println(sayHello("EN.Abouelfetouh"))
	fmt.Println(sayHello("EN.Hisham"))
}
