package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var vertexNumber, edgesNumber int
	_, _ = fmt.Sscanf(line, "%d %d", &vertexNumber, &edgesNumber)

	graphList := make([][]int, vertexNumber)
	edges := make([][]int, vertexNumber)
	for i := range edges {
		edges[i] = make([]int, vertexNumber)
	}

	for i := 1; i <= edgesNumber; i++ {
		line, _ = reader.ReadString('\n')

		var start, end int
		_, _ = fmt.Sscanf(line, "%d %d", &start, &end)

		start--
		end--

		graphList[start] = append(graphList[start], end)
		graphList[end] = append(graphList[end], start)

		edges[start][end] = i
		edges[end][start] = i
	}

	line, _ = reader.ReadString('\n')

	var robotsNumber int
	_, _ = fmt.Sscanf(line, "%d", &robotsNumber)

	robotsStarts := make([]int, robotsNumber)

	line, _ = reader.ReadString('\n')
	for i, num := range strings.Fields(line) {
		robotsStarts[i], _ = strconv.Atoi(num)
		robotsStarts[i]--
	}

	step := make([][]bool, vertexNumber)
	for i := range step {
		step[i] = make([]bool, robotsNumber)
	}
	for i := range robotsStarts {
		step[robotsStarts[i]][i] = true
	}

	for i := 0; i < vertexNumber; i++ {
		if FindCommonVertex(step) {
			fmt.Println(i)
			return
		}

		if FindCommonEdges(robotsNumber, edgesNumber, step, edges, graphList) {
			fmt.Printf("%d.5", i)
			return
		}

		step = NextStep(step, graphList)
	}

	fmt.Println(-1)
}

func FindCommonVertex(step [][]bool) bool {
	for i := range step {
		find := true

		for j := range step[i] {
			if !step[i][j] {
				find = false
			}
		}

		if find {
			return true
		}
	}

	return false
}

func FindCommonEdges(robotNumber, edgesNumber int, step [][]bool, edges, graphList [][]int) bool {
	nextStep := make([][]bool, edgesNumber)
	for i := range nextStep {
		nextStep[i] = make([]bool, robotNumber)
	}

	for start := range step {
		for robot := range step[start] {
			if !step[start][robot] {
				continue
			}

			for _, end := range graphList[start] {
				if edges[start][end] != 0 {
					nextStep[edges[start][end]-1][robot] = true
				}
			}
		}
	}

	return FindCommonVertex(nextStep)
}

func NextStep(step [][]bool, graphList [][]int) [][]bool {
	nextStep := make([][]bool, len(step))
	for i := range nextStep {
		nextStep[i] = make([]bool, len(step[i]))
	}

	for start := range step {
		for robot := range step[start] {
			if !step[start][robot] {
				continue
			}

			for _, end := range graphList[start] {
				nextStep[end][robot] = true
			}
		}
	}

	return nextStep
}
