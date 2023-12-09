package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Measure struct {
	sequence    []int
	predictionl int
	predictionr int
}

func allZeros(a []int) bool {
	for _, v := range a {
		if v != 0 {
			return false
		}
	}
	return true
}

func prediction(m Measure) (int, int) {
	diffs := make([][]int, 0)
	diffs = append(diffs, m.sequence)
	old := m.sequence
	for {
		ns := make([]int, 0)

		for i, n := range old {
			if i == len(old)-1 {
				break
			}
			ns = append(ns, old[i+1]-n)
		}

		diffs = append(diffs, ns)
		old = ns
		if allZeros(ns) {
			break
		}
	}
	var leftValue, rightValue int

	for i := len(diffs) - 1; i >= 0; i -= 1 {

		if i < 0 {
			break
		}

		s := diffs[i]

		if allZeros(s) {
			leftValue = 0
			rightValue = 0
			nd := append(diffs[i], 0)
			nd = append([]int{0}, nd...)
			diffs[i] = nd

			continue
		}

		p := diffs[i]
		newl := leftValue + p[len(p)-1]
		newr := p[0] - rightValue
		nd := append(diffs[i], newl)
		nd = append([]int{newr}, nd...)
		diffs[i] = nd
		leftValue = newl
		rightValue = newr

	}
	fmt.Println(diffs)
	return leftValue, rightValue
}
func main() {

	file, err := os.Open("inp.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	measures := make([]Measure, 0)
	sumPredictionsR := 0
	sumPredictionsL := 0

	for scanner.Scan() {
		line := scanner.Text()
		var sequence []int
		for _, n := range regexp.MustCompile(`-?\d+`).FindAllString(line, -1) {
			num, _ := strconv.Atoi(n)
			sequence = append(sequence, num)
		}
		m := Measure{sequence: sequence}
		m.predictionl, m.predictionr = prediction(m)
		sumPredictionsR += m.predictionr
		sumPredictionsL += m.predictionl
		measures = append(measures, m)
	}

	fmt.Println(sumPredictionsL)
	fmt.Println(sumPredictionsR)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
