package main

import "fmt"

func main() {
	fmt.Println("Inserisci una sequenza di numeri terminata da -1")

	var n int
	var print bool
	var input int

	for {
		fmt.Scan(&input)

		if input == -1 {
			break
		} else if input > 100 {
			n = input
			print = true
			break
		}
	}

	if print {
		fmt.Println("Il primo numero maggiore di 100 è", n)
	} else {
		fmt.Println("Non è stato inserito nessun numero maggiore di 100")
	}
}
