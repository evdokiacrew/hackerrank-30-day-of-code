/*
Task
Given a base-10 integer, n, convert it to binary (base-2).
Then find and print the base-10 integer denoting the maximum number of consecutive 1's in n's binary representation.

Input Format
A single integer, n.

Output Format
Print a single base-10 integer denoting the maximum number of consecutive 1's in the binary representation of n.
*/

package main

import "fmt"

func main() {
	//reading n in STDIN
	var n, count int
	fmt.Scanf("%d", &n)
	//while loop for count maximum number of consecutive 1 with bitwise logic
	for n != 0 {
		n = (n & (n << 1))
		count++
	}
	fmt.Println(count)
}
