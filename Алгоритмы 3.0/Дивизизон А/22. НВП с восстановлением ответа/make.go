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

	var n int
	_, _ = fmt.Sscanf(line, "%d", &n)

	nums := make([]int, n)
	values := make([]int, n)
	steps := make([]int, n)

	line, _ = reader.ReadString('\n')

	for i, num := range strings.Fields(line) {
		nums[i], _ = strconv.Atoi(num)
	}

	indexMaxValue := 0

	for i := len(nums) - 1; i >= 0; i-- {
		steps[i] = len(nums)

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] && values[i] < values[j]+1 {
				values[i] = values[j] + 1
				steps[i] = j
			}
		}

		if values[i] >= values[indexMaxValue] {
			indexMaxValue = i
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	for i := indexMaxValue; i < len(nums); i = steps[i] {
		_, _ = fmt.Fprintf(writer, "%d ", nums[i])
	}
	_ = writer.Flush()
}
