package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanln(&n)

	result := 0

	for i := 1; i <= n; i++ {
		result += (1 + i) * i

		if i > n/2 {
			result -= (1 + (i - (n - i))) * (i - (n - i)) / 2
		}
	}

	fmt.Println(result)
}
