package main

import "testing"

func TestWrappingPaper(t *testing.T) {
	t.Run("Part 1", func(t *testing.T) {
		got := WrappingPaper("2x3x4")
		want := 58
	
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})

	t.Run("Part 2", func(t *testing.T) {
		got := WrappingPaper2("2x3x4")
		want := 34
	
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}

	})
}