package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var stones []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		nums := strings.Fields(line)
		for _, num := range nums {
			val, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			stones = append(stones, val)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stones, nil
}

func hasEvenDigits(num int) bool {
	return len(strconv.Itoa(num))%2 == 0
}

func countStonesAfterBlinks(stones []int, numBlinks int) int {
	counts := make(map[int]int)

	for _, stone := range stones {
		counts[stone]++
	}

	for i := 0; i < numBlinks; i++ {
		newCounts := make(map[int]int)

		for stone, count := range counts {
			if stone == 0 {
				newCounts[1] += count
			} else if hasEvenDigits(stone) {
				left, right := splitStone(stone)
				newCounts[left] += count
				newCounts[right] += count
			} else {
				newCounts[stone*2024] += count
			}
		}

		counts = newCounts
	}

	totalStones := 0
	for _, count := range counts {
		totalStones += count
	}

	return totalStones
}

func splitStone(num int) (int, int) {
	numStr := strconv.Itoa(num)
	mid := len(numStr) / 2
	left, _ := strconv.Atoi(numStr[:mid])
	right, _ := strconv.Atoi(numStr[mid:])
	return left, right
}

func main() {
	stones, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	numBlinks := 25
	totalStones := countStonesAfterBlinks(stones, numBlinks)
	fmt.Println("Total number of stones after", numBlinks, "blinks:", totalStones)

	numBlinks = 75
	totalStones = countStonesAfterBlinks(stones, numBlinks)
	fmt.Println("Total number of stones after", numBlinks, "blinks:", totalStones)
}
