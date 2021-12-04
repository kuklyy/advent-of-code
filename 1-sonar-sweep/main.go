package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// https://adventofcode.com/2021/day/1
func main() {
	linearMesurement()
	slidingWindow()
}

// 1692
func linearMesurement() {
	input, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	prevInput := -1
	increseCount := 0

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if prevInput == -1 {
			prevInput = line
			continue
		}

		if line > prevInput {
			increseCount++
		}

		prevInput = line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(increseCount)
}

// 1724
func slidingWindow() {
	input, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	measurements := make([]int, 0, 3000)

	index := 0

	windowSize := 3

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		for i := index; i <= index+windowSize-1; i++ {
			if i == len(measurements) {
				measurements = append(measurements, line)
			} else {
				measurements[i] = measurements[i] + line
			}
		}

		index++

	}

	increseCount := 0

	prevValue := -1

	for _, measure := range measurements[windowSize-1 : len(measurements)-windowSize+1] {

		if prevValue == -1 {
			prevValue = measure
			continue
		}

		if measure > prevValue {
			increseCount++
		}
		prevValue = measure
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(increseCount)
}
