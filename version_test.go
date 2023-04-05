package main

import (
	"fmt"
	"testing"
)

func TestNewVersion(t *testing.T) {
	result, err := NewVersion("1.2")
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if result.At(MAJOR) != 1 || result.At(MINOR) != 2 {
		t.Error("Unexpected result:", result)
	}
}

func TestNewVersions(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		// errors
		{"", "Empty version string"},
		{"a", "Partial version number is non-numeric: a"},
		{"1.b", "Partial version number is non-numeric: b"},
		{"1.2.", "Partial version number is empty: 1.2."},
		{".", "Partial version number is empty: ."},
		{".1", "Partial version number is empty: .1"},
		{"1.-2", "Partial version number is negative: -2"},

		//good
		{"1", "[1]"},
		{"1.2", "[1 2]"},
		{"1.2.3", "[1 2 3]"},
		{"1.2.3.0", "[1 2 3 0]"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Version tests number %d", i), func(t *testing.T) {
			result, err := NewVersion(tt.input)
			if err == nil && tt.expected != fmt.Sprintf("%v", result) {
				t.Errorf("Test of %v: actual %v, expected %v",
					tt.input, result, tt.expected)
			}
			if err != nil && tt.expected != err.Error() {
				t.Errorf("Test of %v: actual %v, expected %v",
					tt.input, err.Error(), tt.expected)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	sm1, _ := NewVersion("1.2")
	sm2, _ := NewVersion("1.3")

	result := Compare(sm1, sm2)
	if result != -1 {
		t.Error("Expected 1 as result")
	}
}

func TestCompares(t *testing.T) {
	var tests = []struct {
		input1   string
		input2   string
		expected int
	}{
		{"0.1", "1.1", -1},
		{"1.1", "0.1", 1},
		{"0.1", "0.1", 0},
		{"0.1", "0.1.0.0", 0},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Compare tests number %d", i), func(t *testing.T) {
			sm1, _ := NewVersion(tt.input1)
			sm2, _ := NewVersion(tt.input2)
			result := Compare(sm1, sm2)
			if tt.expected != result {
				t.Errorf("Test of %v %v: actual %v, expected %v",
					tt.input1, tt.input2, result, tt.expected)
			}
		})
	}
}

func TestOrdered(t *testing.T) {
	input := []string{"0.1", "1.1", "1.2", "1.2.9.9.9.9", "1.3", "1.3.4", "1.10"}
	versions := make([]SemanticVersion, len(input))
	for i, v := range input {
		sm, _ := NewVersion(v)
		versions[i] = sm
	}

	for i := 1; i < len(versions)-1; i++ {
		if Compare(versions[i], versions[i+1]) != -1 {
			t.Errorf("%v compared to %v expected -1", versions[i], versions[i+1])
		}
	}
}
