package common

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) *bufio.Scanner {
	input, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	return scanner
}
