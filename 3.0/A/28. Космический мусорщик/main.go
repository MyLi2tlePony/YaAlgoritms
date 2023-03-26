package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var decoder = map[byte]int{'N': 0, 'S': 1, 'W': 2, 'E': 3, 'U': 4, 'D': 5}

func main() {
	reader := bufio.NewReader(os.Stdin)

	alphabet := make([][]int, len(decoder))

	for i := 0; i < 6; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")

		alphabet[i] = make([]int, len(line))
		for j := range line {
			alphabet[i][j] = decoder[line[j]]
		}
	}

	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\r\n")

	var let string
	var parameter int
	_, _ = fmt.Sscanf(line, "%s %d", &let, &parameter)

	hash := make([]map[int]int, len(decoder))
	for i := range hash {
		hash[i] = make(map[int]int)
	}

	fmt.Println(CalculateMoves(decoder[let[0]], parameter, hash, alphabet))
}

func CalculateMoves(let, parameter int, hash []map[int]int, alphabet [][]int) int {
	if parameter == 1 {
		return 1
	}

	if result, ok := hash[let][parameter]; ok {
		return result
	}

	result := 1

	for i := range alphabet[let] {
		result += CalculateMoves(alphabet[let][i], parameter-1, hash, alphabet)
	}

	hash[let][parameter] = result
	return result
}
