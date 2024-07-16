package main

import (
	"fmt"
	"log"
	"os"
)

type Point struct {
	X int
	Y int
}

func main() {
	buf, err := os.ReadFile("2015/day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	ans := VisitedHouses2(string(buf))
	fmt.Println(ans)
}

func VisitedHouses(s string) int {

	p := Point{X: 0, Y: 0}
	grid := make(map[Point]int)

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

func VisitedHouses2(s string) int {
	santa := Point{X: 0, Y: 0}
	robo := Point{X: 0, Y: 0}
	grid := make(map[Point]int)

	grid[santa]++
	grid[robo]++

	for i, ch := range s {
		var pp *Point

		if i%2 == 0 {
			pp = &santa
		} else {
			pp = &robo
		}

		func(p *Point) {

			if ch == '^' {
				p.Y++
			} else if ch == 'v' {
				p.Y--
			} else if ch == '>' {
				p.X++
			} else if ch == '<' {
				p.X--
			}
			grid[*p]++
		}(pp)
	}

	return len(grid)
}
