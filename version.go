package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SemanticVersion []int

// Creates a new SemanticVersion given a string, e.g. "1.2.3",
// or returns an error if the string cannot be parsed correctly.
func NewVersion(str string) (SemanticVersion, error) {
	if str == "" {
		return nil, fmt.Errorf("Empty version string")
	}

	splitString := strings.Split(str, ".")
	result := make(SemanticVersion, len(splitString))
	for i, v := range splitString {
		if v == "" {
			return nil, fmt.Errorf("Partial version number is empty: %s", str)
		}

		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("Partial version number is non-numeric: %s", v)
		}
		if num < 0 {
			return nil, fmt.Errorf("Partial version number is negative: %s", v)
		}

		result[i] = num
	}

	return result, nil
}

const MAJOR = 0
const MINOR = 1
const PATCH = 2

// Returns the partial version at the specified index,
// or 0 if the index is beyond what is explicitly defined.
// TODO: Check requirements that trailing 0's do not have special meaning.
func (sm SemanticVersion) At(index int) int {
	if index < len(sm) {
		return (sm)[index]
	} else {
		return 0
	}
}

// Compares two semantic versions objects.
// Returns -1 if a < b, 0 if a == b, 1 if a > b
func Compare(a SemanticVersion, b SemanticVersion) int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	for i := 0; i < maxLen; i++ {
		if a.At(i) > b.At(i) {
			return 1
		} else if a.At(i) < b.At(i) {
			return -1
		}
	}

	return 0
}
