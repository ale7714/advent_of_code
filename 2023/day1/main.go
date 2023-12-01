package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func replaceNums(line string) string {
	line = strings.Replace(line, "one", "o1e", -1)
	line = strings.Replace(line, "two", "t2o", -1)
	line = strings.Replace(line, "three", "t3e", -1)
	line = strings.Replace(line, "four", "f4r", -1)
	line = strings.Replace(line, "five", "f5e", -1)
	line = strings.Replace(line, "six", "s6x", -1)
	line = strings.Replace(line, "seven", "s7n", -1)
	line = strings.Replace(line, "eight", "e8t", -1)
	line = strings.Replace(line, "nine", "n9e", -1)
	return line
}

func extractDigits(text string, spelled bool) int {
	if spelled {
		text = replaceNums(text)
	}

	var digits []int

	for _, ch := range text {
		if ch >= '0' && ch <= '9' {
			digit, _ := strconv.Atoi(string(ch))
			digits = append(digits, digit)
		}
	}

	return digits[0]*10 + digits[len(digits)-1]
}

func main() {

	file, err := os.Open("inp.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumCalibration1 := 0
	sumCalibration2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		value1 := extractDigits(line, false)
		sumCalibration1 += value1
		value2 := extractDigits(line, true)
		sumCalibration2 += value2

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: ", sumCalibration1)
	fmt.Println("Part 2: ", sumCalibration2)
}
