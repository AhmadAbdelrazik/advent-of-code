package main

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	t.Run("Part1", func(t *testing.T) {
		tests := []string{
			`[1,2,3]`,
			`{"a":2,"b":4}`,
			`[[[3]]]`,
			`{"a":{"b":4},"c":-1}`,
			`{}`,
			`[]`,
		}
		want := []int{6, 6, 3, 3, 0, 0}

		for i := range tests {
			t.Run(fmt.Sprintf("test %v", tests[i]), func(t *testing.T) {
				got := JsonParser(tests[i], 1)

				if got != want[i] {
					t.Errorf("got '%v' want '%v'", got, want[i])
				}
			})
		}
	})

	t.Run("Part2", func(t *testing.T) {
		tests := []string{
			`[1,2,3]`,
			`[1,{"c":"red","b":2},3]`,
			`{"d":"red","e":[1,2,3,4],"f":5}`,
			`[1,"red",5]`,
		}
		want := []int{6, 4, 0, 6}
		
		for i := range tests {
			t.Run(fmt.Sprintf("test %v", tests[i]), func(t *testing.T) {
				got := JsonParser(tests[i], 2)
				
				if got != want[i] {
					t.Errorf("got '%v' want '%v'", got, want[i])
				}
			})
		}
	})
}
