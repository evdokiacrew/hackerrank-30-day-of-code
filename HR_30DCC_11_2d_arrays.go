/*
Day 11: 2D Arrays
https://www.hackerrank.com/challenges/30-2d-arrays/problem

Context
Given a 6x6 2D Array, A:

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0

We define an hourglass in A to be a subset of values with indices falling in this pattern in 'As graphical representation:

a b c
  d
e f g

There are 16 hourglasses in , and an hourglass sum is the sum of an hourglass' values.

Task
Calculate the hourglass sum for every hourglass in A, then print the maximum hourglass sum.

Input Format
There are 6 lines of input, where each line contains 6 space-separated integers describing 2D Array A;
every value in will be in the inclusive range of to .

Output Format
Print the largest (maximum) hourglass sum found in A.
*/

package main

import "fmt"

func main() {
	//create a 6x6 array, read standard input, write to our array
	A := [6][6]int{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Scanf("%d", &A[i][j])
		}
	}

	//looking for the maximum value of hourglass in our array
	var maximumHourglass, currentHourglass int
	maximumHourglass = -63 //set minimum possible value of hourglass with all elements -9
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			currentHourglass = A[i][j] + A[i][j+1] + A[i][j+2] + A[i+1][j+1] + A[i+2][j] + A[i+2][j+1] + A[i+2][j+2]
			if currentHourglass > maximumHourglass {
				maximumHourglass = currentHourglass
			}
		}
	}

	//print our maximum value to STDOUT
	fmt.Println(maximumHourglass)
}
