package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	input := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input [][]int) int {
	m := 0

	for _, cl := range input {
		t := 0
		for _, c := range cl {
			t += c
		}

		if t > m {
			m = t
		}
	}

	return m
}

func part2(input [][]int) int {
	m1, m2, m3 := 0, 0, 0

	for _, cl := range input {
		t := 0
		for _, c := range cl {
			t += c
		}

		if t > m1 {
			m3, m2, m1 = m2, m1, t
		} else if t > m2 {
			m3, m2 = m2, t
		} else if t > m3 {
			m3 = t
		}
	}

	return m1 + m2 + m3
}

func loadInput(path string) [][]int {
	file, _ := os.Open(path)
	defer file.Close()

	input := make([][]int, 1)
	ci := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			input = append(input, []int{})
			ci++
		} else {
			cal, _ := strconv.Atoi(str)
			input[ci] = append(input[ci], cal)
		}
	}

	return input
}
