package main

import "fmt"

func main() {
	var cupsQty int = 3

	// cupsQty = "four" // X Compile-time error
	fmt.Println("Number of cupts:", cupsQty)
}
