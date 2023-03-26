package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regValidateExpression = regexp.MustCompile(`(\d+ \d+)|[^\d( )+\-*]`)
var regExpression = regexp.MustCompile(`[()+\-*]|\d+`)

var operandsValue = map[string]int{"(": 0, ")": 0, "-": 1, "+": 1, "*": 2, "/": 2}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	if regValidateExpression.MatchString(strings.TrimRight(line, "\r\n")) {
		fmt.Println("WRONG")
		return
	}

	expression := regExpression.FindAllString(line, -1)
	operators := make([]string, 0)
	resultExpression := make([]string, 0)

	for i := range expression {
		if _, err := strconv.Atoi(expression[i]); err == nil {
			resultExpression = append(resultExpression, expression[i])
		} else {
			switch expression[i] {
			case "*", "/":
				for len(operators) > 0 && operandsValue[operators[len(operators)-1]] >= 2 {
					resultExpression = append(resultExpression, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}

				operators = append(operators, expression[i])
			case "-", "+":
				for len(operators) > 0 && operandsValue[operators[len(operators)-1]] >= 1 {
					resultExpression = append(resultExpression, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}

				operators = append(operators, expression[i])
			case "(":
				operators = append(operators, "(")
			case ")":
				for len(operators) > 0 && operators[len(operators)-1] != "(" {
					resultExpression = append(resultExpression, operators[len(operators)-1])
					operators = operators[:len(operators)-1]
				}

				if len(operators) < 1 {
					fmt.Println("WRONG")
					return
				}

				operators = operators[:len(operators)-1]
			}
		}
	}

	for i := len(operators) - 1; i >= 0; i-- {
		if operators[i] == "(" || operators[i] == ")" {
			fmt.Println("WRONG")
			return
		}

		resultExpression = append(resultExpression, operators[i])
	}

	result := make([]int, 0)

	for i := 0; i < len(resultExpression); i++ {
		switch resultExpression[i] {
		case "+":
			if len(result) < 2 {
				fmt.Println("WRONG")
				return
			}

			result[len(result)-2] = result[len(result)-2] + result[len(result)-1]
			result = result[:len(result)-1]
		case "-":
			if len(result) < 2 {
				fmt.Println("WRONG")
				return
			}

			result[len(result)-2] = result[len(result)-2] - result[len(result)-1]
			result = result[:len(result)-1]
		case "*":
			if len(result) < 2 {
				fmt.Println("WRONG")
				return
			}

			result[len(result)-2] = result[len(result)-2] * result[len(result)-1]
			result = result[:len(result)-1]
		default:
			num, _ := strconv.Atoi(resultExpression[i])
			result = append(result, num)
		}
	}

	if len(result) > 1 || len(result) < 1 {
		fmt.Println("WRONG")
		return
	}

	fmt.Println(result[0])
}
