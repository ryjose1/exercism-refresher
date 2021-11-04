// Package encode holds tools to perform Run Length encoding and decoding
// This package assumes no unicode is in the data
package encode

import (
	"fmt"
	"strconv"
	"strings"
)

// GetNumPrefix separates the number characters and the following letter/whitespace character
// from the rest of the input
func GetNumPrefix(input string) (prefix string, suffix string) {
	for i, char := range input {
		if i+1 > len(input) {
			panic("Hit end of string without encountering a letter")
		}
		strconv.
		if _, err := strconv.Atoi(string(char)); err != nil {
			prefix = string(input[:i+1])
			suffix = string(input[i+1:])
			return
		}
	}
	panic("Hit end of string without encountering a letter")
}

// GetRepeatCharPrefix separates repeating letter/whitespace characters
// from the rest of the input
func GetRepeatCharPrefix(input string) (prefix string, suffix string) {
	firstChar := string(input[0])
	for i, char := range input {
		if string(char) != firstChar {
			prefix = string(input[:i])
			suffix = string(input[i:])
			return
		}
	}

	// If the string ends in repeating characters, return the whole thing as the prefix
	prefix = input
	return
}

// GetDiffCharPrefix separates non-repeating letter/whitespace characters
// from the rest of the input
func GetDiffCharPrefix(input string) (prefix string, suffix string) {
	for i, char := range input {
		if i+1 < len(input) {
			// See if the next char is a number
			if _, err := strconv.Atoi(string(input[i+1])); err == nil {
				prefix = string(input[:i+1])
				suffix = string(input[i+1:])
				return
			}
			// See if the next char is the same as the current char
			if string(input[i+1]) == string(char) {
				prefix = string(input[:i])
				suffix = string(input[i:])
				return
			}
		}

	}

	// If the string ends in non-repeating characters, return the whole thing as the prefix
	prefix = input
	return
}

// NextDataChunk separates the input string into a prefix and suffix
// prefix holds the longest part of the input that can have the same logic applied to it
// suffix holds the rest of the input
// convert holds whether the prefix is compressable/has been compressed
func NextDataChunk(input string) (prefix string, suffix string, convert bool) {
	if len(input) == 0 {
		return
	} else if len(input) == 1 {
		prefix = input
		return
	}

	// Handle numbers
	char := string(input[0])
	if _, err := strconv.Atoi(char); err == nil {
		prefix, suffix = GetNumPrefix(input)
		convert = true
		return
	}

	// Handle repeating characters
	nextChar := string(input[1])
	if nextChar == char {
		prefix, suffix = GetRepeatCharPrefix(input)
		convert = true
		return
	}

	// Handle differing characters
	prefix, suffix = GetDiffCharPrefix(input)
	convert = false
	return
}

func EncodeChunk(chunk string) string {
	repeats := len(chunk)
	char := chunk[0]
	return fmt.Sprint(repeats) + string(char)
}

func RunLengthEncode(input string) string {
	chunks := []string{}

	var prefix string
	var convert bool
	suffix := input
	for suffix != "" {
		prefix, suffix, convert = NextDataChunk(suffix)
		if convert {
			prefix = EncodeChunk(prefix)
		}
		chunks = append(chunks, prefix)
	}
	return strings.Join(chunks, "")
}

func DecodeChunk(chunk string) string {
	repeats, err := strconv.Atoi(string(chunk[:len(chunk)-1]))
	if err != nil {
		panic("Couldn't convert number portion of chunk")
	}
	char := string(chunk[len(chunk)-1])
	return strings.Repeat(char, repeats)
}

func RunLengthDecode(input string) string {
	chunks := []string{}

	var prefix string
	var convert bool
	suffix := input
	for suffix != "" {
		prefix, suffix, convert = NextDataChunk(suffix)
		if convert {
			prefix = DecodeChunk(prefix)
		}
		chunks = append(chunks, prefix)
	}
	return strings.Join(chunks, "")
}
