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

// 1714680
func part1() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	horizontalPosition := 0
	depth := 0

	for scanner.Scan() {
		command := strings.Fields(scanner.Text())
		action := command[0]
		strength, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "forward":
			horizontalPosition += strength
		case "down":
			depth += strength
		case "up":
			depth -= strength
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(horizontalPosition * depth)
}

// 1963088820
func part2() {
	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	horizontalPosition := 0
	depth := 0
	aim := 0

	for scanner.Scan() {
		command := strings.Fields(scanner.Text())
		action := command[0]
		strength, err := strconv.Atoi(command[1])
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "forward":
			horizontalPosition += strength
			depth += aim * strength
		case "down":
			aim += strength
		case "up":
			aim -= strength
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(horizontalPosition * depth)
}
