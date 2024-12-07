package main

import (
	"bufio"
	"fmt"
	"math"
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

	part1Total := 0
	part2Total := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		target, nums := parseLine(line)
		if nums == nil {
			continue
		}

		if canFormTarget(nums, target) {
			part1Total += target
		}
		if canFormTargetAllOperators(nums, target) {
			part2Total += target
		}
	}

	fmt.Printf("Total calibration: %d\n", part1Total)
	fmt.Printf("Total calibration with all operators: %d\n", part2Total)
}

func parseLine(line string) (int, []int) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return 0, nil
	}

	targetStr := strings.TrimSpace(parts[0])
	target, err := strconv.Atoi(targetStr)
	if err != nil {
		return 0, nil
	}

	numsStr := strings.Fields(strings.TrimSpace(parts[1]))
	nums := make([]int, len(numsStr))
	for i, ns := range numsStr {
		n, err := strconv.Atoi(ns)
		if err != nil {
			return 0, nil
		}
		nums[i] = n
	}
	return target, nums
}

func canFormTarget(nums []int, target int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}

	slots := len(nums) - 1
	combinations := 1 << slots

	for mask := 0; mask < combinations; mask++ {
		val := nums[0]
		for i := 0; i < slots; i++ {
			op := (mask >> i) & 1
			if op == 0 {
				val = val + nums[i+1]
			} else {
				val = val * nums[i+1]
			}
		}

		if val == target {
			return true
		}
	}

	return false
}

func canFormTargetAllOperators(nums []int, target int) bool {
	if len(nums) == 1 {
		return nums[0] == target
	}

	slots := len(nums) - 1
	combinations := int(math.Pow(3, float64(slots)))

	for mask := 0; mask < combinations; mask++ {
		val := nums[0]
		currentMask := mask
		for i := 0; i < slots; i++ {
			op := currentMask % 3
			currentMask /= 3

			nextNum := nums[i+1]
			switch op {
			case 0:
				val = val + nextNum
			case 1:
				val = val * nextNum
			case 2:
				val = concatNumbers(val, nextNum)
			}
		}

		if val == target {
			return true
		}
	}

	return false
}

func concatNumbers(a, b int) int {
	if b == 0 {
		return a * 10
	}

	digits := 0
	temp := b
	for temp > 0 {
		temp /= 10
		digits++
	}
	return a*int(math.Pow10(digits)) + b
}
