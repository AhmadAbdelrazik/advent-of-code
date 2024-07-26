package main

import (
	"fmt"
	"os"
	"strings"
)

type Reindeer struct {
	Name    string
	Speed   int
	Stamina int
	Rest    int
}

func main() {
	inputRaw, _ := os.ReadFile("2015/day14/input.txt")
	input := string(inputRaw)
	reindeers := toReindeer(input)

	part1 := Part1(reindeers, 2503)
	part2 := Part2(reindeers, 2503)

	fmt.Printf("part1: %v\n", part1)
	fmt.Printf("part2: %v\n", part2)
}

func toReindeer(input string) []Reindeer {
	var Reindeers []Reindeer

	for _, line := range strings.Split(input, "\n") {
		var speed, stamina, rest int
		var name string

		fmt.Sscanf(line, "%v can fly %v km/s for %v seconds, but then must rest for %v seconds.", &name, &speed, &stamina, &rest)

		Reindeers = append(Reindeers, Reindeer{Name: name, Speed: speed, Rest: rest, Stamina: stamina})
	}

	fmt.Printf("Reindeers: %v\n", Reindeers)

	return Reindeers
}

func Part1(Reindeers []Reindeer, seconds int) int {
	highest := 0

	for _, reindeer := range Reindeers {
		dis := DistanceAtSecond(reindeer, seconds)
		if dis > highest {
			highest = dis
		}
	}

	return highest
}

func DistanceAtSecond(reindeer Reindeer, seconds int) int {
	lapTime := reindeer.Rest + reindeer.Stamina
	lapDistance := reindeer.Speed * reindeer.Stamina

	fullLaps := seconds / lapTime
	remainingSeconds := seconds % lapTime

	totalDistance := lapDistance * fullLaps

	if remainingSeconds > reindeer.Stamina {
		totalDistance += lapDistance
	} else {
		totalDistance += remainingSeconds * reindeer.Speed
	}

	return totalDistance
}

func Part2(Reindeers []Reindeer, seconds int) int {
	ScoreBoard := make(map[Reindeer]int)

	for i := 1; i <= seconds; i++ {
		var topReindeers []Reindeer
		highest := 0
		for _, r := range Reindeers {
			val := DistanceAtSecond(r, i)
			if val > highest {
				highest = val
				topReindeers = append([]Reindeer{}, r)
			} else if val == highest {
				topReindeers = append(topReindeers, r)
			}
		}

		for _, r := range topReindeers {
			ScoreBoard[r]++
		}
	}

	winner := 0

	for _, score := range ScoreBoard {
		if score > winner {
			winner = score
		}
	}
	return winner
}
