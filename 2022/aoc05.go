package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"koko.lv/aoc"
)

type Step = int

const (
	Positions Step = iota
	Instruction
)

func aoc05() {
	// data := aoc.FileLines("inputs/050")
	data := aoc.FileLines("inputs/051")

	fmt.Println("Day 5, part 1. Message:", aoc05Solver(data, false))
	fmt.Println("Day 5, part 2. Message:", aoc05Solver(data, true))
}

func aoc05Solver(data []string, newCraftMover bool) string {
	curStep := Positions
	var curPositions []string
	var instructions []string
	for _, v := range data {
		switch v {
		case "":
			curStep = Instruction
		default:
			if curStep == Positions {
				curPositions = append(curPositions, v)
			} else {
				instructions = append(instructions, v)
			}
		}
	}
	return solveAoc05(makeMatrix(curPositions), instructions, newCraftMover)
}

func solveAoc05(matrix [][]string, instructions []string, newCraftMover bool) string {
	r := regexp.MustCompile(`(\d+)`)
	arrangement := make(map[int]string)
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := 0; j < len(matrix[i]); j++ {
			arrangement[j] += strings.Trim(matrix[i][j], " ")
		}
	}

	for _, inst := range instructions {
		row, col, toCol := getInstructions(inst, r)
		column := strings.Split(arrangement[col], "")
		chunk := len(column) - row
		data := column[chunk:]
		if !newCraftMover {
			data = aoc.Reverse(data)
		}
		arrangement[toCol] += strings.Join(data, "")
		arrangement[col] = strings.Join(column[0:chunk], "")
	}

	res := ""
	for i := 0; i < len(arrangement); i++ {
		col := strings.Split(arrangement[i], "")
		res += col[len(col)-1:][0]
	}
	return res
}

func getInstructions(i string, r *regexp.Regexp) (int, int, int) {
	res := r.FindAllString(i, 3)
	what, _ := strconv.Atoi(res[0])
	from, _ := strconv.Atoi(res[1])
	to, _ := strconv.Atoi(res[2])
	return what, from - 1, to - 1
}

func makeMatrix(positions []string) [][]string {
	positions = positions[:len(positions)-1]
	matrix := make([][]string, len(positions))
	for i, line := range positions {
		s := strings.Split(line, "")
		matrix[i] = make([]string, (len(s)+1)/4)
		for j, x := 0, 0; j < len(s); j, x = j+4, x+1 {
			matrix[i][x] = s[j+1]
		}
	}

	return matrix
}
