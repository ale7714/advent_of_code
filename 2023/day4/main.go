package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	num     int
	winning []int
	own     []int
	matches []int
}

func toIntArray(a string) []int {
	var r []int

	re := regexp.MustCompile(`\s+`)

	for _, e := range re.Split(strings.TrimSpace(a), -1) {

		num, err := strconv.Atoi(e)

		if err != nil {

			fmt.Println("toIntArray", e, err, re.Split(a, -1))
		}

		r = append(r, num)
	}
	return r
}

func getMatches(g Game) []int {
	var matches []int

	for _, w := range g.winning {
		for _, o := range g.own {
			if w == o {
				matches = append(matches, w)
			}
		}
	}

	return matches
}

func getPoints(g Game) int {
	if len(g.matches) == 0 {
		return 0
	}

	p := 1

	if len(g.matches) > 1 {
		for i := 0; i < len(g.matches)-1; i++ {
			p *= 2
		}
	}

	return p
}

func getWinningCards(games []Game) int {
	r := 0
	copies := make(map[int]int)

	for _, g := range games {
		if copies[g.num] == 0 {
			copies[g.num] = 1
		} else {
			copies[g.num] += 1

		}

		if len(g.matches) == 0 {
			continue
		}

		cards := copies[g.num]
		for i := g.num + 1; i <= g.num+len(g.matches); i++ {
			copies[i] += cards
		}
	}

	for _, c := range copies {

		r += c
	}

	return r
}

func main() {

	file, err := os.Open("inp.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPoints := 0
	var cards []Game

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		num, err := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(line[0]))
		lists := strings.Split(line[1], " | ")

		if err != nil {
			fmt.Println(err)
		}

		g := Game{
			num:     num,
			winning: toIntArray(lists[0]),
			own:     toIntArray(lists[1]),
		}

		g.matches = getMatches(g)
		cards = append(cards, g)

		totalPoints += getPoints(g)
	}

	fmt.Println("Part 1:", totalPoints)
	fmt.Println("Part 2:", getWinningCards(cards))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
