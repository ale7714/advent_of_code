package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func sumMulInstructions(input string) int {
	// Regex to match valid `mul` instructions
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	submatches := regex.FindAllStringSubmatch(input, -1)

	sum := 0
	for _, match := range submatches {
		n, _ := strconv.Atoi(match[1])
		m, _ := strconv.Atoi(match[2])
		sum += n * m
	}

	return sum
}

func sumValidMulInstructions(input string) int {
	// Regex to match `mul`, `do()`, and `don't()` instructions
	regex := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	matches := regex.FindAllStringSubmatch(input, -1)

	mulEnabled := true
	sum := 0

	for _, match := range matches {
		switch match[1] {
		case "do()":
			mulEnabled = true
		case "don't()":
			mulEnabled = false
		default: // `mul` instructions
			if mulEnabled {
				n, _ := strconv.Atoi(match[2])
				m, _ := strconv.Atoi(match[3])
				sum += n * m
			}
		}
	}

	return sum
}

func parseInput(filePath string) string {
	// Read the file and join lines into a single string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	return sb.String()
}

func main() {
	// Specify the input file
	filePath := "input.txt"

	// Parse the input
	input := parseInput(filePath)

	sum := sumMulInstructions(input)
	log.Printf("Part 1: Sum of all `mul` instructions: %d\n", sum)

	sum = sumValidMulInstructions(input)
	log.Printf("Part 2: Sum of valid `mul` instructions: %d\n", sum)
}
