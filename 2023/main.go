package main

import (
	"log"
	"os"
)

const (
	DEV  = "dev"
	PROD = "prod"
)

var data []string

func main() {
	env := DEV
	if len(os.Args) > 1 && os.Args[1] == PROD {
		env = PROD
	}
	log.Println("AOC 2023 - ", env)
	if len(os.Args) > 2 {
		log.Println("Calculating all days...")
		aoc01(env)
		aoc02(env)
		aoc03(env)
		aoc04(env)
		aoc05(env)
		aoc06(env)
	}
	aoc07(env)
	log.Println("Done!")
}
