package main

import (
	"fmt"
	"os"
	// "path/filepath"
)

func main() {
	buf, err := os.ReadFile("./2015/day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	ln := NotQuietLisp(string(buf))
	fmt.Println(ln)
}


func NotQuietLisp(s string) int {
	counter := 0
	for _, ch := range s {
		if ch == '(' {
			counter++
		} else {
			counter--
		}
	}
	return counter
}