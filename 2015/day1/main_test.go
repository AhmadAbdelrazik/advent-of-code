package main

import "testing"

func TestCounter(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		got := NotQuietLisp("))(((((")
		want := 3

		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		got := NotQuietLisp2("()())")
		want := 5

		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})
}
