package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(sassi(1))
	fmt.Println(sassi(2))
	fmt.Println(sassi(3))
	fmt.Println(sassi(100))
}

func sassi(height int) int {
	if height == 1 {
		return 1
	}

	return int(math.Pow(float64(height), 2)) + sassi(height-1)
}
