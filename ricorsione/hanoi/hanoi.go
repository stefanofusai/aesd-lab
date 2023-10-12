package main

import "fmt"

func main() {
	for _, mv := range []int{1, 2, 3} {
		fmt.Printf("Solution for %d moves:\n", mv)
		hanoi(mv, 'A', 'B', 'T')
		fmt.Print("------ end ------\n\n")
	}
}

func hanoi(n int, from rune, to rune, temp rune) {
	if n == 1 {
		fmt.Printf("----- %c > %c -----\n", from, to)
		return
	}

	hanoi(n-1, from, temp, to)
	fmt.Printf("----- %c > %c -----\n", from, to)
	hanoi(n-1, temp, to, from)
}
