package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Operation struct {
	From int
	To   int
	Num  int
}

type Stack []rune

func (s *Stack) Prepend(r rune) {
	*s = append([]rune{r}, *s...)
}

func (s *Stack) Push(r Stack) {
	*s = append(*s, r...)
}

func (s *Stack) Pop(n int) Stack {
	i := len(*s) - n
	r := (*s)[i:]
	*s = (*s)[:i]

	return r
}

type InputType struct {
	Stacks     []Stack
	Operations []Operation
}

func main() {
	input := loadInput("input.txt")
	input2 := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input2))
}

func part1(input InputType) string {
	for _, operation := range input.Operations {
		for i := 0; i < operation.Num; i++ {
			r := input.Stacks[operation.From-1].Pop(1)
			input.Stacks[operation.To-1].Push(r)
		}
	}

	res := ""

	for _, stack := range input.Stacks {
		res += string(stack.Pop(1))
	}

	return res
}

func part2(input InputType) string {
	for _, operation := range input.Operations {
		r := input.Stacks[operation.From-1].Pop(operation.Num)
		input.Stacks[operation.To-1].Push(r)
	}

	res := ""

	for _, stack := range input.Stacks {
		res += string(stack.Pop(1))
	}

	return res
}

func loadInput(path string) InputType {
	file, _ := os.Open(path)
	defer file.Close()

	input := InputType{}

	stackPattern := regexp.MustCompile(`[A-Z]{1}|\ \ \ \ `)
	digitsPattern := regexp.MustCompile(`\d{1,2}`)

	scanner := bufio.NewScanner(file)

	isOperation := false

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 || line[1] == '1' {
			isOperation = true
			continue
		}

		if !isOperation {
			matches := stackPattern.FindAllString(line, -1)

			if len(input.Stacks) == 0 {
				input.Stacks = make([]Stack, len(matches))
			}

			for i, match := range matches {
				if strings.TrimSpace(match) != "" {
					input.Stacks[i].Prepend(rune(match[0]))
				}
			}
		} else {
			matches := digitsPattern.FindAllString(line, -1)

			num, _ := strconv.Atoi(matches[0])
			from, _ := strconv.Atoi(matches[1])
			to, _ := strconv.Atoi(matches[2])

			input.Operations = append(input.Operations, Operation{from, to, num})
		}
	}

	return input
}
