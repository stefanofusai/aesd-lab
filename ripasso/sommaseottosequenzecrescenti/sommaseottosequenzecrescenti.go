package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input int
	inputs := []int{}

	fmt.Print("Inserisci dei numeri ('-1' per uscire): ")

	for scanner.Scan() {
		inputs = append(inputs, input)
	}

	var sum int

	for index, element := range inputs {
		if index == len(inputs)-1 {
			break
		}

		sum += element

		if element > inputs[index+1] {
			fmt.Println("Somma finora", sum)
			sum = 0
		}
	}
}
