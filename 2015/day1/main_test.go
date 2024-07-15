package main

import "testing"


func TestCounter(t *testing.T) {
		got := NotQuietLisp("))(((((")
		want := 3

		if got != want {
			t.Errorf("got '%v' want '%v'", got, want)
		}
}