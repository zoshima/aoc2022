package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		Input string
		Want  int
	}{
		{
			Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			Want:  7,
		},
		{
			Input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			Want:  5,
		},
		{
			Input: "nppdvjthqldpwncqszvftbrmjlhg",
			Want:  6,
		},
		{
			Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			Want:  10,
		},
		{
			Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			Want:  11,
		},
	}

	for _, tc := range testCases {
		got := part1(tc.Input)

		if tc.Want != got {
			t.Fatalf("want %d, got %d", tc.Want, got)
		}
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		Input string
		Want  int
	}{
		{
			Input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			Want:  19,
		},
		{
			Input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			Want:  23,
		},
		{
			Input: "nppdvjthqldpwncqszvftbrmjlhg",
			Want:  23,
		},
		{
			Input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			Want:  29,
		},
		{
			Input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			Want:  26,
		},
	}

	for _, tc := range testCases {
		got := part2(tc.Input)

		if tc.Want != got {
			t.Fatalf("want %d, got %d", tc.Want, got)
		}
	}
}
