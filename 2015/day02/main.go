package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	buf, err := os.ReadFile("2015/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	equations := strings.Split(string(buf), "\n")
	ans := 0
	for _, equation := range equations {
		ans += WrappingPaper2(equation)
	}
	fmt.Println(ans)
}

func WrappingPaper2(eq string) int {
	numberStrings := strings.Split(eq, "x")
	numbers := make([]int, 3)

	for i, n := range numberStrings {
		number, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = number
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	answer := numbers[0] * numbers[1] * numbers[2]
	answer += (numbers[0] + numbers[1]) * 2 

	return answer
}

func WrappingPaper(eq string) int {
	numberStrings := strings.Split(eq, "x")
	numbers := make([]int, 3)

	for i, n := range numberStrings {
		number, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numbers[i] = number
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	answer := numbers[0] * numbers[1] * 3
	answer += numbers[0] * numbers[2] * 2
	answer += numbers[1] * numbers[2] * 2

	return answer
}

