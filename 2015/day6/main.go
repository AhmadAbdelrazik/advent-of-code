package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	buf,_ := os.ReadFile("./2015/day6/input.txt")
	ans := Part2(string(buf))
	fmt.Println(ans)
}


func Part1(input string) int {
	var grid [1000][1000]int
	for _, line := range strings.Split(input, "\n") {
		var prefix string
		if strings.HasPrefix(line, "turn on") {
			prefix = "turn on"
		} else if strings.HasPrefix(line, "turn off") {
			prefix = "turn off"
		} else if strings.HasPrefix(line, "toggle") {
			prefix = "toggle"
		}

		var x0, y0, x1, y1 int
		cropped, _ := strings.CutPrefix(line, prefix)
		fmt.Sscanf(cropped, "%d,%d through %d,%d", &x0, &y0, &x1, &y1)

		for i := x0; i <= x1; i++ {
			for j := y0; j <= y1; j++ {
				switch prefix {
				case "turn on":
					grid[i][j] = 1
				case "turn off":
					grid[i][j] = 0
				case "toggle":
					if grid[i][j] == 1 {
						grid[i][j] = 0
					} else {
						grid[i][j] = 1
					}
				}
			}
		}
	}

	var ans int 
	for i := 0; i < 1000; i++ {
		for j:= 0; j < 1000; j++ {
			ans += grid[i][j]
		}
	}

	return ans
}

func Part2(input string) int {
	var grid [1000][1000]int
	for _, line := range strings.Split(input, "\n") {
		var prefix string
		if strings.HasPrefix(line, "turn on") {
			prefix = "turn on"
		} else if strings.HasPrefix(line, "turn off") {
			prefix = "turn off"
		} else if strings.HasPrefix(line, "toggle") {
			prefix = "toggle"
		}

		var x0, y0, x1, y1 int
		cropped, _ := strings.CutPrefix(line, prefix)
		fmt.Sscanf(cropped, "%d,%d through %d,%d", &x0, &y0, &x1, &y1)

		for i := x0; i <= x1; i++ {
			for j := y0; j <= y1; j++ {
				switch prefix {
				case "turn on":
					grid[i][j] += 1
				case "turn off":
					if grid[i][j] > 0 {
						grid[i][j] -= 1
					}
				case "toggle":
					grid[i][j] += 2
				}
			}
		}
	}

	var ans int 
	for i := 0; i < 1000; i++ {
		for j:= 0; j < 1000; j++ {
			ans += grid[i][j]
		}
	}

	return ans
}
