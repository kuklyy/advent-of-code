package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

// 5585
func part1() {
	const lineWidth = 1000
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	matrix := make([][]int, lineWidth)

	for i := range matrix {
		matrix[i] = make([]int, lineWidth)
	}

	for scanner.Scan() {
		line := scanner.Text()
		pointsPair := strings.Split(line, " -> ")

		start := strings.Split(pointsPair[0], ",")
		end := strings.Split(pointsPair[1], ",")

		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])

		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

		if x1 == x2 || y1 == y2 {
			for y, row := range matrix {
				for x := range row {

					if (y-y1)*(x2-x1)-(x-x1)*(y2-y1) == 0 {
						if y <= maxInt(y1, y2) && y >= minInt(y1, y2) && x >= minInt(x1, x2) && x <= maxInt(x1, x2) {
							matrix[x][y] = matrix[x][y] + 1
						}
					}
				}
			}
		}

	}

	result := 0

	for _, row := range matrix {
		for _, column := range row {
			if column >= 2 {
				result++
			}
		}
	}

	fmt.Println(result)
}

// 17193
func part2() {
	const lineWidth = 1000
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	matrix := make([][]int, lineWidth)

	for i := range matrix {
		matrix[i] = make([]int, lineWidth)
	}

	for scanner.Scan() {
		line := scanner.Text()
		pointsPair := strings.Split(line, " -> ")

		start := strings.Split(pointsPair[0], ",")
		end := strings.Split(pointsPair[1], ",")

		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])

		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

		for y, row := range matrix {
			for x := range row {

				if (y-y1)*(x2-x1)-(x-x1)*(y2-y1) == 0 {
					if y <= maxInt(y1, y2) && y >= minInt(y1, y2) && x >= minInt(x1, x2) && x <= maxInt(x1, x2) {
						matrix[x][y] = matrix[x][y] + 1
					}
				}
			}
		}

	}

	result := 0

	for _, row := range matrix {
		for _, column := range row {
			if column >= 2 {
				result++
			}
		}
	}

	fmt.Println(result)
}

func minInt(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}

	return y
}
