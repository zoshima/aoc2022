package main

import (
	"testing"
)

func Test(t *testing.T) {
	input := loadInput("./input_test.txt")

	want := 21
	got := part1(input)

	if got != want {
		t.Fatalf("p1: want %d, got %d", want, got)
	}

	want = 8
	got = part2(input)

	if got != want {
		t.Fatalf("p2: want %d, got %d", want, got)
	}
}
