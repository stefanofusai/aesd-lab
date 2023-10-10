package main

import (
	"fmt"
	"sort"
)

func main() {
	const N = 10
	numeri := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&numeri[i])
	}

	pariDispari(numeri)
	fmt.Println("Il minimo dispari è", minDispari(numeri))
}

func stranoProdotto(numeri []int) int {
	var prod int = 1

	for _, n := range numeri {
		if n > 7 && n%4 == 0 {
			prod *= n
		}
	}

	return prod
}

func pariDispari(numeri []int) {
	for _, n := range numeri {
		if n%2 == 0 {
			fmt.Printf("Il numero %v è pari\n", n)
		} else {
			fmt.Printf("Il numero %v è dispari\n", n)
		}
	}
}

func minDispari(numeri []int) int {
	sort.Ints(numeri)

	for _, n := range numeri {
		if n%2 != 0 {
			return n
		}
	}

	return -1
}
