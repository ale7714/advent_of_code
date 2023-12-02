package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id    int
	cubes map[string]int
}

var bag map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func (g Game) Valid(b map[string]int) bool {
	for color, count := range g.cubes {
		if count > b[color] {
			return false
		}
	}
	return true
}

func (g Game) MinSet() int {
	return g.cubes["red"] * g.cubes["green"] * g.cubes["blue"]
}

func parseGame(line string) Game {
	parsed := strings.Split(line, ":")
	id, _ := strconv.Atoi(strings.TrimPrefix(parsed[0], "Game "))

	game := Game{
		id: id,
		cubes: map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		},
	}

	for _, cubes := range strings.Split(parsed[1], "; ") {
		for _, set := range strings.Split(strings.TrimSpace(cubes), ", ") {
			parsed = strings.Split(strings.TrimSpace(set), " ")
			num, _ := strconv.Atoi(parsed[0])
			if num > game.cubes[parsed[1]] {
				game.cubes[parsed[1]] = num
			}
		}

	}

	return game
}

func main() {

	file, err := os.Open("inp.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumValidGames := 0
	sumMinimumSet := 0

	for scanner.Scan() {
		game := parseGame(scanner.Text())
		if game.Valid(bag) {
			sumValidGames += game.id
		}
		sumMinimumSet += game.MinSet()
	}

	fmt.Println("Part 1: ", sumValidGames)
	fmt.Println("Part 1: ", sumMinimumSet)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
