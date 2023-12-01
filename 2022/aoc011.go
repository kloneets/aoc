package main

import (
	"fmt"
	"strconv"

	"koko.lv/aoc"
)

func aoc011() {
	calorieMap := aoc.FileLines("inputs/011")

	totalCalories := 0
	totalMax := 0
	secondMax := 0
	thirdMax := 0

	// lazy hard code.. It can be done universally, but who cares?
	for _, v := range calorieMap {
		if v == "" {
			if totalMax < totalCalories {
				secondMax = totalMax
				totalMax = totalCalories
			} else if secondMax < totalCalories {
				thirdMax = secondMax
				secondMax = totalCalories
			} else if thirdMax < totalCalories {
				thirdMax = totalCalories
			}
			totalCalories = 0
		} else {
			n, _ := strconv.Atoi(v)
			totalCalories += n
		}
	}
	if totalMax < totalCalories {
		secondMax = totalMax
		totalMax = totalCalories
	} else if secondMax < totalCalories {
		thirdMax = secondMax
		secondMax = totalCalories
	} else if thirdMax < totalCalories {
		thirdMax = totalCalories
	}
	fmt.Println("Day 1, part 2. Three maxes:", totalMax+secondMax+thirdMax)

}
