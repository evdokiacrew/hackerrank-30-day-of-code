package main

import "fmt"

func main() {
	var (
		n, m int
	)
	fmt.Scan(&n)
	for i := 1; i <= 10; i++ {
		m = i * n
		fmt.Println(n, "x", i, "=", m)
	}
}
