package main

import (
	"fmt"
	"maps"
	"math"
	"os"
	"strings"
)

type Graph map[string]map[string]int

func main() {
	inputRaw, _ := os.ReadFile("2015/day13/input.txt")
	input := string(inputRaw)
	val := Part1(makeGraph(input))
	fmt.Printf("val: %v\n", val)
}

func makeGraph(input string) Graph {
	graph := make(Graph)

	for _, line := range strings.Split(input, "\n") {
		var from, to, sign string
		happiness := 0
		fmt.Sscanf(strings.Trim(line, "."), "%v would %v %v happiness units by sitting next to %v",
			&from, &sign, &happiness, &to)

		if sign == "lose" {
			happiness = -happiness
		}

		if graph[from] == nil {
			graph[from] = make(map[string]int)
		}
		if graph[to] == nil {
			graph[to] = make(map[string]int)
		}

		graph[from][to] += happiness
		graph[to][from] += happiness
	}

	return graph
}

func Part1(graph Graph) int {
	checked := make(map[string]bool)
	for f := range graph {
		return Recursion(graph, checked, f, f)
	}
	return 0
}

func Recursion(graph Graph, check map[string]bool, first, person string) int {
	checked := maps.Clone(check)
	checked[person] = true

	highest := math.MinInt
	for p, h := range graph[person] {
		if checked[p] {
			continue
		}
		val := h + Recursion(graph, checked, first, p)
		if val > highest {
			highest = val
		}
	}

	if highest != math.MinInt {
		return highest
	} else {
		return graph[first][person]
	}
}