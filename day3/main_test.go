package main

import (
	"testing"
)

func Test(t *testing.T) {
	input := loadInput("./input_test.txt")

	want := 157
	got := part1(input)

	if got != want {
		t.Fatalf("p1: want %d, got %d", want, got)
	}

	want = 70
	got = part2(input)

	if got != want {
		t.Fatalf("p2: want %d, got %d", want, got)
	}
}

func TestRuneToInt(t *testing.T) {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, r := range alphabet {
		got := runeToInt(r)
		if i+1 != got {
			t.Fatalf("%q: want %d, got %d", r, i+1, got)
		}
	}
}
