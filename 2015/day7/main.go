package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var timeout int = 0

func main() {
	part := flag.Int("part", 1, "Part 1 or 2")
	flag.Parse()

	buf, _ := os.ReadFile("./2015/day7/input.txt")
	var answer int
	if *part == 1 {
		answer = Part1(string(buf), "a")
	}

	fmt.Println(answer)
}

func Part1(input, target string) int {
	equation := make(map[string]string)

	inputs := strings.Split(input, "\n")

	for _, line := range inputs {
		var eq, out string
		// fmt.Println(line)

		switch {
		case strings.Contains(line, "AND"):
			var in1, in2 string
			fmt.Sscanf(line, "%v AND %v -> %v", &in1, &in2, &out)
			eq = fmt.Sprintf("%v AND %v", in1, in2)
		case strings.Contains(line, "OR"):
			var in1, in2 string
			fmt.Sscanf(line, "%v OR %v -> %v", &in1, &in2, &out)
			eq = fmt.Sprintf("%v OR %v", in1, in2)
		case strings.Contains(line, "NOT"):
			var in string
			fmt.Sscanf(line, "NOT %v -> %v", &in, &out)
			eq = fmt.Sprintf("NOT %v", in)
		case strings.Contains(line, "LSHIFT"):
			var in1 string
			var idx int
			fmt.Sscanf(line, "%v LSHIFT %v -> %v", &in1, &idx, &out)
			eq = fmt.Sprintf("%v LSHIFT %v", in1, idx)
		case strings.Contains(line, "RSHIFT"):
			var in1 string
			var idx int
			fmt.Sscanf(line, "%v RSHIFT %v -> %v", &in1, &idx, &out)
			eq = fmt.Sprintf("%v RSHIFT %v", in1, idx)
		default:
			var in1 string
			fmt.Sscanf(line, "%v -> %v", &in1, &out)
			eq = in1
		}

		equation[out] = eq
	}	

	return int(feedBackward(equation, target))
	return 0
}

func feedBackward(equation map[string]string, target string) uint16 {
	if i, err := strconv.Atoi(target); err == nil {
		return uint16(i)
	}

	timeout++
	if timeout > 5000 {
		log.Fatal("closed")
	}
	switch {
	case strings.Contains(equation[target], "AND"):
		var in1, in2 string
		fmt.Sscanf(equation[target], "%v AND %v", &in1, &in2)
		res1 := feedBackward(equation, in1)
		res2 := feedBackward(equation, in2)
		equation[target] = fmt.Sprint(res1 & res2)
		return res1 & res2
	case strings.Contains(equation[target], "OR"):
		var in1, in2 string
		fmt.Sscanf(equation[target], "%v OR %v", &in1, &in2)
		res1 := feedBackward(equation, in1)
		res2 := feedBackward(equation, in2)
		equation[target] = fmt.Sprint(res1 | res2)
		return res1 | res2
	case strings.Contains(equation[target], "NOT"):
		var in1 string
		fmt.Sscanf(equation[target], "NOT %v", &in1)
		res1 := feedBackward(equation, in1)
		equation[target] = fmt.Sprint(^uint16(0) &^ res1)
		return (^uint16(0) &^ res1)
	case strings.Contains(equation[target], "LSHIFT"):
		var in1 string
		var i int
		fmt.Sscanf(equation[target], "%v LSHIFT %v", &in1, &i)
		res := feedBackward(equation, in1)
		equation[target] = fmt.Sprint(res << i)
		return res << i
	case strings.Contains(equation[target], "RSHIFT"):
		var in1 string
		var i int
		fmt.Sscanf(equation[target], "%v RSHIFT %v", &in1, &i)
		res := feedBackward(equation, in1)
		equation[target] = fmt.Sprint(res >> i)
		return res >> i
	default:
		var in1 string
		fmt.Sscanf(equation[target], "%v", &in1)
		if i, err := strconv.Atoi(in1); err != nil {
			res := feedBackward(equation, in1)
			return res
		} else {
			equation[target] = fmt.Sprint(i)
			return uint16(i)
		}
	}
}
