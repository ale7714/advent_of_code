package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func makeBoards(boardsContent []string) [][][]string {
	var boards [][][]string
	for _, content := range boardsContent {
		board := make([][]string, 5)
		for i := range board {
			board[i] = make([]string, 5)
		}

		rows := strings.Split(content, "\n")
		for i, row := range rows {
			for j, v := range regexp.MustCompile(`\s+`).Split(strings.TrimSpace(row), -1) {
				board[i][j] = v
			}
		}

		boards = append(boards, board)
	}
	return boards
}

func markBoard(board [][]string, number string) [][]string {
	var x, y int
	found := false
	for i, row := range board {
		for j, val := range row {
			if val == number {
				x, y = i, j
				found = true
				break
			}
		}
	}

	if found {
		board[x][y] = "x"
	}

	return board
}

func winningRow(row []string) bool {
	win := true

	for _, v := range row {
		if v != "x" {
			win = false
			break
		}
	}

	return win
}

func transpose(a [][]string) [][]string {
	newArr := make([][]string, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}

func winningBoard(board [][]string) bool {
	win := false

	for _, row := range board {
		if winningRow(row) {
			win = true
			break
		}
	}
	return win
}

func getScore(board [][]string, number string) int {
	var sum int
	n, _ := strconv.Atoi(number)

	for _, row := range board {
		for _, val := range row {
			if val != "x" {
				v, _ := strconv.Atoi(val)
				sum += v
			}
		}
	}
	return sum * n
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n\n")
	numbers, boards := strings.Split(lines[0], ","), makeBoards(lines[1:])

	var winBoards [][][]string
	var winningNumbers []int
	for _, board := range boards {
		for j, number := range numbers {
			markedBoard := markBoard(board, number)
			if winningBoard(markedBoard) || winningBoard(transpose(markedBoard)) {
				winBoards = append(winBoards, markedBoard)
				winningNumbers = append(winningNumbers, j)
				break
			}
		}
	}

	min := winningNumbers[0]
	var index int
	for i, v := range winningNumbers {
		if v < min {
			min = v
			index = i
		}
	}

	fmt.Println(getScore(winBoards[index], numbers[min]))

	max := winningNumbers[0]
	for i, v := range winningNumbers {
		if v > max {
			max = v
			index = i
		}
	}

	fmt.Println(getScore(winBoards[index], numbers[max]))
}
