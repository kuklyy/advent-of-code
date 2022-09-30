package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/kuklyy/advent-of-code/common"
)

func main() {
	part1()
	part2()
}

// 2967
func part1() {
	template := make([]string, 0, 4096)
	insertionRules := make([][]string, 0, 128)

	scanner, input := common.ReadFile("input.txt")

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
		stepTemplate := make([][]string, 0, len(template))
		for j := 0; j < len(template)-1; j++ {
			stepTemplate = append(stepTemplate, []string{fmt.Sprintf("%s%s", template[j], template[j+1]), ""})
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
	}

	countMap := make(map[string]int)

	for _, v := range template {
		if _, ok := countMap[v]; !ok {
			countMap[v] = 0
		}
		countMap[v] += 1
	}

	maxValue := math.MinInt64
	minValue := math.MaxInt64

	for _, v := range countMap {
		if v > maxValue {
			maxValue = v
		}

		if v < minValue {
			minValue = v
		}
	}

	fmt.Println(maxValue - minValue)
}

func part2() {
	pairMap := make(map[string]int)
	insertionRules := make([][]string, 0, 128)

	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	readRules := false

	countMap := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			readRules = true
			continue
		}

		if !readRules {
			data := strings.Split(line, "")
			for i := 0; i < len(data)-1; i++ {
				key := fmt.Sprintf("%s%s", data[i], data[i+1])
				appendPair(pairMap, 1, key)
				if i == len(data)-2 {
					appendPair(countMap, 1, data[i+1])
				}
			}
		} else {
			parsedRule := strings.Split(line, " -> ")
			insertionRules = append(insertionRules, []string{parsedRule[0], parsedRule[1]})
		}
	}

	step := 40

	for i := 0; i < step; i++ {
		stepPairMap := make(map[string]int)

		for _, rule := range insertionRules {
			for pair, count := range pairMap {
				if pair == rule[0] {
					appendPair(stepPairMap, count, fmt.Sprintf("%s%s", string(pair[0]), rule[1]))
					appendPair(stepPairMap, count, fmt.Sprintf("%s%s", rule[1], string(pair[1])))
				}
			}
		}

		pairMap = stepPairMap
	}

	for k, v := range pairMap {
		appendPair(countMap, v, string(k[0]))
	}

	maxValue := math.MinInt64
	minValue := math.MaxInt64

	for _, v := range countMap {
		if v > maxValue {
			maxValue = v
		}

		if v < minValue {
			minValue = v
		}
	}

	fmt.Println(maxValue - minValue)

}

func appendPair(pairMap map[string]int, count int, key string) {
	if _, ok := pairMap[key]; !ok {
		pairMap[key] = 0
	}
	pairMap[key] += count
}
