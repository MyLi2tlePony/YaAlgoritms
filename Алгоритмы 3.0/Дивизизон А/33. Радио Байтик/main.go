package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var v int
	_, _ = fmt.Sscanf(line, "%d", &v)

	points := make([]Point, v)

	for i := 0; i < v; i++ {
		line, _ = reader.ReadString('\n')

		var x, y int
		_, _ = fmt.Sscanf(line, "%d %d", &x, &y)

		points[i] = Point{x, y}
	}

	matrix, distances := CalculateDistances(points)
	sort.Slice(distances, func(i, j int) bool {
		return distances[i] < distances[j]
	})

	radius, colors := FindRadiusAndColors(matrix, distances)

	fmt.Printf("%0.8f\n", radius)

	for i := range colors {
		fmt.Printf("%d ", colors[i])
	}
}

func Marked(index, distance int, colors []uint8, matrix [][]int) bool {
	for j := 0; j < len(matrix[index]); j++ {
		if j == index {
			continue
		}

		if matrix[index][j] < distance {
			if colors[j] == 0 {
				colors[j] = 3 - colors[index]

				if !Marked(j, distance, colors, matrix) {
					return false
				}
			} else if colors[index] == colors[j] {
				return false
			}
		}
	}

	return true
}

func FindRadiusAndColors(matrix [][]int, distances []int) (float64, []uint8) {
	var iasdf int
	for len(distances) > 1 {
		distance := distances[len(distances)/2]
		colors := make([]uint8, len(matrix))

		if distance == 112896 || distance == 112225 || len(distances) == 3 {
			iasdf = 10
			_ = iasdf
		}

		isMarked := true
		for i := range colors {
			if colors[i] == 0 {
				colors[i] = 1

				if !Marked(i, distance, colors, matrix) {
					isMarked = false
					break
				}
			}
		}

		if isMarked {
			distances = distances[len(distances)/2:]
		} else {
			distances = distances[:len(distances)/2]
		}

	}

	colors := make([]uint8, len(matrix))
	for i := range colors {
		if colors[i] == 0 {
			colors[i] = 1

			Marked(i, distances[0], colors, matrix)
		}
	}

	return math.Sqrt(float64(distances[0])) / 2, colors
}

func CalculateDistances(points []Point) ([][]int, []int) {
	matrix := make([][]int, len(points))
	for i := range matrix {
		matrix[i] = make([]int, len(points))
	}

	distances := make(map[int]struct{})

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			distance := (points[i].x-points[j].x)*(points[i].x-points[j].x) + (points[i].y-points[j].y)*(points[i].y-points[j].y)
			distances[distance] = struct{}{}

			matrix[i][j] = distance
			matrix[j][i] = distance
		}
	}

	arr := make([]int, len(distances))
	j := 0

	for i := range distances {
		arr[j] = i
		j++
	}

	return matrix, arr
}
