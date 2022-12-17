package main

import (
	"testing"
)

func Test(t *testing.T) {
	input := loadInput("./input_test.txt")

	want := 13
	got := part1(input)

	if got != want {
		t.Fatalf("p1: want %d, got %d", want, got)
	}

	want = 1
	got = part2(input)

	if got != want {
		t.Fatalf("p2: want %d, got %d", want, got)
	}

	input2 := InputType{
		{'R', 5},
		{'U', 8},
		{'L', 8},
		{'D', 3},
		{'R', 17},
		{'D', 10},
		{'L', 25},
		{'U', 20},
	}
	want = 36
	got = part2(input2)

	if got != want {
		t.Fatalf("p2: want %d, got %d", want, got)
	}
}
