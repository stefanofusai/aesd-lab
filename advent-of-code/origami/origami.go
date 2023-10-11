package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

type Instruction struct {
	axis  string
	value int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	coordinates := map[Coordinate]struct{}{}
	instructions := []Instruction{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "fold along x") {
			value, _ := strconv.Atoi(strings.Split(line, "=")[1])
			instructions = append(instructions, Instruction{"x", value})
		} else if strings.Contains(line, "fold along y") {
			value, _ := strconv.Atoi(strings.Split(line, "=")[1])
			instructions = append(instructions, Instruction{"y", value})
		} else if strings.Contains(line, ",") {
			x_y := strings.Split(line, ",")
			x, _ := strconv.Atoi(x_y[0])
			y, _ := strconv.Atoi(x_y[1])
			coordinates[Coordinate{x, y}] = struct{}{}
		} else {
			continue // empty newline
		}
	}

	for _, instruction := range instructions {
		coordinates_new := map[Coordinate]struct{}{}

		for coordinate := range coordinates {
			var coord_value int
			var coordinate_new Coordinate

			is_x := instruction.axis == "x"

			if is_x {
				coord_value = coordinate.x
			} else {
				coord_value = coordinate.y
			}

			if coord_value > 2*instruction.value {
				continue
			}

			if coord_value < instruction.value {
				coordinate_new = coordinate
			} else {
				if is_x {
					coordinate_new = Coordinate{instruction.value - (coordinate.x - instruction.value), coordinate.y}
				} else {
					coordinate_new = Coordinate{coordinate.x, instruction.value - (coordinate.y - instruction.value)}
				}
			}

			coordinates_new[coordinate_new] = struct{}{}
		}

		fmt.Println(len(coordinates_new))
		coordinates = coordinates_new
	}

	fmt.Println()

	var max_x, max_y int

	for coordinate := range coordinates {
		if coordinate.x >= max_x {
			max_x = coordinate.x
		}

		if coordinate.y >= max_y {
			max_y = coordinate.y
		}
	}

	matrix := make([][]int, max_y+1)

	for i := range matrix {
		matrix[i] = make([]int, max_x+1)
	}

	for coordinate := range coordinates {
		matrix[coordinate.y][coordinate.x] = 1
	}

	for _, row := range matrix {
		for _idx, num := range row {
			if num == 0 {
				fmt.Printf(" ")
			} else {
				fmt.Printf("&")
			}

			if _idx == max_x {
				fmt.Println()
			}
		}
	}
}
