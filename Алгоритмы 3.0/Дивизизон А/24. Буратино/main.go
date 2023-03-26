package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	_, _ = fmt.Scanln(&n)

	reader := bufio.NewReader(os.Stdin)
	timePoints := make([]int, n)
	timePointValues := make([]int, n)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')

		var hour, min, sec, value int
		_, _ = fmt.Sscanf(line, "%d:%d:%d %d", &hour, &min, &sec, &value)

		timePoints[i] = (hour-9)*3600 + min*60 + sec
		timePointValues[i] = value
	}

	pins := CalcPins(0, timePoints, timePointValues)
	pins += CalcPins(5*3600, timePoints, timePointValues)

	fmt.Println(pins)
}

func CalcPins(start int, timePoints, timePointValues []int) int {
	values := make([]int, 4*3600)

	for j, i := 0, 0; i < len(values); i++ {
		for j < len(timePoints)-1 && i+start >= timePoints[j+1] {
			j++
		}

		values[i] = timePointValues[j]
	}

	powers := make([]int, 4*3600)
	indexMaxPower := 0

	for i := len(values) - 1; i >= 0; i-- {
		if values[i] <= len(values)-i {
			powers[i]++
		}

		maxPow := 0
		for j := i + values[i]; j < len(powers); j++ {
			if powers[j] > maxPow {
				maxPow = powers[j]
			}
		}

		powers[i] += maxPow

		if powers[indexMaxPower] < powers[i] {
			indexMaxPower = i
		}
	}

	return powers[indexMaxPower]
}
