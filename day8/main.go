package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type InputType [][]int

func (t *InputType) TraverseWest(row, col int) (int, bool) {
	var c int

	for i := col - 1; i >= 0; i-- {
		c++
		if (*t)[row][i] >= (*t)[row][col] {
			return c, false
		}
	}

	return c, true
}

func (t *InputType) TraverseEast(row, col int) (int, bool) {
	var c int

	for i := col + 1; i < len((*t)[row]); i++ {
		c++
		if (*t)[row][i] >= (*t)[row][col] {
			return c, false
		}
	}

	return c, true
}

func (t *InputType) TraverseNorth(row, col int) (int, bool) {
	var c int

	for i := row - 1; i >= 0; i-- {
		c++
		if (*t)[i][col] >= (*t)[row][col] {
			return c, false
		}
	}

	return c, true
}

func (t *InputType) TraverseSouth(row, col int) (int, bool) {
	var c int

	for i := row + 1; i < len((*t)); i++ {
		c++
		if (*t)[i][col] >= (*t)[row][col] {
			return c, false
		}
	}

	return c, true
}

func main() {
	input := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input InputType) int {
	res := len(input)*4 - 4

	numRows := len(input) - 1
	for row := 1; row < numRows; row++ {
		numCols := len(input[row]) - 1
		for col := 1; col < numCols; col++ {
			_, inf := input.TraverseNorth(row, col)
			if inf {
				res++
				continue
			}

			_, inf = input.TraverseWest(row, col)
			if inf {
				res++
				continue
			}

			_, inf = input.TraverseSouth(row, col)
			if inf {
				res++
				continue
			}

			_, inf = input.TraverseEast(row, col)
			if inf {
				res++
				continue
			}
		}
	}

	return res
}

func part2(input InputType) int {
	res := 0

	for row := 1; row < len(input)-1; row++ {
		for col := 1; col < len(input[row])-1; col++ {
			n, _ := input.TraverseNorth(row, col)
			w, _ := input.TraverseWest(row, col)
			s, _ := input.TraverseSouth(row, col)
			e, _ := input.TraverseEast(row, col)

			score := n * w * s * e

			if score < 0 {
				score *= -1
			}

			if score > res {
				res = score
			}
		}
	}

	return res
}

func loadInput(path string) InputType {
	file, _ := os.Open(path)
	defer file.Close()

	var input InputType

	cl := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, make([]int, len(line)))

		for i, c := range line {
			v, _ := strconv.Atoi(string(c))
			input[cl][i] = v
		}

		cl++
	}

	return input
}
