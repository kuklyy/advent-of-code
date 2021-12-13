package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kuklyy/advent-of-code-2021/common"
)

func main() {
	part1()
	part2()
}

const tableSize = 1600

type FoldCommand struct {
	axis string
	i    int
}

// 687
func part1() {
	scanner, input := common.ReadFile("input.txt")
	defer input.Close()

	readPos := true

	table := make([][]int, tableSize)
	for i := range table {
		table[i] = make([]int, tableSize)
	}

	foldCommands := make([]FoldCommand, 0, 16)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readPos = false
			continue
		}

		if readPos {
			pos := strings.Split(line, ",")
			x, _ := strconv.Atoi(pos[0])
			y, _ := strconv.Atoi(pos[1])
			table[y][x] += 1
		} else {
			command := strings.Split(strings.Fields(line)[2], "=")
			i, _ := strconv.Atoi(command[1])
			foldCommands = append(foldCommands, FoldCommand{axis: command[0], i: i})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	xLen := tableSize
	yLen := tableSize
	ans := 0

	for _, v := range foldCommands[:1] {
		if v.axis == "y" {
			foldPaperY(table, v.i, yLen)
			yLen = v.i
		} else {
			foldPaperX(table, v.i, xLen)
			xLen = v.i
		}
	}

	for _, row := range table {
		for _, v := range row {
			if v > 0 {
				ans += 1
			}
		}
	}

	fmt.Println(ans)
}

// FGKCKBZG
func part2() {
	scanner, input := common.ReadFile("input.txt")
	defer input.Close()

	readPos := true

	table := make([][]int, tableSize)
	for i := range table {
		table[i] = make([]int, tableSize)
	}

	foldCommands := make([]FoldCommand, 0, 16)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readPos = false
			continue
		}

		if readPos {
			pos := strings.Split(line, ",")
			x, _ := strconv.Atoi(pos[0])
			y, _ := strconv.Atoi(pos[1])
			table[y][x] += 1
		} else {
			command := strings.Split(strings.Fields(line)[2], "=")
			i, _ := strconv.Atoi(command[1])
			foldCommands = append(foldCommands, FoldCommand{axis: command[0], i: i})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	xLen := tableSize
	yLen := tableSize

	for _, v := range foldCommands {
		if v.axis == "y" {
			foldPaperY(table, v.i, yLen)
			yLen = v.i
		} else {
			foldPaperX(table, v.i, xLen)
			xLen = v.i
		}
	}

	// wow you just print table to see the results
	// btw x*y of marked positions gives exactly 8 results in range of uppercase letters in octal xd
	for y, row := range table {
		if y >= yLen {
			break
		}
		for x, v := range row {
			if x >= xLen {
				break
			}
			if v > 0 {
				fmt.Printf(" * ")
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Printf("\n")
	}
}

func foldPaperY(table [][]int, i int, len int) {
	for x := 0; x < tableSize; x++ {
		mirrorBase := 2
		for y := 0; y < tableSize; y++ {
			if y > i {
				mirrorPos := y - mirrorBase
				if mirrorPos >= 0 {
					table[mirrorPos][x] += table[y][x]
				}
				table[y][x] = 0
				mirrorBase += 2
			} else if x == i {
				table[y][x] = 0
			}
		}
	}
}

func foldPaperX(table [][]int, i int, len int) {
	for y, row := range table {
		mirrorBase := 2
		for x := range row {
			if x > i {
				mirrorPos := x - mirrorBase
				if mirrorPos >= 0 {
					table[y][mirrorPos] += table[y][x]
				}
				table[y][x] = 0
				mirrorBase += 2
			} else if x == i {
				table[y][x] = 0
			}
		}
	}
}
