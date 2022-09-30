package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/kuklyy/advent-of-code/common"
)

const rowSize = 100
const tableSize = 100

func main() {
	part1()
	part2()
}

// 588
func part1() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	table := make([][]int, tableSize)

	i := 0

	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), "") {
			vN, _ := strconv.Atoi(v)
			table[i] = append(table[i], vN)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	for y, row := range table {
		for x, v := range row {
			if isLessThenNeighbors(x, y, table) {
				sum += v + 1
			}
		}
	}

	fmt.Println(sum)
}

func getDiagonalNeighbors(x, y int, table [][]int) []int {
	siblingsPos := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	siblings := make([]int, 0, 4)

	for _, pos := range siblingsPos {
		x = pos[0]
		y = pos[1]
		if x >= 0 && x < len(table[0]) && y >= 0 && y < len(table) {
			siblings = append(siblings, table[y][x])
		}
	}

	return siblings
}

func isLessThenNeighbors(x, y int, table [][]int) bool {
	siblings := getDiagonalNeighbors(x, y, table)
	element := table[y][x]

	for _, sibling := range siblings {
		if element >= sibling {
			return false
		}
	}

	return true
}

type field struct {
	value     int
	isInBasin bool
}

// 964712
func part2() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	table := make([][]field, tableSize)

	i := 0

	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), "") {
			vN, _ := strconv.Atoi(v)
			table[i] = append(table[i], field{vN, false})
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	basins := make([]map[string]bool, 0, rowSize*tableSize)

	for i := range basins {
		basins[i] = make(map[string]bool)
	}

	for y, row := range table {
		for x, field := range row {
			if field.value == 9 || field.isInBasin {
				continue
			}

			basin := make(map[string]bool)
			saveBasin(x, y, table, basin)
			basins = append(basins, basin)
		}
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	fmt.Println(len(basins[0]) * len(basins[1]) * len(basins[2]))

	// visualize(table, basins[:4])
}

//lint:ignore U1000 is not required for solution
func visualize(table [][]field, basins []map[string]bool) {
	for y, row := range table {
		for x, field := range row {
			if field.isInBasin {
				isInBasinSlice := false
				for _, basin := range basins {
					if basin[fmt.Sprint([]int{x, y})] {
						isInBasinSlice = true
						break
					}
				}
				if isInBasinSlice {
					fmt.Printf(" * ")
				} else {
					fmt.Printf(" . ")
				}
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Printf("\n")
	}
}

func saveBasin(x, y int, table [][]field, basin map[string]bool) {
	basin[fmt.Sprint([]int{x, y})] = true
	table[y][x].isInBasin = true

	neighbors := getDiagonalNeighborsPos(x, y, table)

	for _, neighbor := range neighbors {
		if !table[neighbor[1]][neighbor[0]].isInBasin {
			saveBasin(neighbor[0], neighbor[1], table, basin)
		}
	}
}

func getDiagonalNeighborsPos(x, y int, table [][]field) [][]int {
	siblingsPos := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
	siblings := make([][]int, 0, 4)

	for _, pos := range siblingsPos {
		sx := pos[0]
		sy := pos[1]
		if sx >= 0 && sx < len(table[0]) && sy >= 0 && sy < len(table) {
			if table[sy][sx].value < 9 {
				siblings = append(siblings, []int{sx, sy})
			}
		}
	}

	return siblings
}
