package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

type Pair struct {
	First  Range
	Second Range
}

type InputType []Pair

func main() {
	input := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input InputType) int {
	res := 0

	for _, pair := range input {
		if (pair.First.Start >= pair.Second.Start && pair.First.End <= pair.Second.End) ||
			(pair.Second.Start >= pair.First.Start && pair.Second.End <= pair.First.End) {
			res++
		}
	}

	return res
}

func part2(input InputType) int {
	res := 0

	for _, pair := range input {
		if pair.First.Start <= pair.Second.End &&
			pair.Second.Start <= pair.First.End {
			res++
		}
	}

	return res
}

func loadInput(path string) InputType {
	file, _ := os.Open(path)
	defer file.Close()

	var input InputType

	parseRange := func(s string) Range {
		n := strings.Split(s, "-")
		i, _ := strconv.Atoi(n[0])
		j, _ := strconv.Atoi(n[1])

		return Range{i, j}
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")

		input = append(input, Pair{
			First:  parseRange(ranges[0]),
			Second: parseRange(ranges[1]),
		})
	}

	return input
}
