package main

import (
	"fmt"
	"maps"
	"math"
	"os"
	"strings"
)

type Road struct {
	To   string
	From string
}


func (r Road) flip() Road{
	return Road{To: r.From, From: r.To}
}

func main() {
	buf, _ := os.ReadFile("2015/day9/input.txt")

	ans := Part1(string(buf))
	fmt.Println(ans)
}


func Part1(input string) int {
	cities := make(map[string]bool)
	distances := make(map[Road]int)

	for _, line := range strings.Split(input, "\n") {
		var from, to string
		var distance int
		fmt.Sscanf(line, "%v to %v = %v", &from, &to, &distance)
		cities[from] = false
		cities[to] = false

		road := Road{From: from, To: to}
		distances[road] = distance
		distances[road.flip()] = distance
	}

	ans := math.MaxInt

	for f := range cities {
		val := Rec(f, cities, distances, 0)
		ans = min(ans, val)		
	}

	return ans
}

func Rec(from string, cities map[string]bool, distances map[Road]int, distance int) int {
	dst := make(map[string]bool)
	maps.Copy(dst, cities)

	check := false
	dst[from] = true
	
	ans := math.MaxInt
	for f, s := range dst {
		if !s {
			val := Rec(f, dst, distances, distances[Road{From: from, To: f}])
			ans = min(ans, val)
			check = true
		}
	}

	if check {
		return ans + distance
	} else {
		return distance
	}
}