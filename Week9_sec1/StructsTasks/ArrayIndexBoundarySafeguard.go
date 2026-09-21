package main

import "fmt"

func getElement(slice []string, index int) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("index %d out of bounds", index)
		}
	}()

	result = slice[index]
	return result, nil
}

func main() {
	fruits := []string{"Apple", "Banana", "Cherry"}

	fmt.Println("--- Test 1: Valid Index ---")
	val, err := getElement(fruits, 1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element:", val)
	}

	fmt.Println("\n--- Test 2: Index Out Of Bounds ---")
	val, err = getElement(fruits, 5)
	if err != nil {
		fmt.Println("Error Captured:", err)
	} else {
		fmt.Println("Element:", val)
	}

	fmt.Println("\n--- Test 3: Negative Index ---")
	val, err = getElement(fruits, -1)
	if err != nil {
		fmt.Println("Error Captured:", err)
	} else {
		fmt.Println("Element:", val)
	}
}
