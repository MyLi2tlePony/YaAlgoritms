package main

import (
	"bufio"
	"fmt"
	"os"
)

type Os struct {
	start int
	end   int
}

func main() {
	var m int
	_, _ = fmt.Scanln(&m)

	var n int
	_, _ = fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(0)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	sectors := make([]Os, 0)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')

		var start, end int
		_, _ = fmt.Sscanf(line, "%d %d", &start, &end)

		for j := 0; j < len(sectors); j++ {
			if sectors[j].start <= end && start <= sectors[j].end {
				sectors[j], sectors[len(sectors)-1] = sectors[len(sectors)-1], sectors[j]
				sectors = sectors[:len(sectors)-1]
				j--
			}
		}

		sectors = append(sectors, Os{start: start, end: end})
	}

	fmt.Println(len(sectors))
}
