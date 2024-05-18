package version

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBunchOParsing(t *testing.T) {
	t.Parallel()
	compareParsedVerison(t, &Version{Numbers: []int{1, 2}, Separators: []string{"."}, NumberFirst: true}, "1.2", "1.2")
	compareParsedVerison(t, &Version{Numbers: []int{92}, NumberFirst: true}, "92", "92")
	compareParsedVerison(t, &Version{Numbers: []int{2}, NumberFirst: true}, "2", "000000002")
	compareParsedVerison(t, &Version{Numbers: []int{2, 1, 0}, Separators: []string{".", "."}, NumberFirst: true}, "2.1.0", "2.1.0")
	compareParsedVerison(t, &Version{Numbers: []int{9, 8, 0}, Separators: []string{".", "-"}, NumberFirst: true}, "9.8-0", "9.8-0")
	compareParsedVerison(t, &Version{Numbers: []int{23, 9}, Separators: []string{"."}, NumberFirst: true}, "23.9", "23.09")
	compareParsedVerison(t, &Version{Numbers: []int{76, 1}, Separators: []string{"a", ".", "d"}}, "a76.1d", "a76.1d")
	compareParsedVerison(t, &Version{Numbers: []int{2, 37, 9}, Separators: []string{"_", "_"}, NumberFirst: true}, "2_37_9", "2_37_9")
}

func compareParsedVerison(t *testing.T, expected *Version, recreated, provided string) {
	t.Helper()
	v := Parse(provided)
	if msg := cmp.Diff(expected, v); len(msg) > 0 {
		t.Error(msg)
	}
	result := v.String(nil)
	if msg := cmp.Diff(recreated, result); len(msg) > 0 {
		t.Error(msg)
	}
}

func TestAddAndResetSmaller(t *testing.T) {
	t.Parallel()
	test := func(t *testing.T, provided, expected []int, index, amount int) {
		t.Helper()
		v := &Version{
			Numbers: provided,
		}
		v.AddAndResetSmaller(index, amount)
		if msg := cmp.Diff(expected, v.Numbers); len(msg) > 0 {
			t.Error(msg)
		}
	}

	test(t, []int{1, 2, 3}, []int{1, 3, 0}, 1, 1)
	test(t, []int{1, 2, 3}, []int{1, 4, 0}, 1, 2)
	test(t, []int{1, 2, 3}, []int{1, 1, 0}, 1, -1)
	test(t, []int{1, 2, 3}, []int{1, 2, 4}, 2, 1)
	test(t, []int{1, 2, 3}, []int{4, 0, 0}, 0, 3)
}
