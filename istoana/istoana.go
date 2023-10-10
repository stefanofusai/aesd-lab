package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Gentilissimo utente, prego inserisci una riga di caratteri:")

	scanner.Scan()
	input := scanner.Text()

	fmt.Println("E adesso inserisci un carattere:")

	scanner.Scan()
	char := scanner.Text()
	fmt.Println("Il carattere è contenuto", strings.Count(input, char), "volte")

	input_uppercase := strings.ToUpper(input)

	for c := 'A'; c <= 'Z'; c++ {
		letter := string(c)
		count := strings.Count(input_uppercase, letter)

		if count > 0 {
			fmt.Println(letter, " ", strings.Repeat("*", count))
		}
	}

	fmt.Println("Vai con 2!")

	scanner.Scan()
	input_2 := scanner.Text()
	input_2_uppercase := strings.ToUpper(input_2)

	var is_anag bool = true

	for c := 'A'; c <= 'Z'; c++ {
		letter := string(c)
		count_1 := strings.Count(input_uppercase, letter)
		count_2 := strings.Count(input_2_uppercase, letter)

		if count_1 != count_2 {
			is_anag = false
			break
		}
	}

	if is_anag {
		fmt.Println("ciao (è un anagramma)")
	} else {
		fmt.Println("non ciao (non è un anagramma)")
	}
}
