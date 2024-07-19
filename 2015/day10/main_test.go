package main

import (
	"fmt"
	"testing"
)

func TestElves(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		tests := []string{
			"1",
			"11",
			"21",
			"1211",
			"111221",
		}
		want := []string{
			"11",
			"21",
			"1211",
			"111221",
			"312211",
		}

		for i := range tests {
			t.Run(fmt.Sprintf("test_%v",tests[i]), func(t *testing.T) {
				got := Part1(tests[i], 1)
				if got != len(want[i]) {
					t.Fatalf("got '%v' want '%v'", got, want[i])
				}
			})
		} 
	})
}
