package main

import "fmt"

func main() {
	name := "Americano"
	price := 2.99
	ready := true
	count := 6

	fmt.Printf("Type of the name is: %T\n", name)
	fmt.Printf("Type of the price is: %T\n", price)
	fmt.Printf("Type of the ready is: %T\n", ready)
	fmt.Printf("Type of the count is: %T\n", count)
}
