/*
Day 28: RegEx, Patterns, and Intro to Databases
https://www.hackerrank.com/challenges/30-regex-patterns/problem

Task
Consider a database table, Emails, which has the attributes First Name and Email ID.
Given N rows of data simulating the Emails table, print an alphabetically-ordered list of people whose email address ends in @gmail.com.

Input Format
The first line contains an integer, N, total number of rows in the table.
Each of the N subsequent lines contains 2 space-separated strings denoting a person's first name and email ID, respectively.

Output Format
Print an alphabetically-ordered list of first names for every user with a gmail account. Each name must be printed on a new line.
*/

package main

import (
	"fmt"
	"regexp"
	"sort"
)

func main() {
	//reading first string in STDIN as number of rows N in the table
	var N int
	fmt.Scanf("%d", &N)

	//loop reading strings in STDIN and append to our slice if regexp true
	var name, id string
	emails := make([]string, 0)
	reg := regexp.MustCompile("@gmail.com$")
	for i := 0; i < N; i++ {
		fmt.Scanf("%s%s", &name, &id)
		if reg.MatchString(id) {
			emails = append(emails, name)
		}
	}

	//sort our slice with '@gmail.com' names
	sort.Strings(emails)

	//loop printing our slice to STDOUT
	for i := range emails {
		fmt.Println(emails[i])
	}
}
