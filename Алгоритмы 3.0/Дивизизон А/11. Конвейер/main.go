package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	n, _ := strconv.Atoi(strings.TrimRight(line, "\r\n"))

	for i := 0; i < n; i++ {
		_, _ = reader.ReadString(' ')
		line, _ = reader.ReadString('\n')
		dockersStr := strings.Fields(line)
		dockers := make([]float64, len(dockersStr))

		for i, docker := range dockersStr {
			d, _ := strconv.ParseFloat(docker, 64)
			dockers[i] = d
		}

		if validDockers(dockers) {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}

func validDockers(dockers []float64) bool {
	sortedDockers := make([]float64, len(dockers))

	for i := range dockers {
		sortedDockers[i] = dockers[i]
	}

	sort.Slice(sortedDockers, func(i, j int) bool {
		return sortedDockers[j] < sortedDockers[i]
	})

	warehouse := make([]float64, 0)

	for i := 0; i < len(dockers); {
		if sortedDockers[len(sortedDockers)-1] == dockers[i] {
			sortedDockers = sortedDockers[:len(sortedDockers)-1]
			i++
		} else if len(warehouse) > 0 && sortedDockers[len(sortedDockers)-1] == warehouse[len(warehouse)-1] {
			sortedDockers = sortedDockers[:len(sortedDockers)-1]
			warehouse = warehouse[:len(warehouse)-1]
		} else {
			warehouse = append(warehouse, dockers[i])
			i++
		}
	}

	for i := len(warehouse) - 1; i >= 0; i-- {
		if warehouse[i] == sortedDockers[len(sortedDockers)-1] {
			sortedDockers = sortedDockers[:len(sortedDockers)-1]
		} else {
			return false
		}
	}

	return true
}
