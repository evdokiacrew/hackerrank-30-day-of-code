/*
Day 27: Testing
https://www.hackerrank.com/challenges/30-testing/problem

Context
Consider the following problem (but do not solve it):

    Problem Statement

    A Discrete Mathematics professor has a class of n students.
    Frustrated with their lack of discipline, the professor decides to cancel class if fewer than k students are present when class starts.
    Given the arrival time of each student, determine if the class is canceled.

    Input Format
	The first line of input contains t, the number of lectures.
    The information for each lecture spans two lines. The first line has two space-separated integers,
    n (the number of students in the class) and k (the cancelation threshold).
    The second line contains n space-separated integers describing the array of students' arrival times (A = a(0), a(1), ... a(n-1)).

    Note: Non-positive arrival times (a(i) <= 0) indicate the student arrived early or on time;
    positive arrival times (a(i) > 0) indicate the student arrived a(i) minutes late.
    If a student arrives exactly on time (a(i)=0), the student is considered to have entered before the class started.

    Output Format
    For each test case, print the word YES if the class is canceled or NO if it is not.

    Example
    When properly solved, this input:

    2
    4 3
    -1 -3 4 2
    4 2
    0 -1 2 1

    Produces this output:

    YES
    NO

    For the first test case, k=3. The professor wants at least 3 students in attendance, but only 2 arrive on time ( and ).
    Thus, the class is canceled.
    For the second test case, k=2. The professor wants at least 2 students in attendance, and 2 do arrive on time ( and ).
    Thus, the class is not canceled.

Task

Create and print a test case for the problem above that meet the following criteria:
	t <= 5
	3 <= n <= 200
	1 <= k <= n
	-10^3 <= a(i) <= 10^3 where i depend [1,n]
    The value of n must be distinct across all the test cases.
    Array A must have at least 1 zero, 1 positive integer, 1 and negative integer.

Input Format
You are not responsible for reading anything from stdin.

Output Format
Print 11 lines of output that can be read by the Professor challenge as valid input.
Your test case must result in the following output when fed as input to the Professor challenge's solution:

YES
NO
YES
NO
YES

*/

package main

import "fmt"

func main() {
	//printing to STDOUT our test cases
	fmt.Println(5)
	fmt.Println("3 3")
	fmt.Println("-1 0 1")
	fmt.Println("4 1")
	fmt.Println("-1 0 1 2")
	fmt.Println("5 5")
	fmt.Println("-1 0 1 2 3")
	fmt.Println("6 1")
	fmt.Println("-1 0 1 2 3 4")
	fmt.Println("7 7")
	fmt.Println("-1 0 1 2 3 4 5")
}
