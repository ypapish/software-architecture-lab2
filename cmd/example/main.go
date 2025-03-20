package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/ypapish/software-architecture-lab2"
)

var (
	inputExpression = flag.String("e", "", "Postfix expression to compute")
	inputFile       = flag.String("f", "", "File containing the postfix expression to compute")
	outputFile      = flag.String("o", "", "File to write the Lisp-like result to")
)

func main() {
	flag.Parse()

	if *inputExpression != "" && *inputFile != "" {
		fmt.Fprintln(os.Stderr, "Error: Both -e and -f flags cannot be used simultaneously")
		os.Exit(1)
	}

	var input io.Reader
	var output io.Writer

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "Error: No input provided. Use -e or -f flag.")
		os.Exit(1)
	}

	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	} else {
		output = os.Stdout
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error computing expression: %v\n", err)
		os.Exit(1)
	}
}
