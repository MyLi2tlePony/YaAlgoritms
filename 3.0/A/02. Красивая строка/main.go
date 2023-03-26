package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var alphabet = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	k, _ := strconv.Atoi(strings.TrimRight(line, "\r\n"))

	line, _ = reader.ReadString('\n')
	maxBeauty := 0

	for i := range alphabet {
		curBeauty := calculateMaxBeauty(line, alphabet[i], k)
		if maxBeauty < curBeauty {
			maxBeauty = curBeauty
		}
	}

	fmt.Println(maxBeauty)
}

func calculateMaxBeauty(line string, char byte, k int) int {
	left, right, maxBeauty := 0, 0, 0

	for right < len(line) {
		for right < len(line) && (k > 0 || line[right] == char) {
			if line[right] != char {
				k--
			}

			right++
		}

		if maxBeauty < right-left {
			maxBeauty = right - left
		}

		for left < len(line) && line[left] == char {
			left++
		}

		left++
		k++
	}

	return maxBeauty
}
