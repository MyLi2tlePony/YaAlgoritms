package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var regBrackets = regexp.MustCompile(`[()]`)
var regValidate = regexp.MustCompile(`[^ ()01!&|^]`)

var operandsValue = map[byte]int{'(': 0, ')': 0, '^': 1, '|': 1, '&': 2, '!': 3}
var openedBracket = map[string]string{")": "("}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	line = strings.TrimRight(line, "\r\n")

	if regValidate.MatchString(line) {
		fmt.Println("WRONG")
		return
	}

	brackets := make([]string, 0)
	for _, bracket := range regBrackets.FindAllString(line, -1) {
		b, ok := openedBracket[bracket]

		if len(brackets) == 0 || !ok || brackets[len(brackets)-1] != b {
			brackets = append(brackets, bracket)
		} else {
			brackets = brackets[:len(brackets)-1]
		}
	}

	if len(brackets) > 0 {
		fmt.Println("WRONG")
		return
	}

	hash := make([]byte, 0)
	operandHash := make([]byte, 0)

	for i := range line {
		switch line[i] {
		case '1', '0':
			hash = append(hash, line[i])
		case '&', '|', '^', '!':
			for len(operandHash) > 0 && operandsValue[operandHash[len(operandHash)-1]] >= operandsValue[line[i]] {
				hash = append(hash, operandHash[len(operandHash)-1])
				operandHash = operandHash[:len(operandHash)-1]
			}

			operandHash = append(operandHash, line[i])
		case '(':
			operandHash = append(operandHash, line[i])
		case ')':
			for operandHash[len(operandHash)-1] != '(' {
				hash = append(hash, operandHash[len(operandHash)-1])
				operandHash = operandHash[:len(operandHash)-1]
			}

			operandHash = operandHash[:len(operandHash)-1]
		}
	}

	for len(operandHash) > 0 {
		hash = append(hash, operandHash[len(operandHash)-1])
		operandHash = operandHash[:len(operandHash)-1]
	}

	result := make([]bool, 0)

	for i := range hash {
		switch hash[i] {
		case '0':
			result = append(result, false)
		case '1':
			result = append(result, true)
		case '!':
			if len(result) < 1 {
				fmt.Println("WRONG")
				return
			}

			result[len(result)-1] = !result[len(result)-1]
		case '&':
			if len(result) < 2 {
				fmt.Println("WRONG")
				return
			}

			result[len(result)-2] = result[len(result)-2] && result[len(result)-1]
			result = result[:len(result)-1]
		case '|':
			if len(result) < 2 {
				fmt.Println("WRONG")
				return
			}

			result[len(result)-2] = result[len(result)-2] || result[len(result)-1]
			result = result[:len(result)-1]
		case '^':
			if len(result) < 2 {
				fmt.Println("WRONG")
				return
			}

			result[len(result)-2] = result[len(result)-2] != result[len(result)-1]
			result = result[:len(result)-1]
		}
	}

	if len(result) != 1 {
		fmt.Println("WRONG")
		return
	}

	if result[0] {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
