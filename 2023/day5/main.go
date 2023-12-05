package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var categories = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

type Destination struct {
	value  int
	rangel int
}

func getLocation(seed int, maps map[string]map[int]Destination) int {
	index := seed
	for _, c := range categories {
		for source, destination := range maps[c] {
			if index >= source && index <= source+destination.rangel {
				delta := index - source
				index = destination.value + delta
				break
			}
		}
	}

	return index
}
func getMinLocation(seeds []int, maps map[string]map[int]Destination) (value int) {
	minLocation := math.MaxInt64

	for _, s := range seeds {
		location := getLocation(s, maps)
		if location < minLocation {
			minLocation = location
		}

	}

	return minLocation
}

func expandRanges(seeds []int, maps map[string]map[int]Destination) int {
	minLocation := math.MaxInt64

	for i := 0; i < len(seeds); i += 2 {
		start := seeds[i]
		rangel := seeds[i+1]
		for j := start; j < start+rangel; j++ {
			location := getLocation(j, maps)
			if location < minLocation {
				minLocation = location
			}
		}

	}

	return minLocation
}

func main() {

	file, err := os.Open("inp.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var seeds []int
	var maps = make(map[string]map[int]Destination)
	currentCategory := ""

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			re := regexp.MustCompile(`\d+`)

			for _, s := range re.FindAllString(line, -1) {
				n, _ := strconv.Atoi(s)
				seeds = append(seeds, n)
			}

			continue
		}

		if line == "" {
			currentCategory = ""
			continue
		}

		if strings.Contains(line, "-to-") {
			category := regexp.MustCompile(`(\w+)-to-(\w+)`).FindAllString(line, -1)[0]

			if len(maps[category]) == 0 {
				maps[category] = make(map[int]Destination)
			}

			currentCategory = category
		} else {
			ranges := regexp.MustCompile(`\d+`).FindAllString(line, -1)

			destination, _ := strconv.Atoi(ranges[0])
			source, _ := strconv.Atoi(ranges[1])
			rangel, _ := strconv.Atoi(ranges[2])

			maps[currentCategory][source] = Destination{value: destination, rangel: rangel}

		}
	}

	fmt.Println("Part 1: ", getMinLocation(seeds, maps))
	fmt.Println("Part 2:", expandRanges(seeds, maps))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
