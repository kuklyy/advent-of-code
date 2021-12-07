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

// 387413
func part1() {

	scanner, file := common.ReadFile("./input.txt")

	fishFamily := make([]int, 0, 3000)

	var input string

	for scanner.Scan() {
		input = scanner.Text()
	}

	for _, fish := range strings.Split(input, ",") {
		fishInterval, _ := strconv.Atoi(fish)
		fishFamily = append(fishFamily, fishInterval)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	days := 80
	newbornFishInterval := 8
	fishInterval := 6

	for i := 0; i < days; i++ {
		newbornFish := make([]int, 0, len(fishFamily))

		for i := range fishFamily {
			fishFamily[i] = fishFamily[i] - 1
			if fishFamily[i] == -1 {
				newbornFish = append(newbornFish, newbornFishInterval)
				fishFamily[i] = fishInterval
			}
		}

		fishFamily = append(fishFamily, newbornFish...)
	}

	fmt.Println(len(fishFamily))
}

// 1738377086345
func part2() {
	scanner, file := common.ReadFile("./input.txt")

	defer file.Close()

	table := make([]int, 9)

	var input string

	for scanner.Scan() {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, strV := range strings.Split(input, ",") {
		v, _ := strconv.Atoi(strV)
		table[v]++
	}

	for i := 0; i < 256; i++ {
		dayTable := make([]int, 9)

		for j := len(table) - 1; j >= 0; j-- {

			if j == 0 {
				dayTable[len(table)-1] = table[0]
				dayTable[6] = dayTable[6] + table[0]
				continue
			}
			dayTable[j-1] = table[j]
		}
		table = dayTable
	}

	sum := 0

	for _, v := range table {
		sum += v
	}

	fmt.Println(sum)
}
