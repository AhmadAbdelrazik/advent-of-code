package main

import "testing"

func TestMD5(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		got := Part1("pqrstuv")
		want := 1048970

		assertEquality(t, got, want)
	})
}

func assertEquality(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}