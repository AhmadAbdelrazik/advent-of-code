package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rawInput, _ := os.ReadFile("2015/day12/input.txt")
	input := string(rawInput)

	ans := JsonParser(input, 2)
	fmt.Printf("ans: %v\n", ans)
}

func JsonParser(input string, part int) int {
	var result interface{}

	if err := json.Unmarshal([]byte(input), &result); err != nil {
		panic(err)
	}

	var answer int
	if part == 1 {
		answer = Part1Recursion(result)
	} else if part == 2 {
		answer = Part2Recursion(result)
		
	}
	return answer
}

func Part1Recursion(val interface{}) int {
	ans := 0

	switch i := val.(type) {
	case string:
		if v, err := strconv.Atoi(i); err == nil {
			return v
		}
	case float32:
	case float64:
		return int(i)
	case map[string]interface{}:
		sum := 0
		for _, s := range i {
			sum += Part1Recursion(s)
		}
		return sum
	case []interface{}:
		sum := 0
		for _, j := range i {
			sum += Part1Recursion(j)
		}
		return sum
	default:
		fmt.Println("uncaught type", i)
	}
	return ans
}

func Part2Recursion(val interface{}) int {
	ans := 0

	switch i := val.(type) {
	case string:
		if v, err := strconv.Atoi(i); err == nil {
			return v
		}
	case float32:
	case float64:
		return int(i)
	case map[string]interface{}:
		sum := 0
		for f, s := range i {
			if f == "red" || s == "red" {
				return 0
			}
			sum += Part2Recursion(s)
		}
		return sum
	case []interface{}:
		sum := 0
		for _, j := range i {
			sum += Part2Recursion(j)
		}
		return sum
	default:
		fmt.Println("uncaught type", i)
	}
	return ans
}