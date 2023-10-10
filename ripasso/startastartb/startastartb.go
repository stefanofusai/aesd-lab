package main

import (
	"fmt"
)

func main() {
	var input string
	inputs := []string{}

	var count_a int
	var count_b int

	for c := 0; c < 10; c++ {
		fmt.Print("Inserisci una string ('esc' per uscire): ")
		fmt.Scan(&input)

		if input == "esc" {
			break
		}

		inputs = append(inputs, input)
	}

	for _, element := range inputs {
		if element[0] == 'a' {
			count_a += 1
		} else if element[0] == 'b' {
			count_b += 1
		}
	}

	fmt.Println("# a:", count_a)
	fmt.Println("# b:", count_b)
}
