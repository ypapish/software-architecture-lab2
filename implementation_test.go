package lab2

import (
	"fmt"
	"testing"
)

func TestPostfixToLisp(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		hasError bool
	}{
		{"4 5 +", "(+ 4 5)", false},
		{"3 2 ^", "(pow 3 2)", false},
		{"3.1 2.17 -", "(- 3.1 2.17)", false},
		{"3 2 ^ 3 /", "(/ (pow 3 2) 3)", false},
		{"4 2 - 3 2 ^ * 5 +", "(+ (* (- 4 2) (pow 3 2)) 5)", false},
		{"2 3 + 4 5 * 6 - ^ 7 8 + /", "(/ (pow (+ 2 3) (- (* 4 5) 6)) (+ 7 8))", false},
		{"3 4 5 * 6 2 / + ^ 7 -", "(- (pow 3 (+ (* 4 5) (/ 6 2))) 7)", false},
		{"10 2 8 * + 3 5 1 - ^ /", "(/ (+ 10 (* 2 8)) (pow 3 (- 5 1)))", false},
		{"", "invalid expression: incorrect number of elements", true},
		{"w b +", "invalid expression: invalid element 'w'", true},
		{"4 - 3 2 ^ ", "invalid expression: incorrect number of operands for '-'", true},
	}

	for _, test := range tests {
		lispExpr, err := PostfixToLisp(test.input)

		if test.hasError {
			if err == nil {
				t.Errorf("for input %q, expected error but got nil", test.input)
			} else if err.Error() != test.expected {
				t.Errorf("for input %q, expected error %q but got %q", test.input, test.expected, err.Error())
			}
		} else {
			if err != nil {
				t.Errorf("for input %q, unexpected error: %v", test.input, err)
			} else if lispExpr != test.expected {
				t.Errorf("for input %q, expected %q but got %q", test.input, test.expected, lispExpr)
			}
		}
	}
}

func ExamplePostfixToLisp() {
	res, _ := PostfixToLisp("2 2 +")
	fmt.Println(res)

	// Output:
	// (+ 2 2)
}
