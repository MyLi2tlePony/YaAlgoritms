package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	var num int
	_, _ = fmt.Sscanf(line, "%d", &num)
	steps := []map[int]int{{num: 0}}

	line, _ = reader.ReadString('\n')
	_, _ = fmt.Sscanf(line, "%d", &num)

	for i := 0; true; i++ {
		if FindVal(num, steps[i]) {
			break
		}

		steps = append(steps, GenerateCheckPoints(steps[i]))
	}

	ShowResults(num, steps)
}

func ShowResults(val int, steps []map[int]int) {
	if len(steps) < 1 {
		return
	}

	prev := steps[len(steps)-1][val]
	steps = steps[:len(steps)-1]
	ShowResults(prev, steps)
	fmt.Println(val)
}

func GenerateCheckPoints(step map[int]int) map[int]int {
	generatedStep := make(map[int]int, 0)

	for i := range step {
		if i/1000 < 9 {
			generatedStep[i+1000] = i
		}

		if i%10 > 1 {
			generatedStep[i-1] = i
		}

		generatedStep[(i*10+i/1000)%10000] = i
		generatedStep[(i/10 + (i%10)*1000)] = i
	}

	return generatedStep
}

func FindVal(val int, steps map[int]int) bool {
	if _, ok := steps[val]; ok {
		return true
	}

	return false
}
