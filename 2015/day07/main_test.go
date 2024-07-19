package main

import "testing"

func TestCircuit(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		given := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`
		tests := []string{
			"d",
			"e",
			"f",
			"g",
			"h",
			"i",
			"x",
			"y",
		}
		want := []int{
			72,
			507,
			492,
			114,
			65412,
			65079,
			123,
			456,
		}

		for i, test := range tests {
			t.Run("signal_of_" + test, func(t *testing.T) {
				got := Part1(given, test)
				assertEquality(t, got, want[i])
			})
		}

	})
}

func assertEquality(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}
