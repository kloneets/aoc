package main

import (
	"fmt"
	"strconv"
	"strings"

	"koko.lv/aoc"
)

func aoc04() {
	data := aoc.FileLines("inputs/041")

	fmt.Println("Day 4, part 1. Overlaps all sections:", aoc04Part1(data))
	fmt.Println("Day 4, part 2. Overlaps all sections:", aoc04Part2(data))
}

func aoc04Part2(data []string) int {
	res := 0
	for _, v := range data {
		firstFirst, firstSecond, secondFirst, secondSecond := aoc04Pairs(v)

		if (firstFirst <= secondFirst && firstSecond >= secondSecond) ||
			(secondFirst <= firstFirst && secondSecond >= firstSecond) ||
			(firstFirst <= secondFirst && firstSecond >= secondFirst) ||
			(secondFirst <= firstFirst && secondSecond >= firstFirst) {
			res += 1
		}
	}
	return res
}

func aoc04Part1(data []string) int {
	res := 0
	for _, v := range data {
		firstFirst, firstSecond, secondFirst, secondSecond := aoc04Pairs(v)
		if (firstFirst <= secondFirst && firstSecond >= secondSecond) ||
			(secondFirst <= firstFirst && secondSecond >= firstSecond) {
			res += 1
		}
	}
	return res
}

func aoc04Pairs(v string) (int, int, int, int) {
	pairs := strings.Split(v, ",")
	first := strings.Split(pairs[0], "-")
	second := strings.Split(pairs[1], "-")

	firstFirst, _ := strconv.Atoi(first[0])
	firstSecond, _ := strconv.Atoi(first[1])

	secondFirst, _ := strconv.Atoi(second[0])
	secondSecond, _ := strconv.Atoi(second[1])
	return firstFirst, firstSecond, secondFirst, secondSecond
}
