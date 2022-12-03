package main

import (
	"bufio"
	"log"
	"os"
)

type InputType []string

func main() {
	input := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input InputType) int {
	res := 0

	for _, sack := range input {
		shared := make(map[rune]bool)

		for i, r := range sack {
			if i < len(sack)/2 {
				shared[r] = false
			} else {
				if v, ok := shared[r]; ok && !v {
					shared[r] = true
					res += runeToInt(r)
				}
			}
		}
	}

	return res
}

func part2(input InputType) int {
	res := 0

groupsLoop:
	for i := 0; i < len(input); i += 3 {
		shared := make(map[rune]bool)

		for _, r := range input[i] {
			shared[r] = false
		}

		for _, r := range input[i+1] {
			if _, ok := shared[r]; ok {
				shared[r] = true
			}
		}

		for _, r := range input[i+2] {
			if v, ok := shared[r]; ok && v {
				res += runeToInt(r)
				continue groupsLoop
			}
		}
	}

	return res
}

func runeToInt(r rune) int {
	if r <= 'Z' {
		return int(r) - int('A') + 27
	}

	return int(r) - int('a') + 1
}

func loadInput(path string) InputType {
	file, _ := os.Open(path)
	defer file.Close()

	var input InputType

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	return input
}
