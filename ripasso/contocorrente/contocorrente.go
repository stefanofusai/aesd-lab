package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	saldo, _ := strconv.Atoi(os.Args[1])
	var spesa int

	for {
		fmt.Print("Inserisci una spesa: ")
		fmt.Scan(&spesa)
		saldo -= spesa

		if saldo <= 0 {
			break
		}
	}

	fmt.Println("Il saldo finale è:", saldo)
}
