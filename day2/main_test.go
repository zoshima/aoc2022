package main

import "testing"

func Test(t *testing.T) {
	input := loadInput("input_test.txt")

	want := 15
	got := day1(input)

	if got != want {
		t.Fatalf("want %d, got %d", want, got)
	}

	want = 12
	got = day2(input)

	if got != want {
		t.Fatalf("want %d, got %d", want, got)
	}
}
