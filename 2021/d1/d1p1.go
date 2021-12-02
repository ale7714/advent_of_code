package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inc := 0
	var prev, curr int

	f, err := os.Open("input1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		curr, _ = strconv.Atoi(scanner.Text())

		if prev < curr {
			inc += 1
		}

		prev = curr
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(inc - 1)
}
