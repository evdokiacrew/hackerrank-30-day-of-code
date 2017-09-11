package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var inputString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString = scanner.Text()
	fmt.Println("Hello, World.")
	fmt.Println(inputString)
}
