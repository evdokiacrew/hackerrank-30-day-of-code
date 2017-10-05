package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func splitOddLetters(S string) {
	evenString := make([]byte, 0, 64)
	oddString := make([]byte, 0, 64)
	for i := 0; i < len(S); i++ {
		if i%2 != 0 {
			oddString = append(oddString, S[i])
		} else {
			evenString = append(evenString, S[i])
		}
	}
	fmt.Println(string(evenString), string(oddString))
}

func main() {
	var T int
	var S string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		S = scanner.Text()
		splitOddLetters(S)
	}
}
