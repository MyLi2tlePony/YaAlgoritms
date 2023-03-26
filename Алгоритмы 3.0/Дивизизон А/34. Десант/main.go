package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var xs = []int{0, +1, 0, -1}
var ys = []int{+1, 0, -1, 0}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var rowNum, colNum int
	_, _ = fmt.Sscanf(line, "%d %d", &rowNum, &colNum)

	points := make([][]int, rowNum)
	hash := make([][]bool, rowNum)
	for i := range points {
		points[i] = make([]int, colNum)
		hash[i] = make([]bool, colNum)

		line, _ = reader.ReadString('\n')
		for j, num := range strings.Fields(line) {
			points[i][j], _ = strconv.Atoi(num)
		}
	}

	for i := range points {
		for j := range points[i] {
			MarkPoint(i, j, points, hash)
		}
	}

	result := 0

	for i := range points {
		for j := range points[i] {
			if !hash[i][j] {
				result++
				hash[i][j] = true

				for k := 0; k < 4; k++ {
					x := i + xs[k]
					y := j + ys[k]

					if x >= 0 && x < len(points) && y >= 0 && y < len(points[i]) {
						MarkPoint(x, y, points, hash)
					}
				}
			}
		}
	}

	fmt.Println(result)
}

func MarkPoint(i, j int, points [][]int, hash [][]bool) {
	if hash[i][j] {
		return
	}

	for k := 0; k < 4; k++ {
		x := i + xs[k]
		y := j + ys[k]

		if x >= 0 && x < len(points) && y >= 0 && y < len(points[i]) {
			if hash[x][y] && points[x][y] == points[i][j] || points[x][y] < points[i][j] {
				hash[i][j] = true
				break
			}
		}
	}

	if !hash[i][j] {
		return
	}

	for k := 0; k < 4; k++ {
		x := i + xs[k]
		y := j + ys[k]

		if x >= 0 && x < len(points) && y >= 0 && y < len(points[i]) {
			MarkPoint(x, y, points, hash)
		}
	}
}
