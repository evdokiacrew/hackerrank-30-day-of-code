package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var S string

	fmt.Scan(&t)
	for t >= 0; t-- {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	S = scanner.Text()
	fmt.Println(t, S)
	}
	// теперь надо понять как все это тестировать, сижу разбираюсь
}
