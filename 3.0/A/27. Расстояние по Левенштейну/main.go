package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	dpI = []int{0, -1, -1}
	dpJ = []int{-1, -1, 0}
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	line1, _ := reader.ReadString('\n')
	line1 = strings.TrimRight(line1, "\r\n")

	line2, _ := reader.ReadString('\n')
	line2 = strings.TrimRight(line2, "\r\n")

	dp := make([][]int, len(line1))
	for i := range dp {
		dp[i] = make([]int, len(line2))
	}

	for i := range dp {
		for j := range dp[i] {
			if line1[i] == line2[j] {
				if i == 0 {
					dp[i][j] = j
				} else if j == 0 {
					dp[i][j] = i
				} else {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				dp[i][j] = GetMin(dp, i, j) + 1
			}
		}
	}

	fmt.Println(dp[len(line1)-1][len(line2)-1])
}

func GetMin(dp [][]int, i, j int) int {
	result := -1

	for k := range dpI {
		if dpI[k]+i >= 0 && dpJ[k]+j >= 0 {
			current := dp[i+dpI[k]][j+dpJ[k]]

			if result == -1 || current < result {
				result = current
			}
		}
	}

	if result == -1 {
		return 0
	}

	return result
}
