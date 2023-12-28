package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter word: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("What")
	}

	fmt.Printf("Word contains %d letters\n", len(input))
	for _, letter := range input {
		fmt.Println(string(letter))
	}

}
