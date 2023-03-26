package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	charFrequency := make(map[byte]int)

	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r\n")

		for i := range line {
			if line[i] != ' ' {
				charFrequency[line[i]]++
			}
		}
	}

	chars := make([]byte, 0)
	maxFrequency := 0

	for i := range charFrequency {
		if charFrequency[i] > maxFrequency {
			maxFrequency = charFrequency[i]
		}

		chars = append(chars, i)
	}

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	for i := maxFrequency - 1; i >= 0; i-- {
		builder := strings.Builder{}

		for j := range chars {
			if charFrequency[chars[j]] == maxFrequency {
				charFrequency[chars[j]]--
				builder.WriteByte('#')
			} else {
				builder.WriteByte(' ')
			}
		}

		maxFrequency--
		fmt.Println(builder.String())
	}

	builder := strings.Builder{}
	for i := range chars {
		builder.WriteByte(chars[i])
	}
	fmt.Println(builder.String())
}
