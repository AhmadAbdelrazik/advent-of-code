package main

import (
	"fmt"
	"testing"
)

func TestCorporatePolicy(t *testing.T) {
	// t.Run("Incrementer", func(t *testing.T) {
	// 	tests := []string{
	// 		"abcdefgh",
	// 		"ghijklmn",
	// 		"hijklmmn",
	// 		"abbcegjk",
	// 		"zzz",
	// 	}
	// 	want := []string{
	// 		"abcdefgi",
	// 		"ghijklmo",
	// 		"hijklmmo",
	// 		"abbcegjl",
	// 		"aaaa",
	// 	}

	// 	for i := range tests {
	// 		t.Run(fmt.Sprintf("test_%v", tests[i]), func(t *testing.T) {
	// 			got := Incrementer(tests[i])
	// 			if got != want[i] {
	// 				t.Errorf("got %q want %q", got, want[i])
	// 			}
	// 		})
	// 	}
	// })

	t.Run("Valid", func(t *testing.T) {
		tests := []string{
			"hijklmmn",
			"abbceffg",
			"abbcegjk",
			"abcdffaa",
			"ghjaabcc",
		}		
		want := []bool{
			false,
			false,
			false,
			true,
			true,
		}

		for i := range tests {
			t.Run(fmt.Sprintf("testing_%v", tests[i]), func(t *testing.T) {
				got := valid(tests[i])
				
				if got != want[i] {
					t.Errorf("got '%v' want '	%v'", got, want[i])
				}
			})
		}
	})

	t.Run("Part1Rules", func(t *testing.T) {
		tests := []string{
			"abcdefgh",
			"ghijklmn",
		}
		want := []string{
			"abcdffaa",
			"ghjaabcc",
		}

		for i := range tests {
			t.Run(fmt.Sprintf("test_%v", tests[i]), func(t *testing.T) {
				got := Part1(tests[i])
				
				if got != want[i] {
					t.Errorf("got %q want %q", got, want[i])
				}
			})
		}
	})
}