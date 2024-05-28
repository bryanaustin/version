package version

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// TestProcessStepBase will test ProcessStepBase
func TestProcessStepBase(t *testing.T) {
	t.Parallel()
	test := func(t *testing.T, provided, base, expected string) {
		t.Helper()
		bv := Parse(base)
		fv := Parse(provided)
		fv = ProcessStepBase(*fv, *bv)
		compareVersions(t, Parse(expected), fv)
	}

	test(t, "1.2.3", "1.2", "1.2.4")
	test(t, "1.2.3", "1", "1.3.0")
	test(t, "1.2.3.4", "1.2.3", "1.2.3.5")
	test(t, "5.5", "4", "4.0")
	test(t, "6.7", "6", "6.8")
	test(t, "5.66.777", "5", "5.67.0")
}

// TestProcessStepIncrement will test ProcessStepIncrement
func TestProcessStepIncrement(t *testing.T) {
	t.Parallel()
	test := func(t *testing.T, provided, base, expected string) {
		t.Helper()
		bv := Parse(base)
		fv := Parse(provided)
		fv = ProcessStepIncrement(*fv, *bv)
		compareVersions(t, Parse(expected), fv)
	}

	test(t, "1.2.3", "0.1", "1.3.0")
	test(t, "1.2.3", "0.0.2", "1.2.5")
	test(t, "1.2.3", "1", "2.0.0")
	test(t, "1.2.3", "1.2", "2.2.0")
	test(t, "1.2.3", "1", "2.0.0")
	test(t, "1.2.3.4", "0.0.1", "1.2.4.0")
	test(t, "5.5", "4", "9.0")
	test(t, "6.7", "0.6", "6.13")
}

// TestProcessStepSet will test ProcessStepSet
func TestProcessStepSet(t *testing.T) {
	t.Parallel()
	test := func(t *testing.T, provided, base, expected string) {
		t.Helper()
		bv := Parse(base)
		fv := Parse(provided)
		fv = ProcessStepSet(*fv, *bv)
		compareVersions(t, Parse(expected), fv)
	}

	test(t, "1.2.3", "0.1", "1.1.0")
	test(t, "1.2.3", "0.0.9", "1.2.9")
	test(t, "1.2.3", "1", "1.0.0")
	test(t, "1.2.3", "5.6", "5.6.0")
	test(t, "1.2.3", "9", "9.0.0")
	test(t, "1.2.3.4", "0.0.1", "1.2.1.0")
	test(t, "5.5", "4", "4.0")
	test(t, "6.7", "0.6", "6.6")
}

// TestProcessStepMinimum will test ProcessStepMinimum
func TestProcessStepMinimum(t *testing.T) {
	t.Parallel()
	test := func(t *testing.T, provided, base, expected string) {
		t.Helper()
		bv := Parse(base)
		fv := Parse(provided)
		fv = ProcessStepMinimum(*fv, *bv)
		compareVersions(t, Parse(expected), fv)
	}

	test(t, "1.2.3", "0.1", "1.2.3")
	test(t, "1.2.3", "0.0.9", "1.2.9")
	test(t, "1.2.3", "1", "1.2.3")
	test(t, "1.2.3", "5.6", "5.6.3")
	test(t, "1.2.3", "9", "9.2.3")
	test(t, "1.2.3.4", "0.0.1", "1.2.3.4")
	test(t, "5.5", "4", "5.5")
}

// TestGreaterThan will test GreatThan
func TestGreaterThan(t *testing.T) {
	t.Parallel()
	test := func(t *testing.T, provided, base string, expected bool) {
		t.Helper()
		fv := Parse(provided)
		bv := Parse(base)
		if result := fv.GreatThan(*bv); result != expected {
			t.Errorf("Expected %t, got %t", expected, result)
		}
	}

	test(t, "1.2.3", "1.2.3", false)
	test(t, "1.2.3", "1.2.0", true)
	test(t, "1.2.3", "1.2.4", false)
	test(t, "2.1.1", "1.2.3", true)
	test(t, "1.2.3", "1.3.1", false)
	test(t, "1.2.0", "1.2", false)
	test(t, "1.2.1", "1.2", true)
	test(t, "1.2", "1.2.3", false)
}

// TestLesserThan will test LessThan
func TestLesserThan(t *testing.T) {
	t.Parallel()
	test := func(t *testing.T, provided, base string, expected bool) {
		t.Helper()
		fv := Parse(provided)
		bv := Parse(base)
		if result := fv.LessThan(*bv); result != expected {
			t.Errorf("Expected %t, got %t", expected, result)
		}
	}

	test(t, "1.2.3", "1.2.3", false)
	test(t, "1.2.3", "1.2.0", false)
	test(t, "1.2.3", "1.2.4", true)
	test(t, "2.1.1", "1.2.3", false)
	test(t, "1.2.3", "1.3.1", true)
	test(t, "1.2.0", "1.2", false)
	test(t, "1.2.1", "1.2", false)
	test(t, "1.3", "1.2.3", false)
	test(t, "1.2", "1.2.3", true)
}

// TestNormalizeNumber will test the normalizeNumbers function
func TestNormalizeNumber(t *testing.T) {
	t.Parallel()
	test := func(t *testing.T, p1, p2, e1, e2 []int) {
		t.Helper()
		r1, r2 := normalizeNumbers(p1, p2)
		if msg := cmp.Diff(r1, e1); len(msg) > 0 {
			t.Error(msg)
		}
		if msg := cmp.Diff(r2, e2); len(msg) > 0 {
			t.Error(msg)
		}
	}

	var in1, in2 []int
	var ot1, ot2 []int

	in1 = []int{1, 2, 3}
	in2 = []int{4, 5, 6}
	ot1 = []int{1, 2, 3}
	ot2 = []int{4, 5, 6}
	test(t, in1, in2, ot1, ot2)

	in1 = []int{1, 2}
	in2 = []int{4, 5, 6}
	ot1 = []int{1, 2, 0}
	ot2 = []int{4, 5, 6}
	test(t, in1, in2, ot1, ot2)

	in1 = []int{1, 2, 3}
	in2 = []int{4}
	ot1 = []int{1, 2, 3}
	ot2 = []int{4, 0, 0}
	test(t, in1, in2, ot1, ot2)
}

//TODO: Test ProcessVersionString

func compareVersions(t *testing.T, expected, gotten *Version) {
	t.Helper()
	if msg := cmp.Diff(expected, gotten); len(msg) > 0 {
		t.Error(msg)
	}
}
