package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func getRating(values []string, position int, m bool) string {
	if len(values) == 1 {
		return values[0]
	}
	var rating []string
	var bit string
	bits := map[string]int{"0": 0, "1": 0}

	for _, v := range values {
		bits[string(v[position])] += 1
	}

	if bits["0"] == bits["1"] {
		if m {
			bit = "1"
		} else {
			bit = "0"
		}
	} else {
		if m {
			if bits["0"] > bits["1"] {
				bit = "0"
			} else {
				bit = "1"
			}
		} else {
			if bits["0"] < bits["1"] {
				bit = "0"
			} else {
				bit = "1"
			}
		}
	}

	for _, v := range values {
		if bit == string(v[position]) {
			rating = append(rating, v)
		}
	}

	return getRating(rating, position+1, m)
}

func main() {
	var bits []map[string]int
	var gamma string
	var epsilon string

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		curr := scanner.Text()

		for i, bit := range curr {
			if len(bits) <= i {
				bits = append(bits, map[string]int{"0": 0, "1": 0})
			}
			bits[i][string(bit)] += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range bits {
		if v["0"] > v["1"] {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(g * e)

	var oxygen []string
	var co2 []string

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(f)

	for scanner.Scan() {
		curr := scanner.Text()

		if gamma[0] == curr[0] {
			oxygen = append(oxygen, curr)
		}

		if epsilon[0] == curr[0] {
			co2 = append(co2, curr)
		}
	}

	oxygenRating := getRating(oxygen, 1, true)
	co2Rating := getRating(co2, 1, false)

	o, _ := strconv.ParseInt(oxygenRating, 2, 64)
	c, _ := strconv.ParseInt(co2Rating, 2, 64)

	fmt.Println(o * c)

}
