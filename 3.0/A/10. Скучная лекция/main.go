package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	chars := make(map[byte]int)
	lastIndex := len(line) - 1

	for i := range line {
		subBefore := i
		subAfter := lastIndex - i

		chars[line[i]] += subBefore + subAfter + 1 + subAfter*subBefore
	}

	sortedChars := make([]byte, 0)

	for i := range chars {
		sortedChars = append(sortedChars, i)
	}

	sort.Slice(sortedChars, func(i, j int) bool {
		return sortedChars[i] < sortedChars[j]
	})

	for i := range sortedChars {
		fmt.Printf("%s: %d\n", string(sortedChars[i]), chars[sortedChars[i]])
	}
}
