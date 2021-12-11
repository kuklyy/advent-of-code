package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/kuklyy/advent-of-code-2021/common"
)

func main() {
	part1()
	part2()
}

// 358737
func part1() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		syntaxLine := make(map[int]string)
		for _, c := range strings.Split(line, "") {
			syntaxError := false
			syntaxLen := len(syntaxLine)

			switch c {
			case "(", "[", "{", "<":
				syntaxLine[syntaxLen] = c

			case ")":
				if syntaxLine[syntaxLen-1] != "(" {
					score += 3
					syntaxError = true
				} else {
					delete(syntaxLine, syntaxLen-1)
				}
			case "]":
				if syntaxLine[syntaxLen-1] != "[" {
					score += 57
					syntaxError = true
				} else {
					delete(syntaxLine, syntaxLen-1)
				}
			case "}":
				if syntaxLine[syntaxLen-1] != "{" {
					score += 1197
					syntaxError = true
				} else {
					delete(syntaxLine, syntaxLen-1)
				}
			case ">":
				if syntaxLine[syntaxLen-1] != "<" {
					score += 25137
					syntaxError = true
				} else {
					delete(syntaxLine, syntaxLen-1)
				}
			}

			if syntaxError {
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(score)
}

// 4329504793
func part2() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	scores := make([]int, 0, 100)

	for scanner.Scan() {
		line := scanner.Text()
		syntaxStack := make(map[int]string)
		syntaxError := false
		for _, c := range strings.Split(line, "") {
			syntaxLen := len(syntaxStack)

			switch c {
			case "(", "[", "{", "<":
				syntaxStack[syntaxLen] = c

			case ")":
				if syntaxStack[syntaxLen-1] != "(" {
					syntaxError = true
				} else {
					delete(syntaxStack, syntaxLen-1)
				}
			case "]":
				if syntaxStack[syntaxLen-1] != "[" {
					syntaxError = true
				} else {
					delete(syntaxStack, syntaxLen-1)
				}
			case "}":
				if syntaxStack[syntaxLen-1] != "{" {
					syntaxError = true
				} else {
					delete(syntaxStack, syntaxLen-1)
				}
			case ">":
				if syntaxStack[syntaxLen-1] != "<" {
					syntaxError = true
				} else {
					delete(syntaxStack, syntaxLen-1)
				}
			}

			if syntaxError {
				break
			}

		}

		if !syntaxError {
			rowScore := 0
			for i := len(syntaxStack) - 1; i >= 0; i-- {
				rowScore *= 5
				switch syntaxStack[i] {
				case "(":
					rowScore += 1
				case "[":
					rowScore += 2
				case "{":
					rowScore += 3
				case "<":
					rowScore += 4
				}
			}
			scores = append(scores, rowScore)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])

}
