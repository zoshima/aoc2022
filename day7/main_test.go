package main

import (
	"testing"
)

func Test(t *testing.T) {
	input := loadInput("./input_test.txt")

	want := 95437
	got := part1(input)

	if got != want {
		t.Fatalf("p1: want %d, got %d", want, got)
	}

	want = 24933642
	got = part2(input)

	if got != want {
		t.Fatalf("p2: want %d, got %d", want, got)
	}
}
