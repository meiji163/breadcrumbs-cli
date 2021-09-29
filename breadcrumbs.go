// For the rules see jjayeon.github.io/breadcrumbs/
package main

import (
	"fmt"
	"time"
    "os"
    "math/rand"
)

const (
	white = 0 // turn right
	gray  = 1 // turn left
	red   = 2 // obstacle
	green = 3 // goal
)

// Game state
type State struct {
	color int
	dir   int // N,E,S,W = 0,1,2,3
	x     int
	y     int
}

// Print the board with the ant
func PrintState(s State, board [][]int) {
	fmt.Print("            ")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Print("\n")

    colorRed := string("\033[31m")
    colorGreen := string("\033[32m")
    colorReset := string("\033[0m")

	for i, row := range board {
		fmt.Printf("\t  %d ", i)
		for j, v := range row {
			if i == s.x && j == s.y {
				switch s.dir {
				case 0:
					fmt.Print(colorGreen,"△ ")
				case 1:
					fmt.Print(colorGreen,"▻ ")
				case 2:
					fmt.Print(colorGreen,"▽ ")
				case 3:
					fmt.Print(colorGreen,"◅ ")
				}
                fmt.Print(colorReset)
			} else {
				switch v {
				case 0:
					fmt.Printf("◻ ")
				case 1:
					fmt.Printf("◼ ")
				case 2:
					fmt.Print(colorRed,"▩ ", colorReset)
				case 3:
                    fmt.Print(colorGreen,"◈ ",colorReset)
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
        fmt.Print("\033[H\033[2J")
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

func PrintTitle(){
    spaces := "                                             "
    for i:=0; i<37; i++ {
        fmt.Print("\033[H\033[2J")
        fmt.Print(`
        ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
        ██░▄▄▀█░▄▄▀█░▄▄█░▄▄▀█░▄▀█▀▄▀█░▄▄▀█░██░█░▄▀▄░█░▄▄▀█░▄▄
        ██░▄▄▀█░▀▀▄█░▄▄█░▀▀░█░█░█░█▀█░▀▀▄█░██░█░█▄█░█░▄▄▀█▄▄▀
        ██░▀▀░█▄█▄▄█▄▄▄█▄██▄█▄▄███▄██▄█▄▄██▄▄▄█▄███▄█▄▄▄▄█▄▄▄
        ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
        `)


        fmt.Println()
        if i % 2 == 0 {
            fmt.Print(
                spaces+" ,       \n"+
                spaces+":@.o.(_)\n"+
                spaces+"/  |  |\\\n")
        }else{
            fmt.Print(
                spaces+" ,      \n"+
                spaces+":@.o.(_)\n"+
                spaces+" | | /\\\n")
        }
        spaces = string(spaces[1:])
		time.Sleep(60 * time.Millisecond)
    }
}

func GenBoard(size int, pRed, pWhite float64) [][]int {
    board := make([][]int, size)
    for i := range board {
        board[i] = make([]int, size)
        for j := range board[i] {
            r := rand.Float64()
            if r < pRed {
                board[i][j] = 2
            } else if r >= pRed && r < pRed + pWhite {
                board[i][j] = 0
            } else {
                board[i][j] = 1
            }
        }
    }
    board[0][0] = 0
    board[size-1][size-1] = 3
    return board
}

func InputBoardSize() int{
    var size int
    for {
        fmt.Print("Enter Board Size (default 6) >")
        n, err := fmt.Fscanln(os.Stdin, &size)
        if (n ==0){
            return 6
        }else if err == nil && (size > 3 && size < 40){
            return size
        }
		time.Sleep(100 * time.Millisecond)
    }
}

func RunGame(state State, board [][]int, kill chan bool) {
    for {
        select{
        case <- kill:
            return
        default :
            fmt.Print("\033[H\033[2J")
            if state.color == red {
                return
            } else if state.color == green {
                return
            }
            PrintState(state, board)
            time.Sleep(200 * time.Millisecond)
            state = nextState(state, board)
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    PrintTitle()
    size := InputBoardSize()

    pRed := 1/float64(2*size)
    pWhite := (1 - pRed)/2
    board := GenBoard(size, pRed, pWhite)
	start := State{board[0][0], 1, 0, 0}
    PrintState(start, board)

    done := make(chan bool, 1)
    for {
        go RunGame(start,board,done)
        fmt.Scanln()
        done <- true
        time.Sleep(200 * time.Millisecond)
        fmt.Print("\033[H\033[2J")
    }
}
