/*
Task
Given a base-10 integer, n, convert it to binary (base-2). Then find and print the base-10 integer denoting the maximum number of consecutive 1's in n's binary representation.

Input Format
A single integer, n.

Output Format
Print a single base-10 integer denoting the maximum number of consecutive 1's in the binary representation of n.
*/

package main

import "fmt"

func main() {
	//reading n of STDIN
	var n int
	fmt.Scanf("%d", &n)
	//printing n binary into another variable
	nBinary := fmt.Printf("%b", n)
	fmt.Println(nBinary)
}
