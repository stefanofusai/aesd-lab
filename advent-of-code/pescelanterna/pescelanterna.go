package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DAYS = 256

func main() {
	lfs := map[int]int{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Initial state: ")

	scanned := scanner.Scan()
	if !scanned {
		return
	}

	input := scanner.Text()
	numbers := strings.Split(input, ",")

	for _, number := range numbers {
		timer, _ := strconv.Atoi(number)
		lfs[timer]++
	}

	for i := 0; i < DAYS; i++ {
		zero_count := lfs[0]

		for i := 0; i < 8; i++ {
			lfs[i] = lfs[i+1]
		}

		lfs[8] = zero_count
		lfs[6] += zero_count
	}

	var total_lfs int

	for _, counts := range lfs {
		total_lfs += counts
	}

	fmt.Println(total_lfs)
}
