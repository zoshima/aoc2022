package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Motion struct {
	Dir rune
	Num int
}

type Position struct {
	X int
	Y int
}

func (p *Position) Move(dir rune) {
	switch dir {
	case 'R':
		p.X += 1
	case 'D':
		p.Y += 1
	case 'U':
		p.Y -= 1
	case 'L':
		p.X -= 1
	}
}

func (a *Position) MoveTo(b Position) bool {
	dx := diff(a.X, b.X)
	dy := diff(a.Y, b.Y)

	moveX, moveY := false, false

	if dx == 2 {
		moveX = true

		if dy == 1 {
			moveY = true
		}
	}

	if dy == 2 {
		moveY = true

		if dx == 1 {
			moveX = true
		}
	}

	if moveX {
		if a.X > b.X {
			a.X -= 1
		} else {
			a.X += 1
		}
	}

	if moveY {
		if a.Y > b.Y {
			a.Y -= 1
		} else {
			a.Y += 1
		}
	}

	return moveX || moveY
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}

	return a - b
}

type InputType []Motion

func main() {
	input := loadInput("input.txt")

	log.Println(part1(input))
	log.Println(part2(input))
}

func part1(input InputType) int {
	head, tail := Position{}, Position{}
	visited := make(map[Position]any)

	for _, motion := range input {
		for i := 0; i < motion.Num; i++ {
			head.Move(motion.Dir)
			if ok := tail.MoveTo(head); ok {
				visited[tail] = nil
			}
		}
	}

	return len(visited) + 1
}

func part2(input InputType) int {
	numKnots := 10
	knots := make([]Position, numKnots)
	visited := make(map[Position]any)

	for _, motion := range input {
		for i := 0; i < motion.Num; i++ {
			knots[0].Move(motion.Dir) // head

		KnotsLoop:
			for j := 1; j < numKnots; j++ {
				if !knots[j].MoveTo(knots[j-1]) {
					break KnotsLoop
				}

				if j == numKnots-1 {
					visited[knots[j]] = nil
				}
			}
		}
	}

	return len(visited) + 1
}

func loadInput(path string) InputType {
	file, _ := os.Open(path)
	defer file.Close()

	var input InputType

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		dir := rune(parts[0][0])
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		input = append(input, Motion{dir, num})
	}

	return input
}
