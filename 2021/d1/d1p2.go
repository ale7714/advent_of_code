package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {
	inc := 0
	var w1, w2 []int

	f, err := os.Open("input1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		curr, _ := strconv.Atoi(scanner.Text())

		if len(w1) < 1 {
			w1 = append(w1, curr)
		} else {
			if len(w1) < 3 {
				w1 = append(w1, curr)
			}
			if len(w2) < 3 {
				w2 = append(w2, curr)
			}
		}

		if len(w1) == len(w2) && len(w2) == 3 {
			if sum(w1) < sum(w2) {
				inc += 1
			}
			w1 = w2
			_, w2 = w2[0], w2[1:]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(inc)
}
