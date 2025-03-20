package lab2

import (
	"errors"
	"io"
	"strings"
)

var ErrInvalidExpression = errors.New("invalid expression")

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	expression := strings.TrimSpace(string(data))

	if expression == "" {
		return ErrInvalidExpression
	}

	result, err := PostfixToLisp(expression)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(result + "\n"))
	return err
}
