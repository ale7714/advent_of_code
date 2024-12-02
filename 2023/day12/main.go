package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "sync"
)

var cache = make(map[string]int)

func readFile(fname string) []string {
	var lines []string
	file, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func recursiveArrangements(springs string, groups []int) int {
	key := springs
	for _, group := range groups {
		key += strconv.Itoa(group) + ","
	}
	if v, ok := cache[key]; ok {
		return v
	}
	if len(springs) == 0 {
		if len(groups) == 0 {
			return 1
		} else {
			return 0
		}
	}
	if strings.HasPrefix(springs, "?") {
		return recursiveArrangements(strings.Replace(springs, "?", ".", 1), groups) +
			recursiveArrangements(strings.Replace(springs, "?", "#", 1), groups)
	}
	if strings.HasPrefix(springs, ".") {
		res := recursiveArrangements(strings.TrimPrefix(springs, "."), groups)
		cache[key] = res
		return res
	}

	if strings.HasPrefix(springs, "#") {
		if len(groups) == 0 {
			cache[key] = 0
			return 0
		}
		if len(springs) < groups[0] {
			cache[key] = 0
			return 0
		}
		if strings.Contains(springs[0:groups[0]], ".") {
			cache[key] = 0
			return 0
		}
		if len(groups) > 1 {
			if len(springs) < groups[0]+1 || string(springs[groups[0]]) == "#" {
				cache[key] = 0
				return 0
			}
			res := recursiveArrangements(springs[groups[0]+1:], groups[1:])
			cache[key] = res
			return res
		} else {
			res := recursiveArrangements(springs[groups[0]:], groups[1:])
			cache[key] = res
			return res
		}
	}

	return 0
}

func main() {
	INPUT := "inp.txt"
	// INPUT := "test_12.txt"

	fileContent := readFile(INPUT)
	sum_part1 := 0

	for _, line := range fileContent {
		var comb []int
		lineContent := strings.Split(line, " ")
		for _, Scomb := range strings.Split(lineContent[1], ",") {
			conv, _ := strconv.Atoi(Scomb)
			comb = append(comb, conv)
		}
		sum_part1 += recursiveArrangements(lineContent[0], comb)

	}
	fmt.Println("Part1:", sum_part1)
	sum_part2 := 0
	//	mutex := sync.RWMutex{}
	//	wg := sync.WaitGroup{}

	for _, line := range fileContent {
		//	wg.Add(1)
		// fmt.Println("Line: ", line)
		//	go func (line string) {
		var comb []int
		lineContent := strings.Split(line, " ")
		springs := ""
		for i := 0; i < 5; i++ {
			springs += lineContent[0]
			if i < 4 {
				springs += "?"
			}
		}
		for i := 0; i < 5; i++ {
			for _, Scomb := range strings.Split(lineContent[1], ",") {
				conv, _ := strconv.Atoi(Scomb)
				comb = append(comb, conv)
			}
		}
		//			mutex.RLock()
		//			mutex.Lock()
		sum_part2 += recursiveArrangements(springs, comb)
		//			mutex.Unlock()
		//			mutex.RUnlock()
		//			wg.Done()
		//	}(line)

	}

	//wg.Wait()
	fmt.Println("Part2:", sum_part2)
}
