package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.

	var integer uint64
	var double float64
	var String string

	// Read and save an integer, double, and String to your variables.

	scanner.Scan()
	integer, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	double, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	String = scanner.Text()

	// Print the sum of both integer variables on a new line.

	fmt.Println(i + integer)

	// Print the sum of the double variables on a new line.

	fmt.Printf("%.1f \n", double+d)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

	fmt.Println(s + String)

}
