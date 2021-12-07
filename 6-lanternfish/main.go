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
		fmt.Println(i)
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
