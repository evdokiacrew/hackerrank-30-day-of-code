/*
Day 29: Bitwise AND
https://www.hackerrank.com/challenges/30-bitwise-and/tutorial

Task
Given set S={1,2,3,...,N}. Find two integers, A and B (where A < B),
from set S such that the value of A&B is the maximum possible and also less than a given integer, K.
In this case, & represents the bitwise AND operator.

Input Format
The first line contains an integer, T, the number of test cases.
Each of the T subsequent lines defines a test case as 2 space-separated integers, N and K, respectively.

Output Format
For each test case, print the maximum possible value of A&B on a new line.
*/

package main

import "fmt"

func main() {
	//reading first string in STDIN as the number of test cases T
	var T, N, K int
	fmt.Scanf("%d", &T)

	//loop T times and print out to STDOUT results of our function
	for i := 0; i < T; i++ {
		fmt.Scanf("%d%d", &N, &K)
		fmt.Println(maximumBitwiseAnd(N, K))
	}
}

//function to find out maximum bitwise of A and B  and lesser then K
func maximumBitwiseAnd(N int, K int) int {
	var bitwiseCurrent, bitwiseMaximum int
	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			bitwiseCurrent = i & j
			if bitwiseCurrent < K && bitwiseCurrent > bitwiseMaximum {
				bitwiseMaximum = bitwiseCurrent
			}
		}
	}
	return bitwiseMaximum
}
