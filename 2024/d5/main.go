package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	constraints := readRules(scanner)

	correctSum, incorrectSum := processUpdates(scanner, constraints)

	fmt.Printf("Sum middle page numbers of correctly ordered updates: %d\n", correctSum)
	fmt.Printf("Sum middle page numbers of incorrectly ordered updates: %d\n", incorrectSum)
}

func readRules(scanner *bufio.Scanner) map[int][]int {
	constraints := make(map[int][]int)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}

		x, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		y, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err1 == nil && err2 == nil {
			constraints[x] = append(constraints[x], y)
		}
	}
	return constraints
}

func processUpdates(scanner *bufio.Scanner, constraints map[int][]int) (int, int) {
	correctSum := 0
	incorrectSum := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		pages := parsePages(line)
		indexMap := buildIndexMap(pages)

		if isCorrectlyOrdered(pages, indexMap, constraints) {
			correctSum += pages[len(pages)/2]
		} else {
			ordered := reorderUpdate(pages, constraints)
			incorrectSum += ordered[len(ordered)/2]
		}
	}

	return correctSum, incorrectSum
}

func parsePages(line string) []int {
	parts := strings.Split(line, ",")
	pages := make([]int, 0, len(parts))
	for _, p := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(p))
		if err == nil {
			pages = append(pages, num)
		}
	}
	return pages
}

func buildIndexMap(pages []int) map[int]int {
	indexMap := make(map[int]int, len(pages))
	for i, p := range pages {
		indexMap[p] = i
	}
	return indexMap
}

func isCorrectlyOrdered(pages []int, indexMap map[int]int, constraints map[int][]int) bool {
	for _, x := range pages {
		afterPages, exists := constraints[x]
		if !exists {
			continue
		}

		xPos := indexMap[x]
		for _, y := range afterPages {
			if yPos, yExists := indexMap[y]; yExists {
				if xPos > yPos {
					return false
				}
			}
		}
	}
	return true
}

func reorderUpdate(pages []int, constraints map[int][]int) []int {
	pageSet := make(map[int]bool, len(pages))
	for _, p := range pages {
		pageSet[p] = true
	}

	adj := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, p := range pages {
		inDegree[p] = 0
	}

	for _, x := range pages {
		if afterPages, exists := constraints[x]; exists {
			for _, y := range afterPages {
				if pageSet[y] {
					adj[x] = append(adj[x], y)
					inDegree[y]++
				}
			}
		}
	}

	queue := make([]int, 0)
	for p, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, p)
		}
	}

	var result []int
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		result = append(result, curr)

		for _, nxt := range adj[curr] {
			inDegree[nxt]--
			if inDegree[nxt] == 0 {
				queue = append(queue, nxt)
			}
		}
	}

	return result
}
