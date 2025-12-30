package main

import "fmt"

func main() {

	size := "Small"
	coffeeName := "Expresso"
	price := 2.60
	fmt.Println("Small Expresso price is $2.50")
	fmt.Println(size, coffeeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffeeName, price)
}
