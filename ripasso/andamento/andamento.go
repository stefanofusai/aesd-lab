package main

import (
	"fmt"
)

func main() {
	fmt.Println("Prego inserire una serie di interi positivi terminati da 0")

	var curr int
	var prev int

	for {
		fmt.Scan(&curr)

		if curr == 0 {
			return
		} else if curr >= prev {
			fmt.Println("+")
		} else {
			fmt.Println("-")
		}

		prev = curr
	}
}
