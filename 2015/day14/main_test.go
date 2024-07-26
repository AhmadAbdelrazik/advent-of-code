package main

import "testing"

func TestReindeerRace(t *testing.T) {
	given := []Reindeer{
		{Name: "Comet", Speed: 14, Stamina: 10, Rest: 127},
		{Name: "Dancer", Speed: 16, Stamina: 11, Rest: 162},
	}
	seconds := 1000
	t.Run("Part1", func(t *testing.T) {
		got := Part1(given, seconds)
		want := 1120
		
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
	})

	t.Run("Part2", func(t *testing.T) {
		got := Part2(given, seconds)
		want := 689
		
		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}

	})
}
