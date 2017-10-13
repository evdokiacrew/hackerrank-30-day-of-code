/*
Day 25: Running Time and Complexity
https://www.hackerrank.com/challenges/30-running-time-and-complexity/problem

Task
A prime is a natural number greater than 1 that has no positive divisors other than 1 and itself.
Given a number, n, determine and print whether it's Prime or Not prime.

Note: If possible, try to come up with a O(sqrt(n)) primality algorithm, or see what sort of optimizations you come up with for an O(n) algorithm.
Be sure to check out the Editorial after submitting your code!

Input Format
The first line contains an integer, T, the number of test cases.
Each of the T subsequent lines contains an integer, n, to be tested for primality.

Output Format
For each test case, print whether n is Prime or Not prime on a new line.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	//reading first string of STDIN as number of test cases T
	var T, n int
	fmt.Scanf("%d", &T)

	//loop through the number of test cases T and print out result to STDOUT
	for i := 0; i < T; i++ {
		fmt.Scanf("%d", &n)
		fmt.Println(primeOrNot(n))
	}
}

//function of determining whether n is Prime number or not
func primeOrNot(n int) string {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return "Not prime"
		}
	}
	return "Prime"
}
