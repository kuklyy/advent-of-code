package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const boardWidth = 5

func main() {
	part1()
	part2()
}

// 4662
func part1() {

	boards := make([][][]int, 0, 3000)

	numbersToMark := make([]int, 0, 3000)

	markedBoards := make([][][]int, 0, 3000)

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	boardIndex := -1
	rowIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(numbersToMark) == 0 {
			for _, numberString := range strings.Split(line, ",") {
				number, err := strconv.Atoi(numberString)
				if err != nil {
					log.Fatal(err)
				}

				numbersToMark = append(numbersToMark, number)
			}
			continue
		}

		if line == "" {
			boardIndex++
			rowIndex = 0
			boards = append(boards, make([][]int, boardWidth))
			continue
		}

		for _, boardMember := range strings.Fields(line) {
			number, err := strconv.Atoi(boardMember)
			if err != nil {
				log.Fatal(err)
			}

			boards[boardIndex][rowIndex] = append(boards[boardIndex][rowIndex], number)
		}

		rowIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var winningIndex int
	var lastNumberIndex int

	winningIndex, lastNumberIndex = calculateWinner(boards, markedBoards, numbersToMark)

	if winningIndex < 0 {
		log.Fatal()
	}

	score := calculateScore(boards[winningIndex], numbersToMark[:lastNumberIndex+1], numbersToMark[lastNumberIndex])

	fmt.Println(score)
}

func calculateWinner(boards [][][]int, markedBoards [][][]int, numbers []int) (int, int) {
	for drawIndex, drawNumber := range numbers {
		for boardIndex, board := range boards {
			for i, row := range board {
				for j, column := range row {
					if len(markedBoards) == boardIndex {
						markedBoards = append(markedBoards, make([][]int, 2))
						markedBoards[boardIndex][0] = make([]int, boardWidth)
						markedBoards[boardIndex][1] = make([]int, boardWidth)
					}

					if column == drawNumber {
						markedBoards[boardIndex][0][i] = markedBoards[boardIndex][0][i] + 1
						markedBoards[boardIndex][1][j] = markedBoards[boardIndex][1][j] + 1

						if markedBoards[boardIndex][0][i] == boardWidth || markedBoards[boardIndex][1][j] == boardWidth {
							winningIndex := boardIndex
							return winningIndex, drawIndex
						}
					}
				}
			}
		}
	}

	return -1, -1
}

// 12080
func part2() {

	boards := make([][][]int, 0, 3000)

	numbersToMark := make([]int, 0, 3000)

	markedBoards := make([][][]int, 0, 3000)

	input, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	boardIndex := -1
	rowIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(numbersToMark) == 0 {
			for _, numberString := range strings.Split(line, ",") {
				number, err := strconv.Atoi(numberString)
				if err != nil {
					log.Fatal(err)
				}

				numbersToMark = append(numbersToMark, number)
			}
			continue
		}

		if line == "" {
			boardIndex++
			rowIndex = 0
			boards = append(boards, make([][]int, boardWidth))
			continue
		}

		for _, boardMember := range strings.Fields(line) {
			number, err := strconv.Atoi(boardMember)
			if err != nil {
				log.Fatal(err)
			}

			boards[boardIndex][rowIndex] = append(boards[boardIndex][rowIndex], number)
		}

		rowIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var winningBoard int
	var lastNumberIndex int

	winningBoard, lastNumberIndex = calculateWinnerPart2(boards, markedBoards, numbersToMark)

	if winningBoard < 0 {
		log.Fatal()
	}

	score := calculateScore(boards[winningBoard], numbersToMark[:lastNumberIndex+1], numbersToMark[lastNumberIndex])

	fmt.Println(score)
}

func calculateWinnerPart2(boards [][][]int, markedBoards [][][]int, numbers []int) (int, int) {
	winningBoards := make([]int, 0, len(boards))
	lastDrawIndex := -1

	for drawIndex, drawNumber := range numbers {
		for boardIndex, board := range boards {
			skipBoard := false
			for _, winningBoard := range winningBoards {
				if boardIndex == winningBoard {
					skipBoard = true
				}
			}
			if skipBoard {
				continue
			}

			for i, row := range board {
				foundWinner := false

				for j, column := range row {
					if len(markedBoards) == boardIndex {
						markedBoards = append(markedBoards, make([][]int, 2))
						markedBoards[boardIndex][0] = make([]int, boardWidth)
						markedBoards[boardIndex][1] = make([]int, boardWidth)
					}

					if column == drawNumber {
						markedBoards[boardIndex][0][i] = markedBoards[boardIndex][0][i] + 1
						markedBoards[boardIndex][1][j] = markedBoards[boardIndex][1][j] + 1

						if markedBoards[boardIndex][0][i] == boardWidth || markedBoards[boardIndex][1][j] == boardWidth {
							foundWinner = true
							winningBoards = append(winningBoards, boardIndex)
							lastDrawIndex = drawIndex
						}
					}

					if foundWinner {
						break
					}
				}

				if foundWinner {
					break
				}
			}
		}
	}

	return winningBoards[len(winningBoards)-1], lastDrawIndex
}

func calculateScore(board [][]int, drawNumbers []int, lastDrawNumber int) int {
	sum := 0
	for _, row := range board {
		for _, value := range row {
			isNumberMarked := false
			for _, drawNumber := range drawNumbers {
				if value == drawNumber {
					isNumberMarked = true
					break
				}
			}
			if !isNumberMarked {
				sum += value
			}
		}
	}

	return sum * lastDrawNumber
}
