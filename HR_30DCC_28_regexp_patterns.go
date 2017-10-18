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
)

func main() {
	//reading first string in STDIN as number of rows N in the table and initialize our map
	var N int
	fmt.Scanf("%d", &N)
	emails := make(map[string]string)

	//reading next string in STDIN and cycle it into our map (emails)
	var name string
	var id string
	for i := 0; i < N; i++ {
		fmt.Scanf("%s%s", &name, &id)
		emails[name] = id
	}

	//COMPLETE!!!
	r := regexp.MustCompile("@gmail.com$")
	for i := 0; i < N; i++ {
		fmt.Println(re.FindString(emails[]))
	}
		
	}
}
