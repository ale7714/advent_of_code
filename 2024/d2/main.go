package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reports, err := parseReports(file)
	if err != nil {
		fmt.Println("Error parsing reports:", err)
		return
	}

	safeCount := 0
	for _, report := range reports {
		if isReportSafe(report) {
			safeCount++
		}
	}
	fmt.Printf("Number of safe reports: %d\n", safeCount)

	tolerantSafeCount := 0
	for _, report := range reports {
		if isReportSafe(report) || isReportSafeWithTolerance(report) {
			tolerantSafeCount++
		}
	}
	fmt.Printf("Number of safe reports with tolerance: %d\n", tolerantSafeCount)
}

func isReportSafe(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	increasing := numbers[1] > numbers[0]
	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		if diff == 0 || abs(diff) > 3 {
			return false
		}

		// Check if direction changes
		if i > 1 {
			currentIncreasing := numbers[i] > numbers[i-1]
			if currentIncreasing != increasing {
				return false
			}
		}
	}
	return true
}

func isReportSafeWithTolerance(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		// Create a new slice without the current number
		tolerantReport := append([]int{}, numbers[:i]...)         // Copy elements before i
		tolerantReport = append(tolerantReport, numbers[i+1:]...) // Copy elements after i

		if isReportSafe(tolerantReport) {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func parseReports(file *os.File) ([][]int, error) {
	var reports [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		var numbers []int
		for _, numStr := range strings.Fields(line) {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("error parsing number: %v", err)
			}
			numbers = append(numbers, num)
		}
		reports = append(reports, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return reports, nil

}
