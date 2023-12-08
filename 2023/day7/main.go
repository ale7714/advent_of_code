package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards []string
	bid   int
	t     string
	rank  int
	pt    string
	prank int
}

var cardStrength = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var jcardStrength = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
	"J": 0,
}

var typeStrength = map[string]int{
	"5":  6,
	"4":  5,
	"FH": 4,
	"3":  3,
	"2P": 2,
	"1P": 1,
	"HC": 0,
}

var types = []string{
	"5",
	"4",
	"FH",
	"3",
	"2P",
	"1P",
	"HC",
}

func (h *Hand) winnings() int {
	return h.bid * h.rank
}

func (h *Hand) pwinnings() int {
	return h.bid * h.prank
}

func isStrongerC(c1 string, c2 string) bool {
	return cardStrength[c1] > cardStrength[c2]
}

func isStrongerT(t1 string, t2 string) bool {
	return typeStrength[t1] > typeStrength[t2]
}

func isPStrongerH(h1 Hand, h2 Hand) bool {
	if isStrongerT(h1.pt, h2.pt) {
		return true
	}

	if typeStrength[h1.pt] == typeStrength[h2.pt] {
		for i, c := range h1.cards {
			if jcardStrength[c] > jcardStrength[h2.cards[i]] {
				return true
			}

			if jcardStrength[c] == jcardStrength[h2.cards[i]] {
				continue
			}

			return false
		}
	}

	return false
}

func isStrongerH(h1 Hand, h2 Hand) bool {
	if isStrongerT(h1.t, h2.t) {
		return true
	}

	if typeStrength[h1.t] == typeStrength[h2.t] {
		for i, c := range h1.cards {
			if isStrongerC(c, h2.cards[i]) {
				return true
			}

			if cardStrength[c] == cardStrength[h2.cards[i]] {
				continue
			}

			return false
		}
	}

	return false
}

func getHandType(cards []string) string {
	matches := make(map[string]int)

	for i := 0; i < len(cards); i++ {
		if matches[cards[i]] == 0 {
			matches[cards[i]] = 1
			continue
		}

		matches[cards[i]] += 1
	}

	if len(matches) == 1 {
		return "5"
	}

	if len(matches) == 2 {
		for _, v := range matches {
			if v == 4 {
				return "4"
			}

			if v == 3 {
				return "FH"
			}
		}
	}

	if len(matches) == 3 {
		for _, v := range matches {
			if v == 3 {
				return "3"
			}

			if v == 2 {
				return "2P"
			}
		}
	}

	if len(matches) == 4 {
		return "1P"
	}

	return "HC"
}

func getPHandType(cards []string) string {
	matches := make(map[string]int)
	js := 0

	for i := 0; i < len(cards); i++ {
		if cards[i] == "J" {
			js += 1
			continue
		}

		if matches[cards[i]] == 0 {
			matches[cards[i]] = 1
			continue
		}

		matches[cards[i]] += 1
	}
	var maxK string
	for k, _ := range matches {
		if maxK == "" || matches[k] > matches[maxK] {
			maxK = k
		}
	}

	matches[maxK] += js

	if len(matches) == 1 {
		return "5"
	}

	if len(matches) == 2 {
		for _, v := range matches {
			if v == 4 {
				return "4"
			}

			if v == 3 {
				return "FH"
			}
		}
	}

	if len(matches) == 3 {
		for _, v := range matches {
			if v == 3 {
				return "3"
			}

			if v == 2 {
				return "2P"
			}
		}
	}

	if len(matches) == 4 {
		return "1P"
	}

	return "HC"
}

func main() {

	file, err := os.Open("inp.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hands = make(map[string][]Hand)
	var phands = make(map[string][]Hand)

	maxRank := 0
	pmaxRank := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		bid, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err)
		}

		cards := strings.Split(line[0], "")
		hand := Hand{
			cards: cards,
			bid:   bid,
			t:     getHandType(cards),
			pt:    getPHandType(cards),
		}

		hands[hand.t] = append(hands[hand.t], hand)
		phands[hand.pt] = append(phands[hand.pt], hand)
		maxRank += 1
		pmaxRank += 1
	}
	winnings := 0

	for _, t := range types {
		if len(hands[t]) == 0 {
			continue
		}
		hs := hands[t]

		sort.Slice(hs, func(i, j int) bool {
			return isStrongerH(hs[i], hs[j])
		})

		for i, _ := range hs {
			hs[i].rank = maxRank
			maxRank -= 1
		}

		hands[t] = hs
	}

	for _, t := range types {
		for _, h := range hands[t] {
			winnings += h.winnings()
		}
	}

	fmt.Println("Part 1", winnings)

	for _, t := range types {
		if len(phands[t]) == 0 {
			continue
		}
		hs := phands[t]
		sort.Slice(hs, func(i, j int) bool {
			return isPStrongerH(hs[i], hs[j])
		})

		for i, _ := range hs {
			hs[i].prank = pmaxRank
			pmaxRank -= 1
		}

		phands[t] = hs
	}

	winnings = 0

	for _, t := range types {
		for _, h := range phands[t] {
			winnings += h.pwinnings()
		}
	}

	fmt.Println("Part 2", winnings)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
