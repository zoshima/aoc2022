package main

import "testing"

func Test(t *testing.T) {
	input := loadInput("./input_test.txt")

	want := 24000
	got := part1(input)

	if got != want {
		t.Fatalf("want %d, got %d", want, got)
	}

	want = 45000
	got = part2(input)

	if got != want {
		t.Fatalf("want %d, got %d", want, got)
	}
}
