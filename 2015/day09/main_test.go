package main

import "testing"

func TestShortestPath(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		given := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`
		got := Part1(given)
		want := 605

		assertEquality(t, got, want)
	})
	t.Run("Part2", func(t *testing.T) {
		given := `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`
		got := Part2(given)
		want := 982

		assertEquality(t, got, want)
	})
}

func assertEquality(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want'%v'", got, want)
	}
}