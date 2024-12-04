package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	grid, err := readGridFromFile("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	count := countXMASInstances(grid)
	fmt.Printf("Found %d instances of XMAS\n", count)

	xShapedCount := countXShapedMAS(grid)
	fmt.Printf("Found %d X-shaped MAS patterns\n", xShapedCount)
}

func readGridFromFile(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var grid []string
	for _, line := range lines {
		if line != "" {
			grid = append(grid, line)
		}
	}
	return grid, nil
}

func checkXMAS(grid []string, row, col, dRow, dCol int) bool {
	rows := len(grid)
	cols := len(grid[0])
	if row+3*dRow >= rows || row+3*dRow < 0 || col+3*dCol >= cols || col+3*dCol < 0 {
		return false
	}
	return grid[row][col] == 'X' &&
		grid[row+dRow][col+dCol] == 'M' &&
		grid[row+2*dRow][col+2*dCol] == 'A' &&
		grid[row+3*dRow][col+3*dCol] == 'S'
}

func countXMASInstances(grid []string) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	directions := [][2]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // diagonal down-right
		{1, -1},  // diagonal down-left
		{0, -1},  // left
		{-1, 0},  // up
		{-1, 1},  // diagonal up-right
		{-1, -1}, // diagonal up-left
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			for _, dir := range directions {
				if checkXMAS(grid, row, col, dir[0], dir[1]) {
					count++
				}
			}
		}
	}
	return count
}

func checkMAS(grid []string, row, col, dRow, dCol int) bool {
	rows := len(grid)
	cols := len(grid[0])
	if row+2*dRow >= rows || row+2*dRow < 0 || col+2*dCol >= cols || col+2*dCol < 0 {
		return false
	}

	forwardMAS := grid[row][col] == 'M' &&
		grid[row+dRow][col+dCol] == 'A' &&
		grid[row+2*dRow][col+2*dCol] == 'S'

	backwardMAS := grid[row][col] == 'S' &&
		grid[row+dRow][col+dCol] == 'A' &&
		grid[row+2*dRow][col+2*dCol] == 'M'

	return forwardMAS || backwardMAS
}

func countXShapedMAS(grid []string) int {
	count := 0
	rows := len(grid)
	cols := len(grid[0])

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			// Check if center is 'A'
			if grid[row][col] != 'A' {
				continue
			}

			upperLeft := checkMAS(grid, row-1, col-1, 1, 1)
			upperRight := checkMAS(grid, row-1, col+1, 1, -1)

			if upperLeft && upperRight {
				count++
			}
		}
	}
	return count

}
