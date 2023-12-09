package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/kloneets/aoc/2023/helpers"
)

func aoc03(env string) (int, int) {
	log.Println("Day 03:", env, "calculating...")
	if env == PROD {
		data = helpers.FileLines("data/03prod")
	} else {
		data = helpers.FileLines("data/03dev")
	}
	v1 := solveD03P1()
	return v1, solveD03P2()
}

type Pos struct {
	x int
	y int
}

type Engine struct {
	Schema    [][]string
	SymbolPos []Pos
	Numbers   *[]int
	XLen      int
	YLen      int
	Picked    map[string][]Pos
}

func (e *Engine) getNumbers() {
	for _, sp := range e.SymbolPos {
		adj := getAdjacent(e.Schema, sp.x, sp.y)
		for _, pos := range adj {
			key := fmt.Sprintf("%03d", pos.x)
			if e.Schema[pos.x][pos.y] != "." && !e.picked(key, pos) {
				_, err := strconv.Atoi(e.Schema[pos.x][pos.y])
				if err == nil {
					e.setPicked(key, pos)
					numberString := e.buildNumber(e.Schema[pos.x][pos.y], pos.x, pos.y, Direction(Both), key)
					number, _ := strconv.Atoi(numberString)
					*e.Numbers = append(*e.Numbers, number)
				}
			}
		}
	}
}

func (e *Engine) setPicked(key string, pos Pos) {
	e.Picked[key] = append(e.Picked[key], pos)
}

func (e *Engine) picked(key string, pos Pos) bool {
	if e.Picked[key] != nil {
		for _, v := range e.Picked[key] {
			if v.x == pos.x && v.y == pos.y {
				return true
			}
		}
	}
	return false
}

type Direction int

const (
	Both  int = 0
	Left      = 1
	Right     = 2
)

func (e *Engine) buildNumber(num string, x int, y int, dir Direction, key string) string {
	if dir == Direction(Both) || dir == Direction(Left) {
		adjLeft := getAdjacentVertical(e.Schema, x, y, e.XLen, e.YLen, Left)
		if len(adjLeft) > 0 && !e.picked(key, adjLeft[0]) {
			s := e.Schema[adjLeft[0].x][adjLeft[0].y]
			if s != "." {
				_, err := strconv.Atoi(s)
				if err == nil {
					e.setPicked(key, adjLeft[0])
					num = s + num
					num = e.buildNumber(num, adjLeft[0].x, adjLeft[0].y, Direction(Left), key)
				}
			}
		}
	}

	if dir == Direction(Both) || dir == Direction(Right) {
		adjRight := getAdjacentVertical(e.Schema, x, y, e.XLen, e.YLen, Right)
		if len(adjRight) > 0 && !e.picked(key, adjRight[0]) {
			s := e.Schema[adjRight[0].x][adjRight[0].y]
			if s != "." {
				_, err := strconv.Atoi(s)
				if err == nil {
					e.setPicked(key, adjRight[0])
					num = num + s
					num = e.buildNumber(num, adjRight[0].x, adjRight[0].y, Direction(Right), key)
				}
			}
		}
	}
	return num
}

func (e *Engine) sumNumbers() int {
	res := 0
	if e.Numbers != nil {
		for _, v := range *e.Numbers {
			res += v
		}
	}
	return res
}

func (e *Engine) calculateGears() int {
	allGears := 0
	for _, sp := range e.SymbolPos {
		adj := getAdjacent(e.Schema, sp.x, sp.y)
		curNum := make([]int, 0)
		for _, pos := range adj {
			key := fmt.Sprintf("%03d", pos.x)
			if e.Schema[pos.x][pos.y] != "." && !e.picked(key, pos) {
				_, err := strconv.Atoi(e.Schema[pos.x][pos.y])
				if err == nil {
					e.setPicked(key, pos)
					numberString := e.buildNumber(e.Schema[pos.x][pos.y], pos.x, pos.y, Direction(Both), key)
					number, _ := strconv.Atoi(numberString)
					curNum = append(curNum, number)
				}
			}
		}
		if len(curNum) > 1 {
			gearRatio := 1
			for _, nv := range curNum {
				gearRatio *= nv
			}
			allGears += gearRatio
		}
	}
	return allGears
}

func parseData(stars bool) Engine {
	schema := make([][]string, 0)
	var symbols []Pos
	for i, v := range data {
		schema = append(schema, strings.Split(v, ""))
		for j, jv := range schema[i] {
			if jv != "." {
				_, err := strconv.Atoi(jv)
				if err != nil && (!stars || (stars && jv == "*")) {
					symbols = append(symbols, Pos{i, j})
				}
			}
		}
	}
	return Engine{
		Schema:    schema,
		SymbolPos: symbols,
		Numbers:   &[]int{},
		XLen:      len(schema),
		YLen:      len(schema[0]),
		Picked:    make(map[string][]Pos, 0),
	}

}

func solveD03P1() int {
	res := 0
	engine := parseData(false)

	engine.getNumbers()
	res = engine.sumNumbers()
	log.Println("Day 03, res 1:", res)
	return res
}

func solveD03P2() int {
	res := 0
	engine := parseData(true)

	res = engine.calculateGears()
	log.Println("Day 03, res 2:", res)
	return res
}

func getAdjacent(schema [][]string, i int, j int) []Pos {
	n := len(schema)
	m := len(schema[0])

	var res []Pos
	if isValidPos(i-1, j-1, n, m) {
		res = append(res, Pos{i - 1, j - 1})
	}
	if isValidPos(i-1, j, n, m) {
		res = append(res, Pos{i - 1, j})
	}
	if isValidPos(i-1, j+1, n, m) {
		res = append(res, Pos{i - 1, j + 1})
	}
	if isValidPos(i+1, j-1, n, m) {
		res = append(res, Pos{i + 1, j - 1})
	}
	if isValidPos(i+1, j, n, m) {
		res = append(res, Pos{i + 1, j})
	}
	if isValidPos(i+1, j+1, n, m) {
		res = append(res, Pos{i + 1, j + 1})
	}
	res = append(res, getAdjacentVertical(schema, i, j, n, m, Direction(Both))...)
	return res
}

func isValidPos(i int, j int, n int, m int) bool {
	return i >= 0 && i < n && j >= 0 && j < m
}

func getAdjacentVertical(schema [][]string, i int, j int, n int, m int, direction Direction) []Pos {
	var res []Pos

	if isValidPos(i, j-1, n, m) && (direction == Direction(Both) || direction == Direction(Left)) {
		res = append(res, Pos{i, j - 1})
	}
	if isValidPos(i, j+1, n, m) && (direction == Direction(Both) || direction == Direction(Right)) {
		res = append(res, Pos{i, j + 1})
	}
	return res
}
