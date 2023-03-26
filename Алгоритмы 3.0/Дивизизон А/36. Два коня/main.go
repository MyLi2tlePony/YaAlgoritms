package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	xS = []int{-2, -2, -1, -1, +1, +1, +2, +2}
	yS = []int{+1, -1, +2, -2, +2, -2, +1, -1}
)

type Board struct {
	arr [][]bool
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var horsePoint1, horsePoint2 string
	_, _ = fmt.Sscanf(line, "%s %s", &horsePoint1, &horsePoint2)

	horse1Board := CreateBoard()
	horse1Board.arr[horsePoint1[0]-'a'][horsePoint1[1]-'1'] = true

	horse2Board := CreateBoard()
	horse2Board.arr[horsePoint2[0]-'a'][horsePoint2[1]-'1'] = true

	for i := 0; i < 10; i++ {
		if BoardEquals(horse1Board, horse2Board) {
			fmt.Println(i)
			return
		}

		horse1Board = CreateBoardByBoard(horse1Board)
		horse2Board = CreateBoardByBoard(horse2Board)
	}

	fmt.Println(-1)
}

func BoardEquals(board1, board2 Board) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board1.arr[i][j] && board1.arr[i][j] == board2.arr[i][j] {
				return true
			}
		}
	}

	return false
}

func CreateBoardByBoard(board Board) Board {
	newBoard := CreateBoard()

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if !board.arr[i][j] {
				continue
			}

			for k := 0; k < 8; k++ {
				row := i + xS[k]
				col := j + yS[k]

				if row >= 0 && row < 8 && col >= 0 && col < 8 {
					newBoard.arr[row][col] = true
				}
			}
		}
	}

	return newBoard
}

func CreateBoard() Board {
	board := make([][]bool, 8)

	for i := range board {
		board[i] = make([]bool, 8)
	}

	return Board{board}
}
