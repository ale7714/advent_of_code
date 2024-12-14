package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Machine struct {
	Ax, Ay, Bx, By, X_target, Y_target int
}

func main() {
	machines, err := parseInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	totalCost := 0
	prizesWon := 0

	for _, machine := range machines {
		_, _, cost, solvable := solveCramersRule(machine)
		if solvable {
			//fmt.Printf("Machine %d: Press A %d times, B %d times. Minimum tokens: %d.\n", i+1, a, b, cost)
			totalCost += cost
			prizesWon++
		} else {
			//fmt.Printf("Machine %d: No solution exists.\n", i+1)
		}
	}

	fmt.Printf("\nTotal tokens: %d\n", totalCost)
	fmt.Printf("Prizes won: %d\n", prizesWon)
}

func solveCramersRule(machine Machine) (int, int, int, bool) {
	denominator := machine.Ax*machine.By - machine.Ay*machine.Bx
	if denominator == 0 {
		return 0, 0, 0, false
	}

	A := float64(machine.X_target*machine.By-machine.Y_target*machine.Bx) / float64(denominator)
	B := float64(machine.Ax*machine.Y_target-machine.Ay*machine.X_target) / float64(denominator)

	if A < 0 || B < 0 || A != float64(int(A)) || B != float64(int(B)) {
		return 0, 0, 0, false
	}

	a := int(A)
	b := int(B)
	cost := 3*a + b
	return a, b, cost, true
}

func parseInput(fileName string) ([]Machine, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []Machine
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	for i := 0; i < len(lines); i += 3 {
		var Ax, Ay, Bx, By, X_target, Y_target int

		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &Ax, &Ay)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &Bx, &By)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &X_target, &Y_target)
		X_target += 10000000000000
		Y_target += 10000000000000
		machines = append(machines, Machine{Ax, Ay, Bx, By, X_target, Y_target})
	}

	return machines, nil
}
