package main

import (
	"fmt"
	"strings"

	"koko.lv/aoc"
)

type Output int

const (
	Lose Output = iota
	Tie
	Win
)

func aoc020() {
	data := aoc.FileLines("inputs/021")

	fmt.Println("Day 2, part 1. Got points:", partOne(data))
	fmt.Println("Day 2, part 2. Got points:", partTwo(data))
}

func partTwo(data []string) int {
	res := 0
	m := make(map[string]Output)
	m["X"] = Lose
	m["Y"] = Tie
	m["Z"] = Win

	m2 := make(map[string]int)
	m2["A"] = 1
	m2["B"] = 2
	m2["C"] = 3

	for _, v := range data {
		if v != "" {
			round := strings.Split(v, " ")
			res += gameResult2(round[0], m[round[1]], m2)
		}
	}

	return res
}

func gameResult2(opponentStep string, o Output, m map[string]int) int {
	r := 0
	switch o {
	case Lose:
		switch m[opponentStep] {
		case 1:
			r += 3
		case 2:
			r += 1
		case 3:
			r += 2
		}
	case Tie:
		r += m[opponentStep] + 3
	case Win:
		switch m[opponentStep] {
		case 1:
			r += 2
		case 2:
			r += 3
		case 3:
			r += 1
		}
		r += 6
	}
	return r
}

func partOne(data []string) int {
	m := make(map[string]int)
	m["A"] = 1
	m["B"] = 2
	m["C"] = 3
	m["X"] = 1
	m["Y"] = 2
	m["Z"] = 3

	res := 0
	for _, v := range data {
		if v != "" {
			round := strings.Split(v, " ")
			res += m[round[1]]
			switch gameResult(round[0], round[1], m) {
			case Win:
				res += 6
			case Tie:
				res += 3
			}
		}
	}
	return res
}

func gameResult(opponentStep string, myStep string, m map[string]int) Output {
	if m[opponentStep] == m[myStep] {
		return Tie
	}
	switch m[opponentStep] {
	case 1:
		if m[myStep] == 2 {
			return Win
		}
	case 2:
		if m[myStep] == 3 {
			return Win
		}
	case 3:
		if m[myStep] == 1 {
			return Win
		}
	}
	return Lose
}
