package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	inputs := []string{}

	fmt.Print("Inserisci delle stringhe ('esc' per uscire): ")

	for {
		fmt.Scan(&input)

		if input == "esc" {
			break
		}

		inputs = append(inputs, input)
	}

	var max int

	for _, element := range inputs {
		var cons int

		for _, char := range element {
			if strings.ContainsRune("aeiou", char) {
				continue
			}

			cons += 1
		}

		if cons >= max {
			max = cons
		}
	}

	fmt.Println(max)
}
