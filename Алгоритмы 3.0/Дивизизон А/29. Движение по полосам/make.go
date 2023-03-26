package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var m, n int
	_, _ = fmt.Sscanf(line, "%d %d", &m, &n)

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	dp[0][0] = 1

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] + dp[i-1][j-1]
		}
	}

	fmt.Println(dp[m][n])
}
