package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	x, y, vx, vy int
}

func main() {
	inputLines, err := readLines("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	robots := parseRobots(inputLines)

	width := 101
	height := 103

	finalPositions := calculateFinalPositions(robots, width, height, 100)
	quadrantCounts := countQuadrants(finalPositions, width, height)
	safetyFactor := calculateSafetyFactor(quadrantCounts)

	fmt.Printf("Safety Factor: %d\n", safetyFactor)

	earliestTime := findEarliestNonOverlappingTime(robots, width, height)
	fmt.Printf("Easter Egg Image Time: %d\n", earliestTime)
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseRobots(lines []string) []Robot {
	robots := make([]Robot, 0, len(lines))
	re := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`)
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match == nil {
			fmt.Printf("Invalid input line: %s\n", line)
			continue
		}
		px, _ := strconv.Atoi(match[1])
		py, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])
		robots = append(robots, Robot{x: px, y: py, vx: vx, vy: vy})
	}
	return robots
}

func calculateFinalPositions(robots []Robot, width, height, time int) []struct{ x, y int } {
	finalPositions := make([]struct{ x, y int }, len(robots))
	for i, robot := range robots {
		finalPositions[i].x = ((robot.x+robot.vx*time)%width + width) % width
		finalPositions[i].y = ((robot.y+robot.vy*time)%height + height) % height
	}
	return finalPositions
}

func countQuadrants(positions []struct{ x, y int }, width, height int) []int {
	midX := width / 2
	midY := height / 2
	quadrantCounts := make([]int, 4)

	for _, pos := range positions {
		if pos.x < midX && pos.y < midY {
			quadrantCounts[0]++
		} else if pos.x > midX && pos.y < midY {
			quadrantCounts[1]++
		} else if pos.x < midX && pos.y > midY {
			quadrantCounts[2]++
		} else if pos.x > midX && pos.y > midY {
			quadrantCounts[3]++
		}
	}
	return quadrantCounts
}

func calculateSafetyFactor(counts []int) int {
	safetyFactor := 1
	for _, count := range counts {
		safetyFactor *= count
	}
	return safetyFactor
}

func findEarliestNonOverlappingTime(robots []Robot, width, height int) int {
	for t := 0; t < width*height*2; t++ {
		positions := calculateFinalPositions(robots, width, height, t)
		if !hasOverlaps(positions) {
			return t
		}
	}
	return -1
}

func hasOverlaps(positions []struct{ x, y int }) bool {
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			if positions[i].x == positions[j].x && positions[i].y == positions[j].y {
				return true
			}
		}
	}
	return false
}
