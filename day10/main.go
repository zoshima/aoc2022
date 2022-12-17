package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type InputType struct {
	NumCycles  int
	Operations []int
}

func main() {
	input := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input InputType) int {
	buf := -1 // index of buffered operation
	x := 1

	for i := 0; i < input.NumCycles; i++ {
		if buf != -1 {
			x += input.Operations[buf]
		}

		buf++
	}

	return x
}

func part2(input InputType) int {
	return 0
}

func loadInput(path string) InputType {
	file, _ := os.Open(path)
	defer file.Close()

	var input InputType

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		v := 0

		if line != "noop" {
			sv := line[len("addx "):]
			v, _ = strconv.Atoi(sv)

			input.NumCycles += 2
		} else {
			input.NumCycles++
		}

		input.Operations = append(input.Operations, v)
	}

	return input
}
