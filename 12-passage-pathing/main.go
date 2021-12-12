package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/kuklyy/advent-of-code-2021/common"
)

func main() {
	part1()
	part2()
}

func part1() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	list := make(map[string]map[string]bool)

	for scanner.Scan() {
		parsedLine := strings.Split(scanner.Text(), "-")
		start := parsedLine[0]
		end := parsedLine[1]
		_, ok := list[start]
		if !ok {
			list[start] = make(map[string]bool)
		}
		list[start][end] = true

		_, ok = list[end]
		if !ok {
			list[end] = make(map[string]bool)
		}
		list[end][start] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	paths := make(map[int]map[string]int)
	paths[0] = make(map[string]int)
	paths[0]["start"] = 1
	pathIndex := 0

	calculatePath(paths, list["start"], paths[0], list, &pathIndex)

	for i, v := range paths {
		if _, ok := v["end"]; !ok {
			delete(paths, i)
		}
	}

	fmt.Println(len(paths))
}

func calculatePath(paths map[int]map[string]int, element map[string]bool, path map[string]int, list map[string]map[string]bool, pathIndex *int) {

	if _, ok := path["end"]; ok {
		return
	}

	for k := range element {
		if strings.ToLower(k) == k {
			if _, ok := path[k]; ok {
				continue
			}
		}

		*pathIndex += 1
		paths[*pathIndex] = copyMap(path)
		paths[*pathIndex][k] += 1
		calculatePath(paths, list[k], paths[*pathIndex], list, pathIndex)
	}
}

func part2() {
	scanner, input := common.ReadFile("input.txt")

	defer input.Close()

	list := make(map[string]map[string]bool)

	for scanner.Scan() {
		parsedLine := strings.Split(scanner.Text(), "-")
		start := parsedLine[0]
		end := parsedLine[1]
		_, ok := list[start]
		if !ok {
			list[start] = make(map[string]bool)
		}
		list[start][end] = true

		_, ok = list[end]
		if !ok {
			list[end] = make(map[string]bool)
		}
		list[end][start] = true
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	paths := make(map[int]map[string]int)
	paths[0] = make(map[string]int)
	paths[0]["start"] = 1
	pathIndex := 0

	calculatePath2(paths, list["start"], paths[0], list, &pathIndex)

	for i, v := range paths {
		if _, ok := v["end"]; !ok {
			delete(paths, i)
		}
	}
	fmt.Println(len(paths))
}

func calculatePath2(paths map[int]map[string]int, element map[string]bool, path map[string]int, list map[string]map[string]bool, pathIndex *int) {
	if _, ok := path["end"]; ok {
		return
	}

	for k := range element {
		if k == "start" {
			continue
		}
		if strings.ToLower(k) == k {
			if _, ok := path[k]; ok {
				canVisitSmall := true
				for p, v := range path {
					if strings.ToLower(p) == p && v == 2 {
						canVisitSmall = false
						break
					}
				}
				if !canVisitSmall {
					continue
				}
			}
		}

		*pathIndex += 1
		paths[*pathIndex] = copyMap(path)
		paths[*pathIndex][k] += 1
		calculatePath2(paths, list[k], paths[*pathIndex], list, pathIndex)
	}
}

func copyMap(in map[string]int) map[string]int {
	out := make(map[string]int)

	for k, v := range in {
		out[k] = v
	}

	return out
}
