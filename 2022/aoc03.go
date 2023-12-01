package main

import (
	"fmt"
	"strings"

	"github.com/juliangruber/go-intersect"
	"koko.lv/aoc"
)

func aoc03() {
	data := aoc.FileLines("inputs/031")
	valMap := makeValues()
	fmt.Println("Day 3, part 1. Sum of priorities:", part1(data, valMap))
	fmt.Println("Day 3, part 2. Sum of priorities:", part2(data, valMap))
}

func part2(data []string, valMap map[string]int) int {
	r := 0
	if data[len(data)-1] == "" { //remove empty line if exists
		data = data[:len(data)-1]
	}
	for i := 0; i < len(data); i += 3 {
		is := intersect.Simple(strings.Split(data[i], ""), strings.Split(data[i+1], ""))
		is = intersect.Simple(is, strings.Split(data[i+2], ""))
		if len(is) > 0 {
			r += valMap[fmt.Sprintf("%v", is[0])]
		}
	}
	return r
}

func part1(data []string, valMap map[string]int) int {
	r := 0
	for _, v := range data {
		if v != "" {
			s := strings.Split(v, "")
			h := len(s) / 2
			intersected := intersect.Simple(s[0:h], s[h:])

			if len(intersected) > 0 {
				r += valMap[fmt.Sprintf("%v", intersected[0])]
			}
		}
	}
	return r
}

func makeValues() map[string]int {
	valMap := make(map[string]int)
	val := 1
	for ch := 'a'; ch <= 'z'; ch++ {
		valMap[string(ch)] = val
		val++
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		valMap[string(ch)] = val
		val++
	}
	return valMap
}
