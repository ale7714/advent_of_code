package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("inp.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		/* solution */

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
