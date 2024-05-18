package version

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestProcessStepBase(t *testing.T) {
	t.Parallel()
	test := func(provided, base, expected string) {
		c := &Config{
			Base: Parse(base),
		}
		v := Parse(provided)
		ProcessStepBase(c, v)
		compareVersions(t, Parse(expected), v)
	}

	test("1.2.3", "1.2", "1.2.4")
	test("1.2.3", "1", "1.3.0")
	test("1.2.3.4", "1.2.3", "1.2.3.5")
	test("5.5", "4", "4.0")
	test("6.7", "6", "6.8")
	test("5.66.777", "5", "5.67.0")
}

//TODO: test ProcessStepIncrement
//TODO: test ProcessStepSet
//TODO: test ProcessStepMinimum
//TODO: test ProcessVersionString

func compareVersions(t *testing.T, expected, gotten *Version) {
	t.Helper()
	if msg := cmp.Diff(expected, gotten); len(msg) > 0 {
		t.Error(msg)
	}
}
