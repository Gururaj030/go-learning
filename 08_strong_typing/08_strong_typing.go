package main

import "fmt"

func main() {

	// Price for one cup of coffee
	price := 5.0

	// cups sold in one day
	quantity := 20

	total := price * float64(quantity)

	fmt.Printf("Total Income per day: %.2f\n", total)

}
