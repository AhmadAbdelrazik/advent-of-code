package main

import "testing"

func TestMatchsticks(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		tests := `""
"abc"
"aaa\"aaa"
"\x27"`
		got := Part1(tests)
		want := 12

		assertEquality(t, got, want)
	})
	t.Run("Part2", func(t *testing.T) {
		tests := `""
"abc"
"aaa\"aaa"
"\x27"`
		got := Part2(tests)
		want := 19

		assertEquality(t, got, want)
	})
}

func assertEquality(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got '%v' want '%v'", got, want)
	}
}