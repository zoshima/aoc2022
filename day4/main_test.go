package main

import (
	"testing"
)

func Test(t *testing.T) {
	input := loadInput("./input_test.txt")

	want := 2
	got := part1(input)

	if got != want {
		t.Fatalf("p1: want %d, got %d", want, got)
	}

	want = 4
	got = part2(input)

	if got != want {
		t.Fatalf("p2: want %d, got %d", want, got)
	}
}
