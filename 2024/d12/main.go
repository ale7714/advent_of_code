package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	totalPerimeterPrice, totalSidesPrice := calculatePrice(grid)
	fmt.Println("Total Perimeter Price:", totalPerimeterPrice)
	fmt.Println("Total Sides Price:", totalSidesPrice)
}

func parseInput(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func calculateCorners(grid [][]rune, x, y int, letter rune) int {
	rows, cols := len(grid), len(grid[0])
	corners := 0

	directions := [][3][2]int{
		{{-1, 0}, {0, -1}, {-1, -1}},
		{{-1, 0}, {0, 1}, {-1, 1}},
		{{1, 0}, {0, -1}, {1, -1}},
		{{1, 0}, {0, 1}, {1, 1}},
	}

	for _, d := range directions {
		up, left, diag := d[0], d[1], d[2]
		upX, upY := x+up[0], y+up[1]
		leftX, leftY := x+left[0], y+left[1]
		diagX, diagY := x+diag[0], y+diag[1]

		upInBounds := upX >= 0 && upY >= 0 && upX < rows && upY < cols
		leftInBounds := leftX >= 0 && leftY >= 0 && leftX < rows && leftY < cols
		diagInBounds := diagX >= 0 && diagY >= 0 && diagX < rows && diagY < cols

		isConvex := (!upInBounds || grid[upX][upY] != letter) &&
			(!leftInBounds || grid[leftX][leftY] != letter)

		isConcave := (upInBounds && grid[upX][upY] == letter) &&
			(leftInBounds && grid[leftX][leftY] == letter) &&
			(diagInBounds && grid[diagX][diagY] != letter)

		if isConvex || isConcave {
			corners++
		}
	}

	return corners
}

func calculateRegionMetrics(grid [][]rune, x, y int, letter rune, visited map[[2]int]bool) (int, int, int) {
	rows, cols := len(grid), len(grid[0])
	stack := [][2]int{{x, y}}
	visited[[2]int{x, y}] = true

	area := 0
	perimeter := 0
	sides := 0

	directions := [][2]int{
		{1, 0},  // Down
		{-1, 0}, // Up
		{0, 1},  // Right
		{0, -1}, // Left
	}

	for len(stack) > 0 {
		cx, cy := stack[len(stack)-1][0], stack[len(stack)-1][1]
		stack = stack[:len(stack)-1]
		area++

		// Count corners for sides calculation
		sides += calculateCorners(grid, cx, cy, letter)

		// Calculate perimeter and neighbors
		for _, d := range directions {
			nx, ny := cx+d[0], cy+d[1]

			if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
				if grid[nx][ny] != letter {
					perimeter++
				} else if !visited[[2]int{nx, ny}] {
					visited[[2]int{nx, ny}] = true
					stack = append(stack, [2]int{nx, ny})
				}
			} else {
				perimeter++
			}
		}
	}

	return area, perimeter, sides
}

func calculatePrice(grid [][]rune) (int, int) {
	rows, cols := len(grid), len(grid[0])
	visited := make(map[[2]int]bool)
	totalPerimeterPrice := 0
	totalSidesPrice := 0

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if !visited[[2]int{x, y}] {
				area, perimeter, sides := calculateRegionMetrics(grid, x, y, grid[x][y], visited)
				totalPerimeterPrice += area * perimeter
				totalSidesPrice += area * sides
				// fmt.Printf("Region %c: Area = %d, Perimeter = %d, Sides = %d, Perimeter Price = %d, Sides Price = %d\n", grid[x][y], area, perimeter, sides, area*perimeter, area*sides)
			}
		}
	}

	return totalPerimeterPrice, totalSidesPrice
}
