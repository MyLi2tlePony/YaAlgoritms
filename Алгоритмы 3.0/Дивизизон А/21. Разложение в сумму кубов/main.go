package main

import "fmt"

var cubes = make([]int, 0)
var hashCubes = make(map[int]struct{})

func main() {
	var number int
	_, _ = fmt.Scanln(&number)

	for i := 1; i*i*i <= number; i++ {
		cubes = append(cubes, i*i*i)
		hashCubes[i*i*i] = struct{}{}
	}

	level := 1

	for {
		if FindTerms(number, level) {
			break
		}

		level++
	}

	fmt.Println(level)
}

func FindTerms(number, level int) bool {
	if _, ok := hashCubes[number]; ok {
		return true
	}

	if level > 1 {
		for i := range cubes {
			if FindTerms(number-cubes[i], level-1) {
				return true
			}
		}
	}

	return false
}
