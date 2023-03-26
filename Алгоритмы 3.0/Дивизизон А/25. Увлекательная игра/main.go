package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var n, a, b int
	_, _ = fmt.Sscanf(line, "%d %d %d", &n, &a, &b)

	maxCost := 0
	if a > b {
		maxCost = a
	} else {
		maxCost = b
	}

	dp := make([]int, n)

	for i := 1; i < len(dp); i++ {
		dp[i] = maxCost
		for k := 0; k < i; k++ {
			before := dp[k] + a
			after := dp[i-k] + b

			max := 0
			if before > after {
				max = before
			} else {
				max = after
			}

			if k == 0 || max < dp[i] {
				dp[i] = max
			}
		}
	}

	fmt.Println(dp[len(dp)-1])
}
