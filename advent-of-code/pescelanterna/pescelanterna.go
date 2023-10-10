package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type LanternFish struct {
	timer int
}

func spawnLanternFish(timer int) *LanternFish {
	return &LanternFish{timer: timer}
}

func (lf *LanternFish) updateTimer() {
	if lf.timer == 0 {
		lf.timer = 6
	} else {
		lf.timer--
	}
}

const DAYS = 80

func main() {
	lfs := []*LanternFish{}
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
		lfs = append(lfs, spawnLanternFish(timer))
	}

	for i := 1; i <= DAYS; i++ {
		// if i == 1 {
		// 	fmt.Printf("After 1 day: ")
		// } else {
		// 	fmt.Printf("After %d days: ", i)
		// }

		new_lfs := []*LanternFish{}

		for _, lf := range lfs {
			if lf.timer == 0 {
				new_lfs = append(new_lfs, spawnLanternFish(8))
			}

			lf.updateTimer()
		}

		lfs = append(lfs, new_lfs...)

		// for index, lf := range lfs {
		// 	if index < len(lfs)-1 {
		// 		fmt.Printf("%d,", lf.timer)
		// 	} else {
		// 		fmt.Println(lf.timer)
		// 	}
		// }

		if i == 1 {
			fmt.Printf("After 1 day there will be %d lanternfishes\n", len(lfs))
		} else {
			fmt.Printf("After %d days there will be %d lanternfishes\n", i, len(lfs))
		}
	}
}
