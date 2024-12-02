package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	input, _ := os.ReadFile("inp.txt")

	grid, start := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
			if r == 'S' {
				start = image.Point{x, y}
			}
		}
	}

	for s, r := range map[string]rune{
		`[7F|].[JL|].`: '|', `.[-7J].[-FL]`: '-', `[7F|][-7J]..`: 'L',
		`[7F|]..[-FL]`: 'J', `..[JL|][-FL]`: '7', `.[-7J][JL|].`: 'F',
	} {
		if ok, _ := regexp.MatchString(s, string([]rune{
			grid[start.Add(image.Point{0, -1})], grid[start.Add(image.Point{1, 0})],
			grid[start.Add(image.Point{0, 1})], grid[start.Add(image.Point{-1, 0})],
		})); ok {
			grid[start] = r
		}
	}

	path, area := []image.Point{start}, 0
	for p, n := start, start; ; path = append(path, n) {
		p, n = n, start

		for _, d := range map[rune][]image.Point{
			'|': {{0, -1}, {0, 1}}, '-': {{1, 0}, {-1, 0}}, 'L': {{0, -1}, {1, 0}},
			'J': {{0, -1}, {-1, 0}}, '7': {{0, 1}, {-1, 0}}, 'F': {{0, 1}, {1, 0}},
		}[grid[p]] {
			if !slices.Contains(path, p.Add(d)) {
				n = p.Add(d)
			}
		}

		area += p.X*n.Y - p.Y*n.X
		if n == start {
			break
		}
	}
	fmt.Println(len(path) / 2)
	fmt.Println(int(math.Abs(float64(area)))/2 - len(path)/2 + 1)
}
