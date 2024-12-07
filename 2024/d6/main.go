package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		grid = append(grid, line)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	rows := len(grid)
	cols := len(grid[0])

	mutableGrid := make([][]rune, rows)
	for i := range grid {
		mutableGrid[i] = []rune(grid[i])
	}

	guardRow, guardCol, dir := findGuard(mutableGrid)

	visitedCountPart1 := simulateAndCount(mutableGrid, guardRow, guardCol, dir)
	fmt.Printf("Number of distinct positions visited: %d\n", visitedCountPart1)

	possibleLoopCount := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == guardRow && c == guardCol {
				continue
			}
			if mutableGrid[r][c] == '.' {
				mutableGrid[r][c] = '#'
				if causesLoop(mutableGrid, guardRow, guardCol, dir) {
					possibleLoopCount++
				}
				mutableGrid[r][c] = '.'
			}
		}
	}
	fmt.Printf("Number of possible positions for obstructions: %d\n", possibleLoopCount)
}

func findGuard(grid [][]rune) (int, int, Direction) {
	rows := len(grid)
	cols := len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			switch grid[r][c] {
			case '^':
				return r, c, Up
			case 'v':
				return r, c, Down
			case '<':
				return r, c, Left
			case '>':
				return r, c, Right
			}
		}
	}
	panic("Guard not found")
}

func simulateAndCount(grid [][]rune, startR, startC int, startDir Direction) int {
	rows := len(grid)
	cols := len(grid[0])

	r, c, d := startR, startC, startDir
	visited := make(map[[2]int]bool)
	visited[[2]int{r, c}] = true

	for {
		newR, newC := forwardPosition(r, c, d)
		if !inBounds(newR, newC, rows, cols) {
			break
		}

		if grid[newR][newC] == '#' {
			d = turnRight(d)
			continue
		}

		r, c = newR, newC
		visited[[2]int{r, c}] = true

		if !inBounds(r, c, rows, cols) {
			break
		}
	}
	return len(visited)
}

func causesLoop(grid [][]rune, startR, startC int, startDir Direction) bool {
	rows := len(grid)
	cols := len(grid[0])

	r, c, d := startR, startC, startDir
	visitedStates := make(map[[3]int]bool)
	visitedStates[[3]int{r, c, int(d)}] = true

	for {
		newR, newC := forwardPosition(r, c, d)

		if !inBounds(newR, newC, rows, cols) {
			return false
		}

		if grid[newR][newC] == '#' {
			d = turnRight(d)
			if visitedStates[[3]int{r, c, int(d)}] {
				return true
			}
			visitedStates[[3]int{r, c, int(d)}] = true
			continue
		}

		r, c = newR, newC
		if visitedStates[[3]int{r, c, int(d)}] {
			return true
		}
		visitedStates[[3]int{r, c, int(d)}] = true
	}
}

func forwardPosition(r, c int, d Direction) (int, int) {
	switch d {
	case Up:
		return r - 1, c
	case Down:
		return r + 1, c
	case Left:
		return r, c - 1
	case Right:
		return r, c + 1
	default:
		return r, c
	}
}

func turnRight(d Direction) Direction {
	return (d + 1) % 4
}

func inBounds(r, c, rows, cols int) bool {
	return r >= 0 && r < rows && c >= 0 && c < cols
}
