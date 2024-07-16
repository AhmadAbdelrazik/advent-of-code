package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Enter the string: ")
	var a string
	fmt.Scan(&a)
	fmt.Printf("Part1 answer = %v\n", Part1(a))
	fmt.Printf("Part2 answer = %v\n", Part2(a))
}

func Part1(s string) int {
	for i := 0;;i++ {
		str := s + fmt.Sprint(i)
		buf := md5.Sum([]byte(str))
		val := hex.EncodeToString(buf[:])
		if strings.HasPrefix(val, "00000") {
			return i
		}
	}
}

func Part2(s string) int {
	for i := 0;;i++ {
		str := s + fmt.Sprint(i)
		buf := md5.Sum([]byte(str))
		val := hex.EncodeToString(buf[:])
		if strings.HasPrefix(val, "000000") {
			return i
		}
	}
}