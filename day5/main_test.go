package main

import (
	"testing"
)

func Test(t *testing.T) {
	input := loadInput("./input_test.txt")
	input2 := loadInput("./input_test.txt")

	want := "CMZ"
	got := part1(input)

	if got != want {
		t.Fatalf("p1: want %v, got %v", want, got)
	}

	want = "MCD"
	got = part2(input2)

	if got != want {
		t.Fatalf("p2: want %v, got %v", want, got)
	}
}
