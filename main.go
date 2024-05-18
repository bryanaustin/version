package version

import (
	"fmt"
	"os"
	"strconv"
)

// Process will run the application
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

// ProcessVerisonString will run though all the steps for a provided version
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

// ProcessVersion will run all the steps the result in a newer version
func ProcessVersion(c *Config) *Version {
	result := new(Version)
	result.Numbers = append([]int(nil), c.Provided.Numbers...)
	result.Separators = append([]string(nil), c.Provided.Separators...)
	result.NumberFirst = c.Provided.NumberFirst

	ProcessStepBase(c, result)
	ProcessStepIncrement(c, result)
	ProcessStepSet(c, result)
	ProcessStepMinimum(c, result)

	return result
}

// ProcessStepBase increase the largest verison number that's smaller then base version.
func ProcessStepBase(c *Config, v *Version) {
	if c.Base != nil {
		// check to see if anything has changed between the versions
		equal := true
		for i := range c.Base.Numbers {
			if v.Numbers[i] != c.Base.Numbers[i] {
				equal = false
				v.Numbers[i] = c.Base.Numbers[i]
			}
		}

		if equal {
			// increment
			v.AddAndResetSmaller(len(c.Base.Numbers), 1)
		} else {
			// reset
			v.ResetSmaller(len(c.Base.Numbers))
		}
	}
}

func ProcessStepIncrement(c *Config, v *Version) {
	if c.Increment != nil {
		for i, x := range c.Increment.Numbers {
			if x > 0 {
				v.AddAndResetSmaller(i, x)
			}
		}
	}
}

func ProcessStepSet(c *Config, v *Version) {
	if c.Set != nil {
		for i, x := range c.Set.Numbers {
			if x > 0 {
				v.SetAndResetSmaller(i, x)
			}
		}
	}
}

func ProcessStepMinimum(c *Config, v *Version) {
	if c.Minimum != nil {
		for i, x := range c.Minimum.Numbers {
			if len(v.Numbers) < i {
				break
			}
			if x > v.Numbers[i] {
				v.Numbers[i] = x
			}
		}
	}
}
