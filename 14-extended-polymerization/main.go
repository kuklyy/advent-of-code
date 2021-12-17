package main

import (
	"fmt"
	"strings"

	"github.com/kuklyy/advent-of-code-2021/common"
)

func main() {
	part1()
}

func part1() {
	template := make([]string, 0, 4096)
	insertionRules := make([][]string, 0, 128)

	scanner, input := common.ReadFile("sample.txt")

	defer input.Close()

	readRules := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readRules = true
			continue
		}

		if !readRules {
			template = append(template, strings.Split(line, "")...)
		} else {
			parsedRule := strings.Split(line, " -> ")
			insertionRules = append(insertionRules, []string{parsedRule[0], parsedRule[1]})
		}
	}

	step := 10

	for i := 0; i < step; i++ {
		stepTemplate := make(map[int][]string)
		for j := 0; j < len(template)-1; j++ {
			stepTemplate[j] = []string{fmt.Sprintf("%s%s", template[j], template[j+1]), ""}
		}

		for _, rule := range insertionRules {
			for k := range stepTemplate {
				if stepTemplate[k][0] == rule[0] {
					stepTemplate[k][1] = rule[1]
				}
			}
		}

		newTemplate := make([]string, 0, 4096)
		newTemplate = append(newTemplate, string(stepTemplate[0][0][0]))

		for _, k := range stepTemplate {
			if k[1] != "" {
				newTemplate = append(newTemplate, k[1])
			}
			newTemplate = append(newTemplate, string(k[0][1]))
		}

		template = newTemplate
		fmt.Println(template)
	}

	fmt.Println(len(template))

	countMap := make(map[string]int)

	for _, v := range template {
		if _, ok := countMap[v]; !ok {
			countMap[v] = 0
		}
		countMap[v] += 1
	}

	fmt.Println(countMap)
}
