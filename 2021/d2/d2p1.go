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
	x, depth := 0, 0

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
		case "down":
			depth += units
		case "up":
			depth -= units
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(depth * x)
}
