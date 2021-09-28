// Breadcrumbs game solver.
// For the rules see jjayeon.github.io/breadcrumbs/
package main

import (
	"fmt"
	"time"
    "container/heap"
)

const (
	white = 0 // turn right
	gray  = 1 // turn left
	red   = 2 // obstacle
	green = 3 // goal
)

type State struct {
	color int
	dir   int // N,E,S,W = 0,1,2,3
	x     int
	y     int
}

// Print the board with the ant
func PrintState(s State, board [][]int) {
	fmt.Print("    ")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Print("\n")
	for i, row := range board {
		fmt.Printf("  %d ", i)
		for j, v := range row {
			if i == s.x && j == s.y {
				switch s.dir {
				case 0:
					fmt.Printf("↑ ")
				case 1:
					fmt.Printf("→ ")
				case 2:
					fmt.Printf("↓ ")
				case 3:
					fmt.Printf("← ")
				}
			} else {
				switch v {
				case 0:
					fmt.Printf(". ")
				case 1:
					fmt.Printf("* ")
				case 2:
					fmt.Printf("$ ")
				case 3:
					fmt.Printf("! ")
				}
			}
		}
		fmt.Printf("\n")
	}
}

// Animation of the solution path
func PrintSolution(path []State, board [][]int) {
	flip := []State{}
	score := 0
	for _, s := range path {
		if s.color != board[s.x][s.y] {
			flip = append(flip, s)
			board[s.x][s.y] = s.color
			score++
		}
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print("\033c")
		PrintState(path[i], board)
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Printf("Score: %d\n", score)
	fmt.Print("Flip ")
	for _, s := range flip {
		fmt.Printf("(%d, %d) ", s.x, s.y)
	}
	fmt.Print("\n")
}

// The state that results from flipping the tile
func flip(s State) State {
	if s.color > 1 {
		return s
	}
	s2 := s
	s2.color = 1 - s2.color
	s2.dir = (s.dir + 2) % 4
	return s2
}

func nextState(s State, board [][]int) State {
	size := len(board)
	next := s

	//boundary cases
	if (s.x == 0 && s.dir == 0) ||
		(s.x == size-1 && s.dir == 2) ||
		(s.y == 0 && s.dir == 3) ||
		(s.y == size-1 && s.dir == 1) {
		next.dir = (s.dir + 2) % 4
		return next
	}
	switch s.dir {
	case 0:
		next.x--
	case 1:
		next.y++
	case 2:
		next.x++
	case 3:
		next.y--
	}
	next.color = board[next.x][next.y]
	if next.color == white {
		next.dir = (next.dir + 1) % 4
	} else if next.color == gray {
		next.dir = (next.dir + 3) % 4
	}
	return next
}

// Find a solution to Breadcrumbs board with the minimum number of flips,
// or return with err code -1 if it is impossible.
// Return map with reversed state path
func Solve(board [][]int) (map[State]State, int) {
	//size := len(board)
	pq := make(PriorityQueue, 0)
	score := map[State]int{} // minimum score
	done := map[State]bool{} // processed states
	prev := map[State]State{} // previous state in the optimal path

	start := State{board[0][0], 1, 0, 0}
	heap.Push(&pq, &Item{value: start, priority: 0, index: 0})
	score[start] = 0

	for len(pq) > 0 {
		// find min score state not yet processed
		minItem := heap.Pop(&pq).(*Item)
		minState := minItem.value

		if minState.color == green {
			return prev, minState.dir
		}
		done[minState] = true

		// next state without flipping
		next0 := nextState(minState, board)
		_, foundState := done[next0]
		if !foundState && next0.color != red {
			_, foundScore := score[next0]
			if !foundScore || (score[minState] < score[next0]) {
				score[next0] = score[minState]
				prev[next0] = minState
				heap.Push(&pq, &Item{value: next0, priority: score[next0]})
			}
		}

		// next state with flipping
		next1 := flip(next0)
		_, foundState = done[next1]
		if !foundState && next1.color != red {
			_, foundScore := score[next1]
			if !foundScore || (score[minState]+1 < score[next1]) {
				score[next1] = score[minState] + 1
				prev[next1] = minState
				heap.Push(&pq, &Item{value: next1, priority: score[next1]})
			}
		}
	}
	return prev, -1
}

func main() {
	size := 6
	board := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 1, 2, 1, 1, 1},
		{0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 2},
		{1, 0, 2, 1, 1, 0},
		{0, 0, 0, 0, 0, 3},
	}
	prev, ex := Solve(board)
	if ex == -1 {
		fmt.Println("Unsolvable board!")
	} else {
		// retrace the path
		s := State{}
		if ex == 1 {
			s = State{board[size-1][size-2], 1, size - 1, size - 2}
		} else if ex == 2 {
			s = State{board[size-2][size-1], 2, size - 2, size - 1}
		}
		path := []State{}
		for {
			path = append(path, s)
			last, ok := prev[s]
			if !ok {
				break
			}
			s = last
		}
		PrintSolution(path, board)
	}
}
