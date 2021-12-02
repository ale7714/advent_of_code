package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	x, depth, aim := 0, 0, 0

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		curr := strings.Fields(scanner.Text())
		command := curr[0]
		units, _ := strconv.Atoi(curr[1])

		switch command {
		case "forward":
			x += units
			depth = depth + (aim * units)
		case "down":
			aim += units
		case "up":
			aim -= units
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(depth * x)
}
