package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	Win      = 6
	Draw     = 3
	Loss     = 0
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

func main() {
	input := loadInput("input.txt")

	log.Println(day1(input))
	log.Println(day2(input))
}

func day1(input [][2]string) int {
	res := 0

	for _, result := range input {
		a := charToMotion(result[0])
		b := charToMotion(result[1])

		res += b

		if a == b {
			res += Draw
		} else if a == Rock {
			if b == Paper {
				res += Win
			} else {
				res += Loss
			}
		} else if a == Paper {
			if b == Scissors {
				res += Win
			} else {
				res += Loss
			}
		} else if b == Rock {
			res += Win
		} else {
			res += Loss
		}
	}

	return res
}

func day2(input [][2]string) int {
	m := map[int]int{
		7: Paper,
		4: Rock,
		1: Scissors,
		8: Scissors,
		5: Paper,
		2: Rock,
		9: Rock,
		6: Scissors,
		3: Paper,
	}

	res := 0

	for _, result := range input {
		a := charToMotion(result[0])
		b := charToOutcome(result[1])

		res += m[a+b] + b
	}

	return res
}

func loadInput(path string) [][2]string {
	file, _ := os.Open(path)
	defer file.Close()

	input := make([][2]string, 0)
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, [2]string{})

		str := scanner.Text()
		res := strings.SplitN(str, " ", 2)

		input[i][0] = res[0]
		input[i][1] = res[1]

		i++
	}

	return input
}

func charToMotion(s string) int {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	}

	return Scissors
}

func charToOutcome(s string) int {
	switch s {
	case "X":
		return Loss
	case "Y":
		return Draw
	}

	return Win
}
