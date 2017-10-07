package main

import "fmt"

func main() {
	//reading first string of STDIN as size of array
	var N int
	fmt.Scanf("%d", &N)
	A := make([]int, N)

	//readin next string of STDIN and cycle it into slice
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &A[i])
	}

	//reversing our slice
	for i, j := 0, len(A)-1; i < j; i, j = i+1, j-1 {
		A[i], A[j] = A[j], A[i]
	}

	//printing our slice
	for i := range A {
		fmt.Printf("%d ", A[i])
	}
}

//this time the code simpler and much faster without any packages (like bufio, os and strvconv)
