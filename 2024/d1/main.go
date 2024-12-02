package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func insertSorted(slice []int, value int) []int {
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= value
	})

	slice = append(slice, 0)
	copy(slice[index+1:], slice[index:])
	slice[index] = value
	return slice
}

func parseFile(file *os.File) ([]int, []int, error) {
	var column1 []int
	var column2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, err
		}
		column1 = insertSorted(column1, num1)
		column2 = insertSorted(column2, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return column1, column2, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	list1, list2, err := parseFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	totalDistance := 0
	similarityScore := 0
	for i, num1 := range list1 {
		distance := num1 - list2[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance

		count := 0
		lastIndex := 0

		for j := lastIndex; j < len(list2); j++ {
			if num1 == list2[j] {
				count++
			}

			if list2[j] > num1 {
				lastIndex = j
				break
			}
		}

		similarityScore += num1 * count
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
	fmt.Printf("Similarity score: %d\n", similarityScore)
}
