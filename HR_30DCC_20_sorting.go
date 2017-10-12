/*
Day 20: Sorting
https://www.hackerrank.com/challenges/30-sorting/problem

Consider the following version of Bubble Sort:

for (int i = 0; i < n; i++) {
    // Track number of elements swapped during a single array traversal
    int numberOfSwaps = 0;

    for (int j = 0; j < n - 1; j++) {
        // Swap adjacent elements if they are in decreasing order
        if (a[j] > a[j + 1]) {
            swap(a[j], a[j + 1]);
            numberOfSwaps++;
        }
    }

    // If no elements were swapped during a traversal, array is sorted
    if (numberOfSwaps == 0) {
        break;
    }
}

Task
Given an array, a, of size n distinct elements, sort the array in ascending order using the Bubble Sort algorithm above.
Once sorted, print the following lines:

    Array is sorted in numSwaps swaps.
    where numSwaps is the number of swaps that took place.

    First Element: firstElement
    where firstElement is the first element in the sorted array.

    Last Element: lastElement
    where lastElement is the last element in the sorted array.

Hint: To complete this challenge, you will need to add a variable that keeps a running tally of all swaps that occur during execution.

Input Format
There are 6 lines of input, where each line contains 6 space-separated integers describing 2D Array A;
every value in will be in the inclusive range of to .

Output Format
Print the following three lines of output:

    Array is sorted in numSwaps swaps.
    where numSwaps is the number of swaps that took place.

    First Element: firstElement
    where firstElement is the first element in the sorted array.

    Last Element: lastElement
    where lastElement is the last element in the sorted array.
*/

package main

import "fmt"

func main() {
	//reading first string of STDIN as size of array
	var n, numberOfSwaps int
	fmt.Scanf("%d", &n)
	a := make([]int, n)

	//readin next string of STDIN and cycle it into slice
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}

	//lets sort
	for i := 0; i < n; i++ {
		// Track number of elements swapped during a single array traversal
		//var numberOfSwaps int
		for j := 0; j < (n - 1); j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				numberOfSwaps++
			}
		}
		// If no elements were swapped during a traversal, array is sorted
		if numberOfSwaps == 0 {
			break
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", numberOfSwaps)
	fmt.Printf("First Element: %d\n", a[0])
	fmt.Printf("Last Element: %d\n", a[n-1])
}
