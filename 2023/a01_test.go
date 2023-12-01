package main

import (
	"testing"
)

func TestSolve01(t *testing.T) {
	t1, t2 := aoc01(DEV)

	if t1 != 142 {
		t.Failed()
	}
	if t2 != 281 {
		t.Failed()
	}
}

func TestTwoNums(t *testing.T) {
	twoNums("6threeseveneightvkqflfp8six3twonebq")
}
