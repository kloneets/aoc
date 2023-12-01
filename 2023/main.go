package main

import (
	"log"
	"os"
)

const (
	DEV  = "dev"
	PROD = "prod"
)

func main() {
	env := DEV
	if len(os.Args) > 1 && os.Args[1] == PROD {
		env = PROD
	}
	log.Println("AOC 2023 - ", env)
	aoc01(env)
	log.Println("Done!")
}
