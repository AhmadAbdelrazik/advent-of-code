package main

import (
	"fmt"
	"time"
)

func main() {
	input := "1321131112"

	fmt.Println(Part1(input, 50))
}

func Part1(in string, iterations int) int {

	out := ""
	start := time.Now()
	for j := 0; j < iterations; j++ {
		iterationStart := time.Now()
		out = ""
		count := 0
		for i := 0; i < len(in); i++ {
			count++
			if i == len(in)-1 || in[i] != in[i+1] {
				out += fmt.Sprint(count)
				out += fmt.Sprintf("%v", in[i]-48)
				count = 0
			}
		}
		fmt.Println("iteration number ", j+1)
		in = out
		fmt.Println("Length = ", len(out))
		fmt.Println("Time passed", time.Since(iterationStart))
	}
	fmt.Println("Total time = ", time.Since(start))
	return len(out)
}
