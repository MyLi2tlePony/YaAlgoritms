package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var n int
	_, _ = fmt.Sscanf(line, "%d", &n)

	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)

		line, _ = reader.ReadString('\n')
		for j := range strings.TrimRight(line, "\r\n") {
			if line[j] == '*' {
				grid[i][j] = true
			} else {
				grid[i][j] = false
			}
		}
	}

	line, _ = reader.ReadString('\n')

	var x, y int
	_, _ = fmt.Sscanf(line, "%d %d", &x, &y)

	hash := make(map[Point]struct{})
	CheckPoints(x-1, y-1, grid, hash)

	fmt.Println(len(hash))
}

func CheckPoints(x, y int, grid [][]bool, hash map[Point]struct{}) {
	if _, ok := hash[Point{x, y}]; ok || grid[x][y] {
		return
	}

	hash[Point{x, y}] = struct{}{}

	CheckPoints(x-1, y, grid, hash)
	CheckPoints(x+1, y, grid, hash)
	CheckPoints(x, y-1, grid, hash)
	CheckPoints(x, y+1, grid, hash)
}
