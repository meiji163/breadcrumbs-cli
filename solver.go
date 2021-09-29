// Breadcrumbs game solver.
package main

import "container/heap"

// Find a solution to Breadcrumbs board with the minimum number of flips,
// or return with err code -1 if it is impossible.
// Return map with reversed state path
func Solve(board [][]int) (map[State]State, int) {
	//size := len(board)
	pq := make(PriorityQueue, 0)


    /* Run Djikstra on the state graph
    
           currentState
       cost=0 /       \ cost=1
             |         |
        notFlipped    Flipped
    */

	score := map[State]int{} // minimum score
	done := map[State]bool{} // processed states
	prev := map[State]State{} // previous state in the optimal path

    // prev[ prev[s] ] -> prev[s] -> s

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

func RetracePath(board [][]int, prev map[State]State, ex int) []State {
    // retrace the path
    s := State{}
    size := len(board)
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
    return path
}
