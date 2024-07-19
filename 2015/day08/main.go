package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	part := flag.Int("part", 1, "Part 1 or 2")
	flag.Parse()

	buf, _ := os.ReadFile("./2015/day8/input.txt")
	var answer int
	if *part == 1 {
		answer = Part1(string(buf))
	} else {
		answer = Part2(string(buf))
	}

	fmt.Println(answer)
}

func Part1(input string) int {
	ans := 0
	for _, line := range strings.Split(input, "\n") {
		ans += len(line)
		lineDiff := len(line) - 2

		for i := 0; i < len(line); i++ {
			if line[i] == '\\' {
				if line[i+1] == '\\' {
					lineDiff -= 1
					i++
				} else if line[i+1] == 'x' {
					lineDiff -= 3
				} else {
					lineDiff -= 1
				}
			}
		}

		ans -= lineDiff
	}
	return ans
}

func Part2(input string) int {
	ans := 0
	for _, line := range strings.Split(input, "\n") {
		ans += len(line) + 2
		ans += strings.Count(line, `"`)
		ans += strings.Count(line, `\`)
		ans -= len(line)
	}
	return ans
}