package main

import (
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/kloneets/aoc/2023/helpers"
)

func aoc07(env string) (int, int) {
	log.Println("Day 07:", env, "calculating...")
	if env == PROD {
		data = helpers.FileLines("data/07prod")
	} else {
		data = helpers.FileLines("data/07dev")
	}
	v1 := solveD07P1()
	return v1, solveD07P2()
}

func solveD07P1() int {
	rankCards := RankCards(data, false)
	regex, _ := regexp.Compile("([0-9]+)$")
	res := 0
	for i := 0; i < len(rankCards); i++ {
		bid := helpers.StringToInt(regex.FindString(rankCards[i]))
		res += bid * (i + 1)
	}
	log.Println("Day 07, res 1:", res)
	return res
}

func solveD07P2() int {
	rankCards := RankCards(data, true)
	regex, _ := regexp.Compile("([0-9]+)$")
	res := 0
	for i := 0; i < len(rankCards); i++ {
		bid := helpers.StringToInt(regex.FindString(rankCards[i]))
		res += bid * (i + 1)
	}
	log.Println("Day 07, res 2:", res)
	return res
}

func RankCards(array []string, alterJ bool) []string {
	sort.Slice(array, func(i, j int) bool {
		// Get the first part of the element
		firstPartI := array[i][:5]
		firstPartJ := array[j][:5]

		sortPriority := "AKQJT98765432"
		if alterJ {
			sortPriority = "AKQJT98765432J"
		}
		firstChars, firstString := countCharacters(firstPartI, alterJ, sortPriority)
		secondChars, secondString := countCharacters(firstPartJ, alterJ, sortPriority)
		// log.Println(firstPartI, firstString)
		maxFirst := maxCharCount(firstChars)
		maxSecond := maxCharCount(secondChars)
		if maxFirst < maxSecond {
			return true
		} else if maxFirst > maxSecond {
			return false
		} else if maxFirst == maxSecond {
			twoPairsFirst := hasTwoPairs(firstChars)
			twoPairsSecond := hasTwoPairs(secondChars)
			// possible full hause
			// check of one of them has two pairs
			if twoPairsFirst && !twoPairsSecond {
				return false
			} else if !twoPairsFirst && twoPairsSecond {
				return true
			}
		}

		for i := 0; i < 5; i++ {
			firstIdx := strings.Index(sortPriority, string(firstString[i]))
			secondIdx := strings.Index(sortPriority, string(secondString[i]))
			if firstIdx > secondIdx {
				return true
			} else if firstIdx < secondIdx {
				return false
			}
		}
		return false
	})

	return array
}

func hasTwoPairs(chars map[string]int) bool {
	pairs := 0
	for _, v := range chars {
		if v >= 2 {
			pairs++
		}
	}
	return pairs > 1
}

func maxCharCount(chars map[string]int) int {
	max := 0
	for _, v := range chars {
		if v > max {
			max = v
		}
	}
	return max
}

func countCharacters(s string, alterJ bool, sortPriority string) (map[string]int, string) {
	res := make(map[string]int)
	for _, c := range s {
		res[string(c)]++
	}

	return res, s
}
