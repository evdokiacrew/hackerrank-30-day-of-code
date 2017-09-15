package main

import "fmt"

func main() {
	var (
		meal, tip, tax float64
	)
	fmt.Scan(&meal)
	fmt.Scan(&tip)
	fmt.Scan(&tax)
	tipCost := tip / 100 * meal
	taxCost := tax / 100 * meal
	totalCost := tipCost + taxCost + meal
	fmt.Printf("The total meal cost is %d dollars.", int(totalCost+0.5))
}
