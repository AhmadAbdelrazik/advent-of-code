package main

import (
	"fmt"
)

func main() {
	input := "1321131112"
	
	var length int 
	for i := 0; i < 40; i++ {
		input, length = Part1(input)
		fmt.Println("iteration number ", i + 1)
	}
	fmt.Println(length)
}

func Part1(in string) (string, int) {

	freq := make(map[byte]int)

	out := ""

	for i := 0; i < len(in); i++ {
		freq[in[i]]++
		if i == len(in)-1 {
			out += fmt.Sprint(freq[in[i]])
			out += fmt.Sprintf("%v", in[i]-48)
			freq[in[i]] = 0
		} else if in[i] != in[i+1] {
			out += fmt.Sprint(freq[in[i]])
			out += fmt.Sprintf("%d", in[i]-48)
			freq[in[i]] = 0
		}
	}

	return out, len(out)
}
