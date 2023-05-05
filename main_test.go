package main

import (
	"testing"
)

func TestReadArguments(t *testing.T) {
	tests := []struct {
		name           string
		input          []string
		expectedResult string
		// hasError bool
	}{
		// Test...
		{
			name:           "Should return production in \"-e production\" input",
			input:          []string{"script", "-e", "production"},
			expectedResult: "production",
		},
		// Test...
		{
			name:           "Should return production in \"--env production\" input",
			input:          []string{"script", "--env", "production"},
			expectedResult: "production",
		},
		// Test...
		{
			name:           "Should return development if no args",
			input:          []string{"script"},
			expectedResult: "development",
		},
		// Test...
		{
			name:           "Should return production in \"N prev args --env production\" input",
			input:          []string{"script", "argX", "argY", "argZ", "--env", "production"},
			expectedResult: "production",
		},
		// Test...
		{
			name:           "Should return development in \"--other production\" input",
			input:          []string{"script", "--other", "production"},
			expectedResult: "development",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := readArguments(tt.input)

			if result != tt.expectedResult {
				t.Errorf("Expected %s, got %s.", tt.expectedResult, result)
			}
		})
	}
}
