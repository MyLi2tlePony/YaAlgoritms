package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var (
	dI = []int{2, 2, 1, -1}
	dJ = []int{-1, 1, 2, 2}
)

type XY struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var n, m int
	_, _ = fmt.Sscanf(line, "%d %d", &n, &m)

	fmt.Println(calcPath(0, 0, n, m, make(map[XY]*big.Int)))
}

func calcPath(i, j, r, c int, hash map[XY]*big.Int) *big.Int {
	if r, ok := hash[XY{i, j}]; ok {
		return r
	}

	if i == c-1 && j == r-1 {
		hash[XY{i, j}] = big.NewInt(1)
		return hash[XY{i, j}]
	}

	if i >= c || j >= r || i < 0 || j < 0 {
		hash[XY{i, j}] = big.NewInt(0)
		return hash[XY{i, j}]
	}

	result := big.NewInt(0)

	for k := range dI {
		result.Add(result, calcPath(i+dI[k], j+dJ[k], r, c, hash))
	}

	hash[XY{i, j}] = result
	return result
}
