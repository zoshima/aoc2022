package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input []int) int {
	x, res := 1, 0
	var def1, def2, numCycles int

	isBreakpoint := func(c int) bool {
		return c == 20 || c == 60 || c == 100 || c == 140 || c == 180 || c == 220
	}

	cycle := func(v int) {
		numCycles++
		x += def1
		def1, def2 = def2, v

		if isBreakpoint(numCycles) {
			res += x * numCycles
		}
	}

	for i := 0; i < len(input); i++ {
		cycle(input[i])

		if input[i] == 0 {
			continue
		}

		// additional cycle
		cycle(0)
	}

	x += def1 + def2

	return res
}

func part2(input []int) int {
	return -1
}

func loadInput(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	input := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) == 1 {
			input = append(input, 0)
			continue
		}

		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		input = append(input, value)
	}

	return input
}
