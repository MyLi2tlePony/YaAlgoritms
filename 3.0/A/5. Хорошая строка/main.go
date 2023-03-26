package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanln(&n)

	chars := make([]int, n)

	for i := range chars {
		_, _ = fmt.Scanln(&chars[i])
	}

	result := 0

	for i := 1; i < len(chars); i++ {
		if chars[i-1] != 0 {
			if chars[i-1] < chars[i] {
				result += chars[i-1]
			} else {
				result += chars[i]
			}
		}
	}

	fmt.Println(result)
}
