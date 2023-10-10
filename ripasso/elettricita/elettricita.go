package main

import (
	"fmt"
)

func main() {
	var input int
	inputs := []int{}

	var input_count int
	fmt.Print("Inserisci almeno tre stringhe ('-1' per uscire): ")

	for {
		fmt.Scan(&input)
		input_count += 1

		if input == -1 {
			if input_count > 3 {
				break
			} else {
				fmt.Println("You can't break yet, you fool!")
				input_count -= 1
				continue
			}
		}

		inputs = append(inputs, input)
	}

	for index, element := range inputs {
		if index == 0 {
			continue
		} else if index == len(inputs)-1 {
			break
		}

		if inputs[index-1] > element && element < inputs[index+1] {
			fmt.Println(element)
		}
	}

}
