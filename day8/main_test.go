package main

import (
	"testing"
)

func TestTraversal(t *testing.T) {
	input := loadInput("./input_test.txt")

	if input.TraverseNorth(3, 2) != 2 {
		t.Fatalf("TN: want 2, got %d", input.TraverseNorth(3, 2))
	}
	if input.TraverseWest(3, 2) != 2 {
		t.Fatalf("TW: want 2, got %d", input.TraverseWest(3, 2))
	}
	if input.TraverseSouth(3, 2) != 1 {
		t.Fatalf("TS: want 1, got %d", input.TraverseSouth(3, 2))
	}
	if input.TraverseEast(3, 2) != 2 {
		t.Fatalf("TE: want 2, got %d", input.TraverseEast(3, 2))
	}
}

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
