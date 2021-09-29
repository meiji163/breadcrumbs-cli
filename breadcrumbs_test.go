package main

import "testing"

func TestSolve(t *testing.T) {
	board1 := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 2, 1, 1, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 2},
		{1, 0, 2, 1, 1, 0},
		{0, 0, 0, 0, 0, 3},
	}
	_, ex := Solve(board1)
	if ex == -1 {
		t.Log("Exit code ", ex, ", but board is solvable")
		t.Fail()
	}
}

func TestUnsolvable(t *testing.T) {
	board2 := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 2, 1, 1, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 2},
		{1, 0, 2, 1, 2, 0},
		{0, 0, 0, 0, 0, 3},
	}

	_, ex := Solve(board2)
	if ex != -1 {
		t.Log("Exit bode ", ex, ", but board is unsolvabe")
		t.Fail()
	}
}
