package main

import (
	"github.com/kloneets/aoc/2023/helpers"
	"log"
	"strconv"
	"strings"
)

func aoc02(env string) (int, int) {
	log.Println("Day 02:", env, "calculating...")
	if env == PROD {
		data = helpers.FileLines("data/02prod")
	} else {
		data = helpers.FileLines("data/02dev")
	}
	v1 := solveD02P1()
	// if env == DEV {
	// // 	data = helpers.FileLines("data/02dev2")
	// }
	return v1, solveD02P2()
}

func solveD02P1() int {
	res := 0
gameLine:
	for _, row := range data {
		game := strings.Split(row, ":")
		sets := strings.Split(game[1], ";")
		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				r := strings.Split(strings.Trim(ball, " "), " ")
				bc, _ := strconv.Atoi(r[0])
				switch r[1] {
				case "red":
					if bc > 12 {
						continue gameLine
					}
				case "green":
					if bc > 13 {
						continue gameLine
					}
				case "blue":
					if bc > 14 {
						continue gameLine
					}
				}
			}
		}
		g := strings.Split(strings.Trim(game[0], " "), " ")
		idx, _ := strconv.Atoi(g[1])
		res += idx
	}
	log.Println("Day 02, res 1:", res)
	return res
}

func solveD02P2() int {
	res := 0
	for _, row := range data {
		game := strings.Split(row, ":")
		sets := strings.Split(game[1], ";")
		maxR := 0
		maxG := 0
		maxB := 0
		for _, set := range sets {
			balls := strings.Split(set, ",")
			for _, ball := range balls {
				r := strings.Split(strings.Trim(ball, " "), " ")
				bc, _ := strconv.Atoi(r[0])
				switch r[1] {
				case "red":
					if bc > maxR {
						maxR = bc
					}
				case "green":
					if bc > maxG {
						maxG = bc
					}
				case "blue":
					if bc > maxB {
						maxB = bc
					}
				}
			}
		}
		pow := maxR * maxG * maxB
		res += pow
	}
	log.Println("Day 02, res 2:", res)
	return res
}
