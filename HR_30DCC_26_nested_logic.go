/*
Day 26: Nested Logic
https://www.hackerrank.com/challenges/30-nested-logic/problem

Task
Your local library needs your help!
Given the expected and actual return dates for a library book, create a program that calculates the fine (if any).
The fee structure is as follows:

    If the book is returned on or before the expected return date, no fine will be charged
    	i.e.: fine = 0.
    If the book is returned after the expected return day but still within the same calendar month and year as the expected return date,
    	fine = 15 Hackos x (the number of days late).
    If the book is returned after the expected return month but still within the same calendar year as the expected return date,
    	fine = 500 Hackos x (the number of months late).
    If the book is returned after the calendar year in which it was expected,
    	there is a fixed fine of 10000 Hackos.

Input Format
The first line contains 3 space-separated integers denoting the respective
day, month, and year on which the book was actually returned.
The second line contains 3 space-separated integers denoting the respective
day, month, and yearon which the book was expected to be returned (due date).

Output Format
Print a single integer denoting the library fine for the book received as input.
*/

package main

import "fmt"

func main() {
	//reading actually and expected dates frin STDIN
	var fine, dayActually, dayExpexted, monthActually, monthExpected, yearActually, yearExpected int
	fmt.Scanf("%d", &dayActually)
	fmt.Scanf("%d", &monthActually)
	fmt.Scanf("%d", &yearActually)
	fmt.Scanf("%d", &dayExpexted)
	fmt.Scanf("%d", &monthExpected)
	fmt.Scanf("%d", &yearExpected)

	//nested 'magic' (logic)
	if yearActually > yearExpected {
		fine = 10000
	}
	if yearActually == yearExpected {
		if monthActually > monthExpected {
			fine = (monthActually - monthExpected) * 500
		}
		if monthActually == monthExpected {
			if dayActually > dayExpexted {
				fine = (dayActually - dayExpexted) * 15
			}
			if dayActually == dayExpexted {
				fine = 0
			}
		}
	}
	fmt.Println(fine)
}
