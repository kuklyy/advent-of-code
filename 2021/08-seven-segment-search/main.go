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

// 525
func part1() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	oneLen := 2
	fourLen := 4
	sevenLen := 3
	eightLen := 7
	result := 0

	for scanner.Scan() {
		for _, v := range strings.Split(strings.Split(scanner.Text(), " | ")[1], " ") {
			vLen := len(v)
			if vLen == oneLen || vLen == fourLen || vLen == sevenLen || vLen == eightLen {
				result += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

// 1083859
func part2() {
	scanner, input := common.ReadFile("input.txt")
	defer input.Close()

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		parsedLine := strings.Split(line, " | ")

		segmentData := make(map[int]map[string]int)

		for i := 0; i < 7; i++ {
			segmentData[i] = make(map[string]int)
		}

		segmentMap := make(map[string]int)

		// parse input
		for _, input := range strings.Fields(parsedLine[0]) {
			inputLen := len(input)

			switch inputLen {

			case 6: // 0, 6, 9
				for _, char := range input {
					// 0
					segmentData[0][string(char)]++
					segmentData[1][string(char)]++
					segmentData[2][string(char)]++
					segmentData[4][string(char)]++
					segmentData[5][string(char)]++
					segmentData[6][string(char)]++

					// 6
					segmentData[0][string(char)]++
					segmentData[1][string(char)]++
					segmentData[3][string(char)]++
					segmentData[4][string(char)]++
					segmentData[5][string(char)]++
					segmentData[6][string(char)]++

					// 9
					segmentData[0][string(char)]++
					segmentData[1][string(char)]++
					segmentData[2][string(char)]++
					segmentData[3][string(char)]++
					segmentData[5][string(char)]++
					segmentData[6][string(char)]++
				}
			case 2: // 1
				for _, char := range input {
					segmentData[2][string(char)]++
					segmentData[5][string(char)]++
				}
			case 5: // 2, 3, 5
				for _, char := range input {
					// 2
					segmentData[0][string(char)]++
					segmentData[2][string(char)]++
					segmentData[3][string(char)]++
					segmentData[4][string(char)]++
					segmentData[6][string(char)]++

					// 3
					segmentData[0][string(char)]++
					segmentData[2][string(char)]++
					segmentData[3][string(char)]++
					segmentData[5][string(char)]++
					segmentData[6][string(char)]++

					// 5
					segmentData[0][string(char)]++
					segmentData[1][string(char)]++
					segmentData[3][string(char)]++
					segmentData[5][string(char)]++
					segmentData[6][string(char)]++
				}
			case 4: // 4
				for _, char := range input {
					segmentData[1][string(char)]++
					segmentData[2][string(char)]++
					segmentData[3][string(char)]++
					segmentData[5][string(char)]++
				}
			case 3: // 7
				for _, char := range input {
					segmentData[0][string(char)]++
					segmentData[2][string(char)]++
					segmentData[5][string(char)]++
				}
			case 7: // 8
				for _, char := range input {
					segmentData[0][string(char)]++
					segmentData[1][string(char)]++
					segmentData[2][string(char)]++
					segmentData[3][string(char)]++
					segmentData[4][string(char)]++
					segmentData[5][string(char)]++
					segmentData[6][string(char)]++
				}
			}

		}

		// parse output
		for i := 0; i < 7; i++ {
			maxValue := 0
			var maxValueKey string
			maxValueIndex := math.MinInt

			for j, s := range segmentData {
				for c, k := range s {
					if k > maxValue {
						maxValue = k
						maxValueKey = c
						maxValueIndex = j
					}
				}
			}

			segmentMap[maxValueKey] = maxValueIndex

			delete(segmentData, maxValueIndex)

			for _, k := range segmentData {
				delete(k, maxValueKey)
			}
		}

		digitalOrigin := make([][]int, 10)
		digitalOrigin[0] = []int{0, 1, 2, 4, 5, 6}
		digitalOrigin[1] = []int{2, 5}
		digitalOrigin[2] = []int{0, 2, 3, 4, 6}
		digitalOrigin[3] = []int{0, 2, 3, 5, 6}
		digitalOrigin[4] = []int{1, 2, 3, 5}
		digitalOrigin[5] = []int{0, 1, 3, 5, 6}
		digitalOrigin[6] = []int{0, 1, 3, 4, 5, 6}
		digitalOrigin[7] = []int{0, 2, 5}
		digitalOrigin[8] = []int{0, 1, 2, 3, 4, 5, 6}
		digitalOrigin[9] = []int{0, 1, 2, 3, 5, 6}

		var sb strings.Builder

		for _, f := range strings.Fields(parsedLine[1]) {
			signals := make([]int, 0, 7)
			for _, char := range f {
				signals = append(signals, segmentMap[string(char)])
			}
			sort.Ints(signals)

			for i := 0; i < len(digitalOrigin); i++ {
				digitMatch := compareSlices(signals, digitalOrigin[i])
				if digitMatch {
					sb.WriteString(fmt.Sprint(i))
					break
				}
			}

		}

		rowInt, _ := strconv.Atoi(sb.String())
		sum += rowInt
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)

}

func compareSlices(signals []int, digital []int) bool {
	if len(signals) != len(digital) {
		return false
	}

	for i := 0; i < len(signals); i++ {
		if signals[i] != digital[i] {
			return false
		}
	}

	return true
}
