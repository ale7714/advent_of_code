package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	trailheads := findTrailheads(grid)

	score := 0
	for _, trailhead := range trailheads {
		score += scoreTrail(grid, trailhead)
	}

	fmt.Printf("Trailheads Score: %d\n", score)

	totalRating := calculateTotalRating(grid, trailheads)
	fmt.Println("Total Rating:", totalRating)
}

func parseInput(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		row := make([]int, len(line))
		for i, char := range line {
			row[i] = int(char - '0')
		}
		grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func findTrailheads(grid [][]int) [][2]int {
	var trailheads [][2]int
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				trailheads = append(trailheads, [2]int{row, col})
			}
		}
	}
	return trailheads
}

func scoreTrail(grid [][]int, start [2]int) int {
	rows, cols := len(grid), len(grid[0])
	queue := [][2]int{start}
	visited := make(map[[2]int]bool)
	reachableNines := make(map[[2]int]bool)

	visited[start] = true

	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		x, y := current[0], current[1]

		if grid[x][y] == 9 {
			reachableNines[current] = true
		}

		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]

			if nx >= 0 && ny >= 0 && nx < rows && ny < cols {
				if !visited[[2]int{nx, ny}] && grid[nx][ny] == grid[x][y]+1 {
					queue = append(queue, [2]int{nx, ny})
					visited[[2]int{nx, ny}] = true
				}
			}
		}
	}

	return len(reachableNines)
}

func countTrails(grid [][]int, x, y int, memo map[[2]int]int) int {
	if val, exists := memo[[2]int{x, y}]; exists {
		return val
	}

	trails := 0
	directions := [][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}

	for _, d := range directions {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid[0]) {
			if grid[nx][ny] == grid[x][y]+1 {
				trails += countTrails(grid, nx, ny, memo)
			}
		}
	}

	if grid[x][y] == 9 {
		trails += 1
	}

	memo[[2]int{x, y}] = trails
	return trails
}

func calculateTotalRating(grid [][]int, trailheads [][2]int) int {
	memo := make(map[[2]int]int)
	totalRating := 0

	for _, trailhead := range trailheads {
		totalRating += countTrails(grid, trailhead[0], trailhead[1], memo)
	}

	return totalRating
}
