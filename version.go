package version

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Version contains numbers and formatting for a version string
type Version struct {
	Numbers     []int
	Separators  []string
	NumberFirst bool
}

var (
	ErrNoNumbers = errors.New("version number contains no numbers")
)

// Parse will parse a version number from the supplied string. Any non-numbers are parsed out
// in their relative positions as formatting, so it should be format agnostic.
func Parse(x string) *Version {
	v := new(Version)
	var currentNumber string
	var currentSeparator string

	for i, char := range x {
		if unicode.IsDigit(char) {
			if i == 0 {
				v.NumberFirst = true
			}
			if currentSeparator != "" {
				v.Separators = append(v.Separators, currentSeparator)
				currentSeparator = ""
			}
			currentNumber += string(char)
		} else {
			if currentNumber != "" {
				n, _ := strconv.Atoi(currentNumber)
				v.Numbers = append(v.Numbers, n)
				currentNumber = ""
			}
			currentSeparator += string(char)
		}
	}

	if currentNumber != "" {
		n, _ := strconv.Atoi(currentNumber)
		v.Numbers = append(v.Numbers, n)
	}
	if currentSeparator != "" {
		v.Separators = append(v.Separators, currentSeparator)
	}

	return v
}

// Vaild will check to see if the parsed version has any numbers
func (v Version) Vaild() error {
	if len(v.Numbers) < 1 {
		return ErrNoNumbers
	}
	return nil
}

// ResetSmaller will reset all numbers to zero at the
// beyond the provided index.
func (v *Version) ResetSmaller(i int) {
	for ; i < len(v.Numbers); i++ {
		v.Numbers[i] = 0
	}
}

// SetAndResetSmaller will set the index i to value x and reset
// smaller version numbers to 0.
func (v *Version) SetAndResetSmaller(i, x int) {
	if i < len(v.Numbers) {
		v.Numbers[i] = x
	}
	v.ResetSmaller(i + 1)
}

// AddAndResetSmaller will change the number at index i by the
// value x.
func (v *Version) AddAndResetSmaller(i, x int) {
	if i < len(v.Numbers) {
		v.Numbers[i] += x
	}
	v.ResetSmaller(i + 1)
}

// String will convert the a parsed version back into string format
func (v Version) String(padding []int) string {
	var numi, sepi int
	numbernext := v.NumberFirst
	b := new(strings.Builder)

	for {
		if numbernext {
			if numi >= len(v.Numbers) {
				break
			}
			var ns string
			if padding != nil && len(padding) > numi {
				// Wish there was a better way to do this using the standard lib
				format := fmt.Sprintf("%%0%dd", padding[numi])
				ns = fmt.Sprintf(format, v.Numbers[numi])
			} else {
				ns = strconv.Itoa(v.Numbers[numi])
			}
			b.WriteString(ns)
			numi++
		} else {
			if sepi >= len(v.Separators) {
				break
			}
			b.WriteString(v.Separators[sepi])
			sepi++
		}
		numbernext = !numbernext
	}

	return b.String()
}
