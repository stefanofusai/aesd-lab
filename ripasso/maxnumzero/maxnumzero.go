package main

import "fmt"

func main() {
	var input int
	inputs := []int{}

	fmt.Print("Inserisci dei numeri ('-1' per uscire): ")

	for {
		fmt.Scan(&input)

		if input == -1 {
			break
		}

		inputs = append(inputs, input)
	}

	var max int
	var elem int

	for _, element := range inputs {
		n := element
		var num_zeri int

		for {
			if n == 0 {
				break
			} else if n%10 == 0 {
				num_zeri++
			}

			n /= 10
		}

		if num_zeri >= max {
			max = num_zeri
			elem = element
		}
	}

	fmt.Println("num zeri", max, "element", elem)
}
