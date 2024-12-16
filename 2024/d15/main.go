package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Directions map for convenience
var directions = map[rune][2]int{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

func main() {
	// Open input file
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Read the warehouse map lines until we reach a line that suggests we've moved on to moves
	var warehouse []string
	var moveLines []string

	// First, we read lines until we encounter the first move line
	// We know the map is typically surrounded by '#' lines. Once we start reading lines
	// that don't fit the map pattern (or after the map ends) the rest should be moves.
	// The puzzle states: The rest of the document describes the moves.
	//
	// We'll assume:
	// 1) The map is fully enclosed in '#' (like a border)
	// 2) After the map, all subsequent lines are part of the moves.
	//
	// We can detect map end by:
	// - The map likely ends once we encounter a line that doesn't have '#' border or is empty.
	// However, from the problem, it might be safer to read all lines first, then separate
	// the map from moves. The moves are only these characters: ^, v, <, >
	// The map lines contain '#' at the edges and likely '.' 'O' '@'.
	//
	// We'll do this: read all lines first, then separate them.
	var allLines []string
	for scanner.Scan() {
		line := scanner.Text()
		allLines = append(allLines, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Identify where map ends and moves begin.
	// The map lines will contain '#' at least in the first and last line.
	// Moves lines contain only the characters ^, v, <, > (and possibly empty).
	// We'll find the first line that is strictly made of move characters and consider that the start of moves.

	isMoveLine := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		for _, ch := range s {
			if ch != '^' && ch != 'v' && ch != '<' && ch != '>' {
				return false
			}
		}
		return true
	}

	// We'll find the last line of the map by scanning from the top until we hit a line that looks like moves
	mapEndIdx := -1
	for i, line := range allLines {
		// If line is empty, skip. Move lines might come after empty lines.
		// If line looks like moves, we stop.
		if isMoveLine(strings.ReplaceAll(line, " ", "")) {
			mapEndIdx = i - 1
			break
		}
	}
	if mapEndIdx == -1 {
		// If we never found a moves line, then possibly all lines are map lines or moves are empty
		// Check from bottom if there's a moves line.
		for i := len(allLines) - 1; i >= 0; i-- {
			if isMoveLine(strings.ReplaceAll(allLines[i], " ", "")) {
				mapEndIdx = i - 1
				break
			}
		}
		if mapEndIdx == -1 {
			// No moves at all, map is all lines
			mapEndIdx = len(allLines) - 1
		}
	}

	if mapEndIdx < 0 {
		mapEndIdx = len(allLines) - 1
	}

	if mapEndIdx >= len(allLines) {
		mapEndIdx = len(allLines) - 1
	}

	warehouse = allLines[:mapEndIdx+1]
	moveLines = allLines[mapEndIdx+1:]

	// Combine all moves into a single string
	var movesBuilder strings.Builder
	for _, ml := range moveLines {
		movesBuilder.WriteString(strings.TrimSpace(ml))
	}
	moves := movesBuilder.String()

	// Convert warehouse into a modifiable structure (slice of rune slices)
	grid := make([][]rune, len(warehouse))
	for i, line := range warehouse {
		grid[i] = []rune(line)
	}

	// Find the robot position
	var robotR, robotC int
	foundRobot := false
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '@' {
				robotR, robotC = r, c
				foundRobot = true
				break
			}
		}
		if foundRobot {
			break
		}
	}

	// Function to check if a position is inside the grid
	inBounds := func(r, c int) bool {
		return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[r])
	}

	// For each move, attempt to move the robot
	for _, move := range moves {
		d := directions[move]
		dr, dc := d[0], d[1]

		nextR := robotR + dr
		nextC := robotC + dc

		// Check what is at the next cell
		if !inBounds(nextR, nextC) {
			// Out of bounds or can't move, skip
			continue
		}
		nextCell := grid[nextR][nextC]

		if nextCell == '.' {
			// Just move the robot
			grid[robotR][robotC] = '.'
			grid[nextR][nextC] = '@'
			robotR, robotC = nextR, nextC
		} else if nextCell == 'O' {
			// Need to push boxes
			// Find the chain of boxes in that direction
			boxPositions := make([][2]int, 0)
			curR, curC := nextR, nextC
			for {
				if !inBounds(curR, curC) {
					// hit boundary without empty space, can't push
					boxPositions = nil
					break
				}
				if grid[curR][curC] == 'O' {
					boxPositions = append(boxPositions, [2]int{curR, curC})
					curR += dr
					curC += dc
				} else {
					// We hit a non-box cell
					break
				}
			}

			if boxPositions == nil {
				// Can't move
				continue
			}

			// Now curR, curC is the cell after the chain of boxes
			if !inBounds(curR, curC) {
				// Can't push outside
				continue
			}
			if grid[curR][curC] == '.' {
				// We can push all boxes forward
				// Move the boxes from the far end
				grid[robotR][robotC] = '.'       // robot moves out
				grid[robotR+dr][robotC+dc] = '@' // robot moves into position of first box
				robotR += dr
				robotC += dc

				// Now push boxes
				// The last box in the chain moves first
				for i := len(boxPositions) - 1; i >= 0; i-- {
					br, bc := boxPositions[i][0], boxPositions[i][1]
					grid[br+dr][bc+dc] = 'O'
					grid[br][bc] = '.'
				}
			} else {
				// The cell after chain is either '#' or 'O'
				// If '#' we can't move
				// If 'O', that means chain ended in a box with no empty space after.
				// So we can't push.
				if grid[curR][curC] == 'O' || grid[curR][curC] == '#' {
					continue
				}
			}
		} else if nextCell == '#' {
			// Can't move into a wall
			continue
		} else if nextCell == '@' {
			// Should never happen, only one robot
			continue
		}
	}

	// After all moves, calculate sum of GPS coordinates of all boxes
	// GPS = 100*row + col
	sum := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 'O' {
				sum += 100*r + c
			}
		}
	}

	fmt.Println(sum)
}
