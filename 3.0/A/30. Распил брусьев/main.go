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

	var balkLen, partNumber int
	line, _ := reader.ReadString('\n')
	_, _ = fmt.Sscanf(line, "%d %d", &balkLen, &partNumber)

	parts := make([]int, partNumber+2)
	line, _ = reader.ReadString('\n')

	for i, num := range strings.Fields(line) {
		parts[i+1], _ = strconv.Atoi(num)
	}

	parts[len(parts)-1] = balkLen

	dp := make([][]int, len(parts))
	for i := range dp {
		dp[i] = make([]int, len(parts))
	}

	for j := 1; j < len(dp); j++ {
		for i := j - 2; i >= 0; i-- {
			dp[i][j] = 0

			for s := i + 1; s < j; s++ {
				current := dp[i][s] + dp[s][j]

				if current < dp[i][j] || s == i+1 {
					dp[i][j] = current
				}
			}

			dp[i][j] += parts[j] - parts[i]
		}
	}

	fmt.Println(dp[0][len(parts)-1])
}
