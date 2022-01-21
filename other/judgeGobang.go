package main

import (
	"fmt"
)

const (
	black = -1
	white = 1
	empty = 0

	dir_x  = 1
	dir_y  = 2
	dir_xy = 3
	dir_yx = 4
)

func main() {
	board := [][]int{
   		    []int{1,0,0,0},
			[]int{0,1,0,0},
			[]int{0,0,1,-1},
			[]int{0,0,0,0},
	}
	for _,v :=range board {
		fmt.Println(v)
	}

	color := judgeWin(board)
	if color == black {
		fmt.Println("black win")
	}
	if color == white {
		fmt.Println("whit win")
	}
	if color == empty {
		fmt.Println("都没 win")
	}

}

func judgeWin(board [][]int) int {
	max_y := len(board) - 1
	max_x := len(board[0]) - 1
	for x := 0; x <= max_x; x++ {
		for y := 0; y <= max_y; y++ {
			now := board[y][x]
			if now == empty {
				continue
			}
			if judge(x, y, now, dir_x, 0, board) ||
				judge(x, y, now, dir_y, 0, board) ||
				judge(x, y, now, dir_xy, 0, board) ||
				judge(x, y, now, dir_yx, 0, board) {
				return now
			}
		}
	}
	return empty
}

func judge(x int, y int, state int, dir int, deep int, board [][]int) bool {
	// 2 == 长度3 index 0 开始
	if deep == 2 {
		return true
	}
	if dir == dir_x {
		x = x + 1
	}

	if dir == dir_y {
		y = y + 1
	}

	if dir == dir_xy {
		x = x + 1
		y = y + 1
	}
	if dir == dir_yx {
		x = x - 1
		y = y + 1
	}
	if x > len(board[0])-1 || x < 0 {
		return false
	}
	if y > len(board)-1 || y < 0 {
		return false
	}
	if board[y][x] != state {
		return false
	}
	return judge(x, y, state, dir, deep+1, board)
}
