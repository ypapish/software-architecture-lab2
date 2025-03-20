package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func isOperator(element string) bool {
	return element == "+" || element == "-" || element == "*" || element == "/" || element == "^"
}

func isNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func PostfixToLisp(expression string) (string, error) {
	var stack []string
	elements := strings.Fields(expression)

	for _, element := range elements {
		if !isNumber(element) && !isOperator(element) {
			return "", fmt.Errorf("invalid expression: invalid element '%s'", element)
		}

		if isOperator(element) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: incorrect number of operands for '%s'", element)
			}

			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			operator := element
			if element == "^" {
				operator = "pow"
			}

			lispExpr := fmt.Sprintf("(%s %s %s)", operator, operand1, operand2)
			stack = append(stack, lispExpr)
		} else {
			stack = append(stack, element)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression: incorrect number of elements")
	}
	return stack[0], nil
}
