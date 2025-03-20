package lab2

import (
	"bytes"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Valid expression",
			input:          "4 2 + 3 *",
			expectedOutput: "(* (+ 4 2) 3)",
			expectedError:  nil,
		},
		{
			name:           "Invalid expression",
			input:          "4 2 + *",
			expectedOutput: "",
			expectedError:  ErrInvalidExpression,
		},
		{
			name:           "Empty expression",
			input:          "",
			expectedOutput: "",
			expectedError:  ErrInvalidExpression,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := &bytes.Buffer{}

			handler := &ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()

			if tt.expectedError != nil {
				if err == nil || !strings.Contains(err.Error(), tt.expectedError.Error()) {
					t.Errorf("Expected error containing: %v, got: %v", tt.expectedError, err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if tt.expectedOutput != "" {
				actualOutput := strings.TrimSpace(output.String())
				if actualOutput != tt.expectedOutput {
					t.Errorf("Expected output: %s, got: %s", tt.expectedOutput, actualOutput)
				}
			}
		})
	}
}
