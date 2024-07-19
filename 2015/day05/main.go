package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	buf, err := os.ReadFile("./2015/day5/input.txt")
	if err != nil {
		panic(err)
	}
	tests := strings.Split(string(buf), "\n")
	ans := 0
	for _, test := range tests {
		if Part2(test) {
			ans++
		}
	}
	fmt.Println(ans)
}

func Part1(s string) bool {
	freq := make(map[rune]int)

	duplicateCheck := false

	for i, ch := range s {
		freq[ch]++
		if i != 0 {
			if s[i] == s[i-1] {
				duplicateCheck = true
			}
		}
	}

	if !duplicateCheck {
		return false
	}

	if strings.Contains(s, "ab") || strings.Contains(s, "bc") || strings.Contains(s, "cd") || strings.Contains(s, "xy") {
		return false
	}

	vowels := freq['a'] + freq['e'] + freq['i'] + freq['o'] + freq['u']

	return vowels >= 3
}

func Part2(s string) bool {
	rule1 := false
	rule2 := false


	for i := 0; i < len(s)-2; i++ {
		goal := s[i : i+2]
		for j := i + 2; j < len(s)-1; j++ {
			if s[j:j+2] == goal {
				rule1 = true
			}
		}
	}

	for i := range s {
		if i > 1 {
			if s[i] == s[i-2] {
				rule2 = true
			}
		}
	}
	return rule1 && rule2
}
