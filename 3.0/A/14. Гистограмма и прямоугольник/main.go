package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Square struct {
	index int
	value int
	left  int
	right int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString(' ')

	squaresNumber, _ := strconv.Atoi(strings.TrimRight(line, " "))

	line, _ = reader.ReadString('\n')
	squares := make([]Square, squaresNumber)

	for i, n := range strings.Fields(line) {
		num, _ := strconv.Atoi(n)
		squares[i] = Square{value: num, index: i}
	}

	stack := make([]Square, 0)

	for i := range squares {
		for len(stack) > 0 && stack[len(stack)-1].value > squares[i].value {
			squares[stack[len(stack)-1].index].right = i
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, squares[i])
	}

	for len(stack) > 0 {
		squares[stack[len(stack)-1].index].right = len(squares)
		stack = stack[:len(stack)-1]
	}

	for i := len(squares) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1].value > squares[i].value {
			squares[stack[len(stack)-1].index].left = i
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, squares[i])
	}

	for len(stack) > 0 {
		squares[stack[len(stack)-1].index].left = -1
		stack = stack[:len(stack)-1]
	}

	maxSquare := 0

	for i := range squares {

		square := squares[i].value * (squares[i].right - squares[i].left - 1)

		if square > maxSquare {
			maxSquare = square
		}
	}

	fmt.Print(maxSquare)
}
