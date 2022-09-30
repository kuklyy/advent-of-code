package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/kuklyy/advent-of-code/common"
)

func main() {
	part1()
	part2()
}

// 323647
func part1() {
	input := make([]int, 0, 3000)

	scanner, file := common.ReadFile("./input.txt")

	defer file.Close()

	for scanner.Scan() {
		for _, strN := range strings.Split(scanner.Text(), ",") {
			n, _ := strconv.Atoi(strN)
			input = append(input, n)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(input)

	horizontalPosition := 0

	if len(input)%2 == 0 {
		horizontalPosition = (input[len(input)/2] + input[len(input)/2-1]) / 2
	} else {
		horizontalPosition = input[len(input)]
	}

	fuelUsed := 0

	for _, position := range input {
		fuelUsed += int(math.Abs(float64(position - horizontalPosition)))
	}

	fmt.Println(fuelUsed)
}

// 87640209
func part2() {
	input := make([]int, 0, 3000)

	scanner, file := common.ReadFile("./input.txt")

	defer file.Close()

	for scanner.Scan() {
		for _, strN := range strings.Split(scanner.Text(), ",") {
			n, _ := strconv.Atoi(strN)
			input = append(input, n)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(input)

	sumTable := make([]int, input[len(input)-1]+1)

	for _, v := range input {
		for j := 0; j <= input[len(input)-1]; j++ {
			sumTable[j] += fuelSum(math.Abs(float64(j - v)))
		}
	}

	sort.Ints(sumTable)
	fmt.Println(sumTable[0])

}

func fuelSum(n float64) int {
	return int(((2.0 + n - 1.0) / 2.0 * n))
}
