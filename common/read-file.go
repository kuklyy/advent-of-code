package common

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string) (*bufio.Scanner, *os.File) {
	input, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)

	return scanner, input
}
