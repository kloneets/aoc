package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/kloneets/aoc/2023/helpers"
)

var blocks []string

func aoc05(env string) (int, int) {
	blocks = []string{
		"seed-to-soil map:",
		"soil-to-fertilizer map:",
		"fertilizer-to-water map:",
		"water-to-light map:",
		"light-to-temperature map:",
		"temperature-to-humidity map:",
		"humidity-to-location map:",
	}
	log.Println("Day 05:", env, "calculating...")
	if env == PROD {
		data = helpers.FileLines("data/05prod")
	} else {
		data = helpers.FileLines("data/05dev")
	}
	v1 := solveD05P1()

	return v1, solveD05P2()
}

type Almanac struct {
	Seeds                 []int
	SeedsToSoil           [][]int
	SoilToFertilizer      [][]int
	FertilizerToWater     [][]int
	WaterToLight          [][]int
	LightToTemperature    [][]int
	TemperatureToHumidity [][]int
	HumidityToSoil        [][]int
}

func buildAlamanac() Almanac {
	var seeds []int
	var seedsToSoil [][]int
	var soilToFertilizer [][]int
	var fertilizerToWater [][]int
	var waterToLight [][]int
	var lightToTemperature [][]int
	var temperatureToHumidity [][]int
	var humidityToSoil [][]int
	lastStop := ""
	for _, row := range data {
		if row == "" {
			continue
		}

		if helpers.ContainsString(row, blocks) {
			lastStop = row
			continue
		}

		switch lastStop {
		case "":
			extract := strings.Split(row, ":")
			seeds = numStringToInts(strings.Trim(extract[1], " "))
		case "seed-to-soil map:":
			seedsToSoil = append(seedsToSoil, numStringToInts(row))
		case "soil-to-fertilizer map:":
			soilToFertilizer = append(soilToFertilizer, numStringToInts(row))
		case "fertilizer-to-water map:":
			fertilizerToWater = append(fertilizerToWater, numStringToInts(row))
		case "water-to-light map:":
			waterToLight = append(waterToLight, numStringToInts(row))
		case "light-to-temperature map:":
			lightToTemperature = append(lightToTemperature, numStringToInts(row))
		case "temperature-to-humidity map:":
			temperatureToHumidity = append(temperatureToHumidity, numStringToInts(row))
		case "humidity-to-location map:":
			humidityToSoil = append(humidityToSoil, numStringToInts(row))
		}
	}
	return Almanac{
		Seeds:                 seeds,
		SeedsToSoil:           seedsToSoil,
		SoilToFertilizer:      soilToFertilizer,
		FertilizerToWater:     fertilizerToWater,
		WaterToLight:          waterToLight,
		LightToTemperature:    lightToTemperature,
		TemperatureToHumidity: temperatureToHumidity,
		HumidityToSoil:        humidityToSoil,
	}
}

func numStringToInts(s string) []int {
	sNums := strings.Split(s, " ")
	nums := []int{}
	for _, sNum := range sNums {
		num, _ := strconv.Atoi(sNum)
		nums = append(nums, num)
	}
	return nums
}

func getMap(key int, mapping [][]int) int {
	for _, m := range mapping {
		if key >= m[1] && key <= m[1]+m[2]-1 {
			return m[0] + (key - m[1])
		}
	}
	return key
}

func (a *Almanac) getMin() int {
	res := -1
	for _, s := range a.Seeds {
		mapping := getMap(
			getMap(
				getMap(
					getMap(
						getMap(
							getMap(
								getMap(s, a.SeedsToSoil),
								a.SoilToFertilizer),
							a.FertilizerToWater),
						a.WaterToLight),
					a.LightToTemperature),
				a.TemperatureToHumidity),
			a.HumidityToSoil)
		if res >= mapping || res == -1 {
			res = mapping
		}
	}
	return res
}

func solveD05P1() int {
	almanac := buildAlamanac()
	res := almanac.getMin()
	log.Println("Day 05, res 1:", res)
	return res
}

func (a *Almanac) expandSeeds() {
	var newSeeds []int
	for i := 0; i < len(a.Seeds); i += 2 {
		for j := a.Seeds[i]; j <= a.Seeds[i]+a.Seeds[i+1]-1; j++ {
			newSeeds = append(newSeeds, j)
		}
	}
	a.Seeds = newSeeds
}

// please never do it this way... I am just lazy. It took about 3 min to calculate
func (a *Almanac) getMin2() int {
	res := -1
	for i := 0; i < len(a.Seeds); i += 2 {
		for j := a.Seeds[i]; j <= a.Seeds[i]+a.Seeds[i+1]-1; j++ {
			mapping := getMap(
				getMap(
					getMap(
						getMap(
							getMap(
								getMap(
									getMap(j, a.SeedsToSoil),
									a.SoilToFertilizer),
								a.FertilizerToWater),
							a.WaterToLight),
						a.LightToTemperature),
					a.TemperatureToHumidity),
				a.HumidityToSoil)
			if res >= mapping || res == -1 {
				res = mapping
			}
		}
	}
	return res
}
func solveD05P2() int {
	almanac := buildAlamanac()
	res := almanac.getMin2()
	log.Println("Day 05, res 1:", res)
	return res
}
