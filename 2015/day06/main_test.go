package main

import "testing"

func TestLamps(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		tests := `turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500`

		got := Part1(tests)
		want := 1000000 - 1000 - 4

		assertEquality(t, got, want)
	})
}

func assertEquality(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}