package main

import (
	"fmt"
	"testing"
)

func TestVisitedHouses(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		tests := []string{
			"^v^v^v^v^v",
			">",
			"^>v<",
		}

		results := []int{
			2,
			2,
			4,
		}

		for i, test := range tests {
			t.Run(fmt.Sprintf("test:%q", test), func(t *testing.T) {
				got := VisitedHouses(test)
				want := results[i]
	
				assertEquality(t, got, want)
			})
		}

	})
}

func assertEquality(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
