package main

import (
	"fmt"
	"strings"

	"koko.lv/aoc"
)

func aoc06() {
	data := aoc.FileLines("inputs/061")
	fmt.Println("Day 6, part 1. Marker after, ", solveAoc06(data, 4, false))
	fmt.Println("Day 6, part 2. Marker after, ", solveAoc06(data, 14, true))
}

func solveAoc06(data []string, letterCount int, second bool) []int {
	res := []int{}
	for _, v := range data {
		i := strings.Split(v, "") // input
		var seq []string
		for idx, letter := range i {
			seq = Aoc06SwapLetters(seq, letter, letterCount)
			if !Aoc06CheckRepeat(seq, letterCount) {
				res = append(res, idx+1)
				break
			}
		}
	}
	return res
}

func Aoc06CheckRepeat(seq []string, letterCount int) bool {
	if len(seq) < letterCount {
		return true
	}
	unique := aoc.UniqueStrings(seq)

	if len(unique) == len(seq) {
		return false
	}

	return true
}

func Aoc06SwapLetters(curState []string, letter string, letterCount int) []string {
	var newState []string
	if len(curState) < letterCount {
		newState = curState
	} else {
		newState = curState[1:]
	}

	newState = append(newState, letter)

	return newState
}
