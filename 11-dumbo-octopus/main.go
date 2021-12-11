package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kuklyy/advent-of-code-2021/common"
)

const tableSize = 10

func main() {
	part1()
	part2()
}

// 1632
func part1() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	table := make([][]int, tableSize)

	i := 0
	for scanner.Scan() {
		for _, c := range strings.Split(scanner.Text(), "") {
			n, _ := strconv.Atoi(c)
			table[i] = append(table[i], n)
		}
		i += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	score := 0

	for d := 0; d < 100; d++ {
		flashMap := make(map[string]bool)

		for y, row := range table {
			for x := range row {
				calculateStepFlashes(x, y, table, flashMap)
			}
		}

		score += len(flashMap)
	}

	fmt.Println(score)

}

// 303
func part2() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	table := make([][]int, tableSize)

	i := 0
	for scanner.Scan() {
		for _, c := range strings.Split(scanner.Text(), "") {
			n, _ := strconv.Atoi(c)
			table[i] = append(table[i], n)
		}
		i += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	step := 0

	for {
		flashMap := make(map[string]bool)

		for y, row := range table {
			for x := range row {
				calculateStepFlashes(x, y, table, flashMap)
			}
		}

		step += 1

		if len(flashMap) == tableSize*tableSize {
			break
		}
	}

	fmt.Println(step)

}

func calculateStepFlashes(x, y int, table [][]int, flashMap map[string]bool) {
	if flashMap[getMapKey(x, y)] {
		return
	}

	table[y][x] += 1

	if table[y][x] == 10 {
		table[y][x] = 0
		flashMap[getMapKey(x, y)] = true

		for _, neighbor := range getNeighbors(x, y, table, flashMap) {
			calculateStepFlashes(neighbor[0], neighbor[1], table, flashMap)
		}
	}

}

func getNeighbors(x, y int, table [][]int, flashMap map[string]bool) [][]int {
	neighborsPos := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}, {x - 1, y + 1}, {x + 1, y + 1}, {x - 1, y - 1}, {x + 1, y - 1}}
	neighbors := make([][]int, 0, len(neighborsPos))

	for _, pos := range neighborsPos {
		nx := pos[0]
		ny := pos[1]
		if nx >= 0 && nx < tableSize && ny >= 0 && ny < tableSize {
			if !flashMap[getMapKey(nx, ny)] {
				neighbors = append(neighbors, []int{nx, ny})
			}
		}
	}

	return neighbors
}

func getMapKey(x, y int) string {
	return fmt.Sprint([]int{x, y})
}
