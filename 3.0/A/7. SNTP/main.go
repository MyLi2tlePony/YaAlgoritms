package main

import (
	"fmt"
)

func main() {
	var hour, min, sec int
	var line string

	_, _ = fmt.Scanln(&line)
	_, _ = fmt.Sscanf(line, "%d:%d:%d", &hour, &min, &sec)
	A := toSec(hour, min, sec)

	_, _ = fmt.Scanln(&line)
	_, _ = fmt.Sscanf(line, "%d:%d:%d", &hour, &min, &sec)
	B := toSec(hour, min, sec)

	_, _ = fmt.Scanln(&line)
	_, _ = fmt.Sscanf(line, "%d:%d:%d", &hour, &min, &sec)
	C := toSec(hour, min, sec)

	delta := C - A
	if C < A {
		delta += 24 * 60 * 60
	}

	delta += delta % 2

	B += delta / 2
	B = B % (24 * 60 * 60)
	hour, min, sec = toTime(B)
	fmt.Printf("%02d:%02d:%02d", hour, min, sec)
}

func toSec(hour, min, sec int) int {
	return (hour*60+min)*60 + sec
}

func toTime(seconds int) (int, int, int) {
	sec := seconds % 60
	min := ((seconds - sec) % (60 * 60)) / 60
	hour := (seconds - sec - min*60) / (60 * 60)

	return hour, min, sec
}
