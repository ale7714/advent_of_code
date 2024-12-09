package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"unicode"
)

func main() {
	grid := readGrid("input.txt")

	height := len(grid)
	if height == 0 {
		fmt.Println("Map is empty. Exiting.")
		return
	}
	width := len(grid[0])
	antennas := getAntennas(grid, height, width)
	antinodes := getAntinodes(antennas, height, width)

	fmt.Println("Number of antinodes:", len(antinodes))

	antinodesRH := antinodeRH(antennas, height, width)
	fmt.Println("Number of antinodes after the effects of resonant harmonics:", len(antinodesRH))

}

func readGrid(filename string) [][]rune {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return grid
}

func getAntennas(grid [][]rune, height int, width int) map[rune][]point {
	antennas := make(map[rune][]point)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := grid[y][x]
			if c != '.' && (unicode.IsLetter(c) || unicode.IsDigit(c)) {
				antennas[c] = append(antennas[c], point{x, y})
			}
		}
	}

	return antennas
}

func getAntinodes(antennas map[rune][]point, height int, width int) map[[2]int]bool {
	antinodes := make(map[[2]int]bool)

	for _, coords := range antennas {
		if len(coords) < 2 {
			continue
		}

		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				A := coords[i]
				B := coords[j]

				p1x := 2*A.x - B.x
				p1y := 2*A.y - B.y

				p2x := 2*B.x - A.x
				p2y := 2*B.y - A.y

				if p1x >= 0 && p1x < width && p1y >= 0 && p1y < height {
					antinodes[[2]int{p1x, p1y}] = true
				}
				if p2x >= 0 && p2x < width && p2y >= 0 && p2y < height {
					antinodes[[2]int{p2x, p2y}] = true
				}
			}
		}
	}

	return antinodes
}

type point struct {
	x, y int
}

func gcd(a, b int) int {
	a = int(math.Abs(float64(a)))
	b = int(math.Abs(float64(b)))
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func antinodeRH(antennas map[rune][]point, height int, width int) map[[2]int]bool {
	antinodes := make(map[[2]int]bool)
	for _, coords := range antennas {
		if len(coords) < 2 {
			continue
		}

		for i := 0; i < len(coords); i++ {
			for j := i + 1; j < len(coords); j++ {
				A := coords[i]
				B := coords[j]
				dx := B.x - A.x
				dy := B.y - A.y
				g := gcd(dx, dy)
				dx /= g
				dy /= g

				xCur, yCur := A.x, A.y
				for {
					if xCur < 0 || xCur >= width || yCur < 0 || yCur >= height {
						break
					}
					antinodes[[2]int{xCur, yCur}] = true
					xCur += dx
					yCur += dy
				}

				xCur, yCur = A.x-dx, A.y-dy
				for {
					if xCur < 0 || xCur >= width || yCur < 0 || yCur >= height {
						break
					}
					antinodes[[2]int{xCur, yCur}] = true
					xCur -= dx
					yCur -= dy
				}
			}
		}
	}

	return antinodes
}
