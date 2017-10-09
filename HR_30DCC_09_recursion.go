package main

import "fmt"

//function to compute factorial, our function get named "factorial" to get scores
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	//reading N of STDIN
	var N int
	fmt.Scanf("%d", &N)
	//call our function and print result in STDOUT
	fmt.Println(factorial(N))
}
