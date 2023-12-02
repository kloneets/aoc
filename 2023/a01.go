package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/kloneets/aoc/2023/helpers"
)

func aoc01(env string) (int, int) {
	log.Println("Day 01:", env, "calculating...")
	if env == PROD {
		data = helpers.FileLines("data/01prod")
	} else {
		data = helpers.FileLines("data/01dev")
	}
	v1 := solveD01P1(env)
	if env == DEV {
		data = helpers.FileLines("data/01dev2")
	}
	return v1, solveD01P2(env)
}

func solveD01P1(env string) int {
	res := 0
	reg, _ := regexp.Compile("([0-9]{1})")
	for r := range data {
		d := reg.FindAllString(data[r], -1)
		num, _ := strconv.Atoi(d[0] + d[len(d)-1])
		res += num
	}

	log.Println("Day 01, res 1:", res)
	return res
}

func solveD01P2(env string) int {
	res := 0

	for r := range data {

		d := twoNums(data[r])
		num, _ := strconv.Atoi(helpers.StringToNumberString(d[0]) + helpers.StringToNumberString(d[len(d)-1]))
		res += num
	}

	log.Println("Day 01, res 2:", res)

	return res
}

func twoNums(hystack string) []string {
	neadle := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var idx []string
	j := 0
	for {
		curIdx := make(map[int]string)
		for _, v := range neadle {
			i := strings.Index(hystack[j:], v)
			if i != -1 {
				curIdx[i] = v
			}

			if len(hystack) <= j {
				break
			}
		}
		sk := 0
		mv := ""
		for k, v := range curIdx {
			if k <= sk {
				mv = v
				sk = k
			}
		}
		if mv != "" {
			idx = append(idx, mv)
		}
		j += 1

		if len(hystack) <= j {
			break
		}
	}
	return idx
}
