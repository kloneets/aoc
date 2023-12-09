package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/kloneets/aoc/2023/helpers"
)

func aoc04(env string) (int, int) {
	log.Println("Day 04:", env, "calculating...")
	if env == PROD {
		data = helpers.FileLines("data/04prod")
	} else {
		data = helpers.FileLines("data/04dev")
	}
	v1 := solveD04P1()

	return v1, solveD04P2()
}

type Card struct {
	Id     int
	Winers []string
	Nums   []string
}

func getCards() []Card {
	cards := make([]Card, 0)
	regex, _ := regexp.Compile("([0-9]+)")
	for _, rawCard := range data {
		m := strings.Split(rawCard, ":")
		id, _ := strconv.Atoi(regex.FindString(m[0]))
		nums := strings.Split(m[1], "|")
		cards = append(cards, Card{
			Id:     id,
			Winers: regex.FindAllString(nums[0], -1),
			Nums:   regex.FindAllString(nums[1], -1),
		})
	}
	return cards
}

func (c *Card) getPoints() int {
	points := 0
	for _, n := range c.Nums {
		for _, w := range c.Winers {
			if n == w {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
				break
			}
		}
	}

	return points
}

func (c *Card) findCopies(copies map[int]int) {
	nextWin := c.Id + 1
	for _, n := range c.Nums {
		for _, w := range c.Winers {
			if n == w {
				for i := 0; i < copies[c.Id]; i++ {
					copies[nextWin]++
				}
				nextWin++
				break
			}
		}
	}
}

func solveD04P1() int {
	cards := getCards()
	res := 0
	for _, card := range cards {
		res += card.getPoints()
	}
	log.Println("Day 04, res 1:", res)
	return res
}

func countCopies(cards []Card) int {
	copies := make(map[int]int)
	res := 0
	for _, card := range cards {
		copies[card.Id]++
		card.findCopies(copies)
	}

	for _, cc := range copies {
		res += cc
	}
	return res
}

func solveD04P2() int {
	cards := getCards()
	res := countCopies(cards)
	log.Println("Day 04, res 2:", res)
	return res
}
