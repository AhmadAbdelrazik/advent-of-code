package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("Part1(\"cqjxjnds\"): %v\n", Part1("cqjxjnds"))
	fmt.Printf("Part1(\"cqjxxyzz\"): %v\n", Part1("cqjxxyzz"))
}

func Incrementer(in string) string {
	inc := true
	input := []byte(in)

	for i := 0; i < len(in); i++ {
		if in[i] == 'i' || in[i] == 'o' || in[i] == 'l' {
			val := input[:i]
			val = append(val, in[i]+1)
			val = append(val, bytes.Repeat([]byte("a"), len(in)-i-1)...)
			return string(val)
		}
	}

	for i := len(in) - 1; i >= 0; i-- {
		if !inc {
			break
		}
		if input[i] == 'z' {
			input[i] = 'a'
			if i == 0 {
				input = append([]byte{'a'}, input...)
			}
		} else {
			inc = false
			input[i] += 1
		}
	}

	return string(input)
}

func valid(in string) bool {
	rule1 := false // abc, bcd, xyz
	rule3 := false // aa, bb, zz * 2

	for i := 0; i < len(in); i++ {
		if i < len(in)-2 && in[i] == in[i+1]-1 && in[i] == in[i+2]-2 {
			rule1 = true
		}
	}

	count := 0
	for i := 0; i < len(in)-1; i++ {
		if in[i] == in[i+1] {
			count++
			i++
		}

		if count == 2 {
			rule3 = true
			break
		}
	}

	return rule1 && rule3
}

func Part1(in string) string {
	for {
		in = Incrementer(in)
		if valid(in) {
			break
		}
	}
	return in
}
