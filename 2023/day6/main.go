package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Record struct {
	lowHold  int
	highHold int
}

func parseLine(line string) []int {
	var result []int

	re := regexp.MustCompile(`\d+`)
	for _, e := range re.FindAllString(line, -1) {
		i, err := strconv.Atoi(e)
		if err != nil {
			fmt.Println(err)
		}

		result = append(result, i)
	}

	return result
}

func beatRecord(time int, distance int) Record {
	var record Record

	for hold := 1; hold <= time; hold++ {
		travelDistance := hold * (time - hold)

		if record.highHold != 0 && record.lowHold != 0 {
			break
		}

		if travelDistance > distance {
			record.lowHold = hold
			record.highHold = time - hold
		}

	}

	return record
}

func main() {

	input, err := os.ReadFile("inp.txt")
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(string(input), "\n")

	times := parseLine(string(lines[0]))
	distances := parseLine(string(lines[1]))

	ways := 1

	for i, time := range times {
		record := beatRecord(time, distances[i])
		ways *= record.highHold - record.lowHold + 1
	}

	fmt.Println("Part 1", ways)

	correctTime, err := strconv.Atoi(strings.Replace(strings.TrimPrefix(string(lines[0]), "Time:"), " ", "", -1))
	if err != nil {
		fmt.Println(err)
	}
	correctDistance, err := strconv.Atoi(strings.Replace(strings.TrimPrefix(string(lines[1]), "Distance:"), " ", "", -1))
	if err != nil {
		fmt.Println(err)
	}

	record := beatRecord(correctTime, correctDistance)

	fmt.Println("Part 2", record.highHold-record.lowHold+1)
}
