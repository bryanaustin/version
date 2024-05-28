package version

import (
	"fmt"
	"os"
	"strconv"
)

// Process will run the version application as intended
func Process(c *Config) {
	r, ok := ProcessConditional(c)
	if ok {
		fmt.Println(strconv.FormatBool(r))
		if !r {
			os.Exit(100)
		}
		return
	}

	s := ProcessVersionString(c)
	fmt.Println(s)
}

// ProcessConditional will check lesser or greater options
func ProcessConditional(c *Config) (result bool, ok bool) {
	if c.Lesser != nil {
		ok = true
		result = c.Provided.LessThan(*c.Lesser)
		return
	}
	if c.Greater != nil {
		ok = true
		result = c.Provided.GreatThan(*c.Greater)
		return
	}
	return
}

// ProcessVerisonString will run though all the intended steps to generate the configured string result
func ProcessVersionString(c *Config) string {
	v := ProcessVersion(c)
	if c.Format != nil {
		v.Separators = c.Format.Separators
		v.NumberFirst = c.Format.NumberFirst
	}

	var padding []int
	if c.Pad != nil {
		padding = c.Pad.Numbers
	}
	return v.String(padding)
}

// ProcessVersion will run all the steps the intended steps to generate the configured version
func ProcessVersion(c *Config) *Version {
	result := new(Version)
	result.Numbers = append([]int(nil), c.Provided.Numbers...)
	result.Separators = append([]string(nil), c.Provided.Separators...)
	result.NumberFirst = c.Provided.NumberFirst

	if c.Base != nil {
		result = ProcessStepBase(*result, *c.Base)
	}

	if c.Increment != nil {
		result = ProcessStepIncrement(*result, *c.Increment)
	}

	if c.Set != nil {
		result = ProcessStepSet(*result, *c.Set)
	}

	if c.Minimum != nil {
		result = ProcessStepMinimum(*result, *c.Minimum)
	}

	return result
}

// ProcessStepBase increase the largest verison number that's smaller then base version.
func ProcessStepBase(v, x Version) *Version {
	vCopy := v.Copy()
	r := &vCopy
	// check to see if anything has changed between the versions
	equal := true
	for i := range x.Numbers {
		if r.Numbers[i] != x.Numbers[i] {
			equal = false
			r.Numbers[i] = x.Numbers[i]
		}
	}

	if equal {
		// increment
		r.AddAndResetSmaller(len(x.Numbers), 1)
	} else {
		// reset
		r.ResetSmaller(len(x.Numbers))
	}
	return r
}

func ProcessStepIncrement(v, x Version) *Version {
	vCopy := v.Copy()
	r := &vCopy
	for i, val := range x.Numbers {
		if val > 0 {
			r.AddAndResetSmaller(i, val)
		}
	}
	return r
}

func ProcessStepSet(v, x Version) *Version {
	vCopy := v.Copy()
	r := &vCopy
	for i, val := range x.Numbers {
		if val > 0 {
			r.SetAndResetSmaller(i, val)
		}
	}
	return r
}

func ProcessStepMinimum(v, x Version) *Version {
	vCopy := v.Copy()
	r := &vCopy
	for i, val := range x.Numbers {
		if len(r.Numbers) < i {
			break
		}
		if val > r.Numbers[i] {
			r.Numbers[i] = val
		}
	}
	return r
}
