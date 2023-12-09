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
	aoc01(env)
	aoc02(env)
	aoc03(env)
	aoc04(env)
	log.Println("Done!")
}
