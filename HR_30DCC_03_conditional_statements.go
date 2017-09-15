package main

import "fmt"

func main() {

	var n uint
	fmt.Scan(&n)
	if n%2 != 0 {
		fmt.Println("Weird")
	} else {
		switch {
		case n < 5:
			fmt.Println("Not Weird")
		case n > 20:
			fmt.Println("Not Weird")
		default:
			fmt.Println("Weird")
		}
	}
}
