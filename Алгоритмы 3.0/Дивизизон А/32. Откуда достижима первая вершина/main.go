package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var v, e int
	_, _ = fmt.Sscanf(line, "%d %d", &v, &e)

	graph := make([][]int, v+1)

	for i := 0; i < e; i++ {
		line, _ = reader.ReadString('\n')

		var start, end int
		_, _ = fmt.Sscanf(line, "%d %d", &end, &start)

		graph[start] = append(graph[start], end)
	}

	hash := make([]bool, v+1)

	CheckVertex(1, graph, hash)

	for i := range hash {
		if hash[i] {
			fmt.Printf("%d ", i)
		}
	}
}

func CheckVertex(vertex int, graph [][]int, checked []bool) {
	if checked[vertex] {
		return
	}

	checked[vertex] = true

	for i := range graph[vertex] {
		CheckVertex(graph[vertex][i], graph, checked)
	}
}
