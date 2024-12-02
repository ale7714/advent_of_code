/*
--- Day 11: Cosmic Expansion ---
You continue following signs for "Hot Springs" and eventually come across
an observatory. The Elf within turns out to be a researcher studying cosmic
expansion using the giant telescope here.

He doesn't know anything about the missing machine parts; he's only visiting for
this research project. However, he confirms that the hot springs are the next-closest
area likely to have people; he'll even take you straight there once he's done
with today's observation analysis.

Maybe you can help him with the analysis to speed things up?

The researcher has collected a bunch of data and compiled the data into a single
giant image (your puzzle input).
The image includes empty space (.) and galaxies (#). For example:

...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
The researcher is trying to figure out the sum of the lengths of the shortest path
between every pair of galaxies. However, there's a catch: the universe expanded
in the time it took the light from those galaxies to reach the observatory.

Due to something involving gravitational effects, only some space expands.
In fact, the result is that any rows or columns that contain no galaxies should
all actually be twice as big.

In the above example, three columns and two rows contain no galaxies:

   v  v  v
 ...#......
 .......#..
 #.........
>..........<
 ......#...
 .#........
 .........#
>..........<
 .......#..
 #...#.....
   ^  ^  ^

These rows and columns need to be twice as big; the result of cosmic expansion
therefore looks like this:

....#........
.........#...
#............
.............
.............
........#....
.#...........
............#
.............
.............
.........#...
#....#.......

Equipped with this expanded universe, the shortest path between every pair of
galaxies can be found. It can help to assign every galaxy a unique number:

....1........
.........2...
3............
.............
.............
........4....
.5...........
............6
.............
.............
.........7...
8....9.......

In these 9 galaxies, there are 36 pairs. Only count each pair once;
order within the pair doesn't matter. For each pair, find any shortest
path between the two galaxies using only steps that move up, down, left,
or right exactly one . or # at a time. (The shortest path between
two galaxies is allowed to pass through another galaxy.)

For example, here is one of the shortest paths between galaxies 5 and 9:

....1........
.........2...
3............
.............
.............
........4....
.5...........
.##.........6
..##.........
...##........
....##...7...
8....9.......
This path has length 9 because it takes a minimum of nine steps to get from galaxy
5 to galaxy 9 (the eight locations marked # plus the step onto galaxy 9 itself).
Here are some other example shortest path lengths:

Between galaxy 1 and galaxy 7: 15
Between galaxy 3 and galaxy 6: 17
Between galaxy 8 and galaxy 9: 5

In this example, after expanding the universe, the sum of the shortest
path between all 36 pairs of galaxies is 374.

Expand the universe, then find the length of the shortest path between
every pair of galaxies. What is the sum of these lengths?

--- Part Two ---
The galaxies are much older (and thus much farther apart)
than the researcher initially estimated.

Now, instead of the expansion you did before, make each empty row or column one million times larger.
That is, each empty row should be replaced with 1000000 empty rows,
and each empty column should be replaced with 1000000 empty columns.

(In the example above, if each empty row or column were merely 10 times larger,
the sum of the shortest paths between every pair of galaxies would be 1030.
If each empty row or column were merely 100 times larger, the sum of the shortest
paths between every pair of galaxies would be 8410. However, your universe will
need to expand far beyond these values.)

Starting with the same initial image, expand the universe according to these new rules,
then find the length of the shortest path between every pair of galaxies.
What is the sum of these lengths?
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func Day11() [2]int {
	return [2]int{
		d11p1(),
		d11p2(),
	}
}

func main() {
	fmt.Println(Day11())
}

func GetShortestDistordedDistance(distortion int) int {
	file, err := os.Open("inp.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	//get all data from input file
	starChart := [][]rune{}
	for scanner.Scan() {
		starChart = append(starChart, []rune(scanner.Text()))
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	//find all rows that doesn't have a '#'
	rowsWithoutHash := []int{}
	for y := 0; y < len(starChart); y++ {
		found := false
		for x := 0; x < len(starChart[y]); x++ {
			if starChart[y][x] == '#' {
				found = true
				break
			}
		}
		if !found {
			rowsWithoutHash = append(rowsWithoutHash, y)
		}
	}

	//find all columns that doesn't have a '#'
	colsWithoutHash := []int{}
	for x := 0; x < len(starChart[0]); x++ {
		found := false
		for y := 0; y < len(starChart); y++ {
			if starChart[y][x] == '#' {
				found = true
				break
			}
		}
		if !found {
			colsWithoutHash = append(colsWithoutHash, x)
		}
	}

	//get the distoreded coordinates of all # symbols
	starPos := [][2]int{}
	for y := 0; y < len(starChart); y++ {
		for x := 0; x < len(starChart[y]); x++ {
			if starChart[y][x] == '#' {
				cumulDistortionX := 0
				cumulDistortionY := 0

				for i := 0; i < len(rowsWithoutHash); i++ {
					if y > rowsWithoutHash[i] {
						cumulDistortionY += distortion - 1
					}
				}

				for i := 0; i < len(colsWithoutHash); i++ {
					if x > colsWithoutHash[i] {
						cumulDistortionX += distortion - 1
					}
				}

				starPos = append(starPos, [2]int{x + cumulDistortionX, y + cumulDistortionY})
			}
		}
	}

	//get the sum of shortest path
	sum := 0
	for i := 0; i < len(starPos)-1; i++ {
		for j := i + 1; j < len(starPos); j++ {
			sum += ManhattanDistance(starPos[i], starPos[j])
		}
	}
	return sum
}

func d11p1() int {
	return GetShortestDistordedDistance(2)
}

func d11p2() int {
	return GetShortestDistordedDistance(1000000)
}

// calculate the Manhattan Distance between A and B and return the distance
func ManhattanDistance(posA [2]int, posB [2]int) int {
	return int(math.Abs(float64(posB[0])-float64(posA[0])) + math.Abs(float64(posB[1])-float64(posA[1])))
}
