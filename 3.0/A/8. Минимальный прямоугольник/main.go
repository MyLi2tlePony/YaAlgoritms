package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x, y, maxY, maxX, minY, minX int
	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	_, _ = fmt.Sscanf(line, "%d", &n)

	for i := 0; i < n; i++ {
		line, _ = reader.ReadString('\n')
		_, _ = fmt.Sscanf(line, "%d %d", &x, &y)

		if x > maxX || i == 0 {
			maxX = x
		}

		if x < minX || i == 0 {
			minX = x
		}

		if y > maxY || i == 0 {
			maxY = y
		}

		if y < minY || i == 0 {
			minY = y
		}
	}

	fmt.Println(minX, minY, maxX, maxY)
}
