package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description,omitempty"`
}

func (p Product) Validate() error {
	if p.Price <= 0 {
		return fmt.Errorf("invalid price for SKU %s: price must be greater than 0 (got %.2f)", p.SKU, p.Price)
	}
	if p.Stock < 0 {
		return fmt.Errorf("invalid stock for SKU %s: stock cannot be negative (got %d)", p.SKU, p.Stock)
	}
	return nil
}

func DecodeProducts(jsonData string) ([]Product, error) {
	var products []Product
	err := json.Unmarshal([]byte(jsonData), &products)
	if err != nil {
		return nil, fmt.Errorf("json decoding error: %w", err)
	}

	for _, p := range products {
		if err := p.Validate(); err != nil {
			return nil, err
		}
	}

	return products, nil
}

func CalculateTotalInventoryValue(products []Product) float64 {
	total := 0.0
	for _, p := range products {
		total += p.Price * float64(p.Stock)
	}
	return total
}

func EncodeProduct(p Product) (string, error) {
	bytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func main() {
	jsonInput := `[
		{
			"sku": "LAP-001",
			"name": "Gaming Laptop",
			"price": 1200.50,
			"stock": 10,
			"description": "High performance laptop with RTX GPU"
		},
		{
			"sku": "MOU-002",
			"name": "Wireless Mouse",
			"price": 25.00,
			"stock": 50
		},
		{
			"sku": "KEY-003",
			"name": "Mechanical Keyboard",
			"price": 85.00,
			"stock": 20,
			"description": ""
		}
	]`

	fmt.Println("--- 1. Decoding JSON Products ---")
	products, err := DecodeProducts(jsonInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range products {
		fmt.Printf("SKU: %s | Name: %s | Price: $%.2f | Stock: %d | Desc: '%s'\n",
			p.SKU, p.Name, p.Price, p.Stock, p.Description)
	}

	fmt.Println("\n--- 2. Calculating Total Inventory Value ---")
	totalValue := CalculateTotalInventoryValue(products)
	fmt.Printf("Total Inventory Value: $%.2f\n", totalValue)

	fmt.Println("\n--- 3. Testing JSON Encoding with 'omitempty' ---")
	pWithoutDesc := Product{
		SKU:   "PAD-004",
		Name:  "Mouse Pad",
		Price: 10.00,
		Stock: 100,
	}

	encodedJSON, err := EncodeProduct(pWithoutDesc)
	if err == nil {
		fmt.Println("Encoded Product (Empty Description Omitted):")
		fmt.Println(encodedJSON)
	}

	fmt.Println("\n--- 4. Testing Validation Failure ---")
	invalidJSON := `[
		{
			"sku": "BAD-001",
			"name": "Faulty Item",
			"price": -50.00,
			"stock": 5
		}
	]`
	_, err = DecodeProducts(invalidJSON)
	if err != nil {
		fmt.Println("Validation Error Caught:", err)
	}
}
