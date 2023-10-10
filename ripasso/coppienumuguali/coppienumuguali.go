package main

import "fmt"

func main() {
	var input int
	inputs := []int{}

	fmt.Print("Inserisci dei numeri ('0' per uscire): ")

	for {
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		inputs = append(inputs, input)
	}

	var c int

	for index, element := range inputs {
		if index+1 == len(inputs) {
			break
		} else if element == inputs[index+1] {
			c += 1
		}
	}

	fmt.Println(c)
}
