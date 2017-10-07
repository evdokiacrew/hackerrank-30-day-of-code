package main

import "fmt"

func main() {
	//reading first string of STDIN as number of elements in map and initialize our map
	var n int
	fmt.Scanf("%d", &n)
	phoneBook := make(map[string]int)

	//reading next string of STDIN and cycle it into our map (phoneBook)
	var name string
	var phoneNumber int
	for i := 0; i < n; i++ {
		fmt.Scanf("%s%d", &name, &phoneNumber)
		phoneBook[name] = phoneNumber
	}

	//reading names, check phonebook with them and print them out
	for {
		var nameCheckEnd string
		fmt.Scanf("%s", &name)
		if name == "" {
			break
		}
		nameCheckEnd = name

		//cant solve this, need to read new names until empty line

		phoneNumber, ok := phoneBook[name]
		if ok == true {
			fmt.Printf("%s=%d\n", name, phoneNumber)
		} else {
			fmt.Println("Not found")
		}
	}
}
