package main

import (
	"fmt"
	"testing"
)

func TestNaughtyWords(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		tests := []string{
			"ugknbfddgicrmopn",
			"aaa",
			"jchzalrnumimnmhp",
			"haegwjzuvuyypxyu",
			"dvszwmarrgswjxmb",
		}

		answers := []bool{
			true,
			true,
			false,
			false,
			false,
		}

		for i, test := range tests {
			t.Run(fmt.Sprintf("test:%q",test), func(t *testing.T) {
				got := Part1(test)
				want := answers[i]

				assertEquality(t, got, want)
			})
		}
	})

	t.Run("Part2", func(t *testing.T) {
		tests := []string{
			"qjhvhtzxzqqjkmpb",
			"xxyxx",
			"aabcdefgaa",
			"uurcxstgmygtbstg",
			"ieodomkazucvgmuy",
			"aaa",
			"xilodxfuxphuiiii",
			"pzkkkkwrlvxiuysn",
			"bkkkkcwegvypbrio",
		}

		answers := []bool{
			true,
			true,
			false,
			false,
			false,
			false,
			true,
			true,
			true,
		}

		for i, test := range tests {
			t.Run(fmt.Sprintf("test:%q",test), func(t *testing.T) {
				got := Part2(test)
				want := answers[i]

				assertEquality(t, got, want)
			})
		}

	})
}

func assertEquality(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}