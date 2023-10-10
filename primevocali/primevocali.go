package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	inputs := []string{}

	fmt.Print("Metti una stringa brutto coglione ('esc' per uscire): ")

	for {
		fmt.Scan(&input)

		if input == "esc" {
			break
		}

		inputs = append(inputs, input)
	}

	for _, element := range inputs {
		index := strings.IndexAny(element, "aeiou")

		if index == -1 {
			fmt.Println("La parola non contiene vocali")
		} else {
			fmt.Println("La prima vocale è: ", string(element[index]))
		}
	}
}
