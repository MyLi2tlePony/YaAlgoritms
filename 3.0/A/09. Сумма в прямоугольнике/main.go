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

	var n, m, k int
	_, _ = fmt.Sscanf(line, "%d %d %d", &n, &m, &k)

	arr := make([][]int, n+1)
	arr[0] = make([]int, m+1)

	for i := 1; i <= n; i++ {
		arr[i] = make([]int, m+1)

		line, _ = reader.ReadString('\n')
		for j, num := range strings.Fields(line) {
			number, _ := strconv.Atoi(num)

			arr[i][j+1] = number + arr[i][j] + arr[i-1][j+1] - arr[i-1][j]
		}
	}

	writer := bufio.NewWriter(os.Stdout)

	for i := 0; i < k; i++ {
		line, _ = reader.ReadString('\n')

		var x1, y1, x2, y2 int
		_, _ = fmt.Sscanf(line, "%d %d %d %d", &x1, &y1, &x2, &y2)

		_, _ = fmt.Fprintln(writer, arr[x2][y2]-arr[x2][y1-1]-arr[x1-1][y2]+arr[x1-1][y1-1])
	}

	_ = writer.Flush()
}
