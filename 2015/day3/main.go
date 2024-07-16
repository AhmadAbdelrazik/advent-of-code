package main

import (
	"fmt"
	"log"
	"os"
)

type point struct {
	X int
	Y int
}

func main() {
	buf, err := os.ReadFile("2015/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ans := VisitedHouses(string(buf))
	fmt.Println(ans)
}

func VisitedHouses(s string) int {
	
	p := point{X: 0, Y: 0}
	grid := make(map[point]int)
	
	grid[p]++
	for _, ch := range s {
		if ch == '^' {
			p.Y++
		} else if ch == 'v' {
			p.Y--
		} else if ch == '>' {
			p.X++
		} else if ch == '<' {
			p.X--
		}
		grid[p]++
	}

	return len(grid)
} 