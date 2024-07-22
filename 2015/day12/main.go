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

	ans := Part1(input)
	fmt.Printf("ans: %v\n", ans)
}

func Part1(input string) int {
	var result interface{}

	if err := json.Unmarshal([]byte(input), &result); err != nil {
		panic(err)
	}

	val := iterator(result)
	return val
}

func iterator(val interface{}) int {
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
			sum += iterator(s)
		}
		return sum
	case []interface{}:
		sum := 0
		for _, j := range i {
			sum += iterator(j)
		}
		return sum
	default:
		fmt.Println("uncaught type", i)
	}
	return ans
}
