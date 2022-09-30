package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// these should be done by bitwise operators, but i want to treat it as a random data structure
// TODO implement alternate solution using bitwise operators
func main() {
	part1()
	part2()
}

// 2250414
func part1() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	const inputLen = 12

	var commonBits [inputLen]int

	for scanner.Scan() {

		for i, char := range scanner.Text() {
			number, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}

			if number == 0 {
				commonBits[i] -= 1
			} else {
				commonBits[i] += 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var gammaRateBin [inputLen]string
	var epsilonRateBin [inputLen]string

	for i, value := range commonBits {
		if value < 0 {
			gammaRateBin[i] = "0"
			epsilonRateBin[i] = "1"
		} else {
			gammaRateBin[i] = "1"
			epsilonRateBin[i] = "0"
		}
	}

	gammaRate, err := strconv.ParseInt(strings.Join(gammaRateBin[:], ""), 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilonRate, err := strconv.ParseInt(strings.Join(epsilonRateBin[:], ""), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(gammaRate * epsilonRate)
}

// 6085575
func part2() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	data := make([]string, 0, 3000)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	oxygenCandidates := make([]string, 0, 3000)
	oxygenCandidates = append(oxygenCandidates, data...)

	charIndex := 0

	for len(oxygenCandidates) > 1 {
		ones := make([]string, 0, len(oxygenCandidates))
		zeros := make([]string, 0, len(oxygenCandidates))

		for _, row := range oxygenCandidates {
			number, err := strconv.Atoi(string(row[charIndex]))
			if err != nil {
				log.Fatal(err)
			}

			if number == 0 {
				zeros = append(zeros, row)
			} else {
				ones = append(ones, row)
			}
		}

		zerosLen := len(zeros)
		onesLen := len(ones)

		if zerosLen == onesLen {
			oxygenCandidates = removeSlice(oxygenCandidates, zeros)
		} else if zerosLen > onesLen {
			oxygenCandidates = removeSlice(oxygenCandidates, ones)
		} else {
			oxygenCandidates = removeSlice(oxygenCandidates, zeros)
		}
		charIndex++
	}

	co2Candidates := make([]string, 0, 3000)
	co2Candidates = append(co2Candidates, data...)
	charIndex = 0

	for len(co2Candidates) > 1 {
		ones := make([]string, 0, len(co2Candidates))
		zeros := make([]string, 0, len(co2Candidates))

		for _, row := range co2Candidates {
			number, err := strconv.Atoi(string(row[charIndex]))
			if err != nil {
				log.Fatal(err)
			}

			if number == 0 {
				zeros = append(zeros, row)
			} else {
				ones = append(ones, row)
			}
		}

		zerosLen := len(zeros)
		onesLen := len(ones)

		if zerosLen == onesLen {
			co2Candidates = removeSlice(co2Candidates, ones)
		} else if zerosLen > onesLen {
			co2Candidates = removeSlice(co2Candidates, zeros)
		} else {
			co2Candidates = removeSlice(co2Candidates, ones)
		}
		charIndex++
	}

	co2Rating, err := strconv.ParseInt(co2Candidates[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	oxygenRating, err := strconv.ParseInt(oxygenCandidates[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(co2Rating * oxygenRating)
}

func removeSlice(s []string, r []string) []string {
	for _, value := range r {
		s = removeElement(s, value)
	}

	return s
}

func removeElement(s []string, r string) []string {
	index := findIndex(s, r)

	if index < 0 {
		return s
	} else {
		return append(s[:index], s[index+1:]...)
	}
}

func findIndex(s []string, v string) int {
	for i, n := range s {
		if n == v {
			return i
		}
	}

	return -1
}
