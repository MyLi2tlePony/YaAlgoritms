package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var rows, cols int
	_, _ = fmt.Sscanf(line, "%d %d", &rows, &cols)

	outputs := make([][]bool, rows)
	maze := make([][]bool, rows)
	steps := make([][]bool, rows)

	for i := range maze {
		maze[i] = make([]bool, cols)
		outputs[i] = make([]bool, cols)
		steps[i] = make([]bool, cols)

		line, _ = reader.ReadString('\n')
		for j, digit := range strings.Fields(line) {
			switch digit {
			case "1":
				maze[i][j] = true
			case "2":
				outputs[i][j] = true
			}
		}
	}

	steps[0][0] = true

	for i := 0; true; i++ {
		if FindOutput(steps, outputs) {
			fmt.Println(i)
			break
		}

		steps = NextStep(steps, maze, outputs)
	}
}

func NextStep(steps, maze, outputs [][]bool) [][]bool {
	newSteps := make([][]bool, len(steps))
	for i := range steps {
		newSteps[i] = make([]bool, len(steps[i]))
	}

	cR := []int{0, 0, 1, -1}
	cC := []int{1, -1, 0, 0}

	for i := range steps {
		for j := range steps[i] {
			if !steps[i][j] {
				continue
			}

			for k := 0; k < 4; k++ {
				r, c := FindLastEnabledPoint(i, j, cR[k], cC[k], maze, outputs)
				newSteps[r][c] = true
			}
		}
	}

	return newSteps
}

func FindLastEnabledPoint(i, j, cR, cC int, maze, outputs [][]bool) (int, int) {
	for true {
		if outputs[i][j] {
			return i, j
		}

		r, c := i+cR, j+cC

		if r >= 0 && r < len(maze) && c >= 0 && c < len(maze[r]) && !maze[r][c] {
			i = r
			j = c
		} else {
			return i, j
		}
	}

	return i, j
}

func FindOutput(steps, outputs [][]bool) bool {
	for i := range steps {
		for j := range steps[i] {
			if steps[i][j] && outputs[i][j] {
				return true
			}
		}
	}

	return false
}
