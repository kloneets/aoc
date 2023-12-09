package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/kloneets/aoc/2023/helpers"
)

func aoc06(env string) (int, int) {
	log.Println("Day 06:", env, "calculating...")
	if env == PROD {
		data = helpers.FileLines("data/06prod")
	} else {
		data = helpers.FileLines("data/06dev")
	}
	v1 := solveD06P1()
	return v1, sovleD06P2()
}

type Race struct {
	T int
	D int
}

func (r *Race) WinOportunities() int {
	res := 0
	for i := 1; i < r.T; i++ {
		if (r.T-i)*i > r.D {
			res++
		}
	}
	return res
}

func solveD06P1() int {
	regex, _ := regexp.Compile("([0-9]+)")
	times := regex.FindAllString(data[0], -1)
	distances := regex.FindAllString(data[1], -1)

	var races []Race

	for i := 0; i < len(times); i++ {
		races = append(races, Race{
			T: helpers.StringToInt(times[i]),
			D: helpers.StringToInt(distances[i]),
		})
	}

	res := 1
	for _, race := range races {
		res *= race.WinOportunities()
	}
	log.Println("Day 06, res 1:", res)
	return res
}

func sovleD06P2() int {
	regex, _ := regexp.Compile("([0-9]+)")
	race := Race{
		T: helpers.StringToInt(regex.FindString(strings.Replace(data[0], " ", "", -1))),
		D: helpers.StringToInt(regex.FindString(strings.Replace(data[1], " ", "", -1))),
	}
	res := race.WinOportunities()
	log.Println("Day 06, res 2:", res)
	return res
}
