package solver

import (
	"slices"
	"testing"
)

func Test_CreateNewString_CreatesAllStrings(t *testing.T) {
	set := []rune("abc")
	wants := []string{
		"a",
		"b",
		"c",
		"aa",
		"ab",
		"ac",
		"ba",
		"bb",
		"bc",
		"ca",
		"cb",
		"cc",
	}

	s := Solver{
		sample: set,
	}

	results := make([]string, 0, len(wants))

	for i := range len(wants) {
		results = append(results, s.CreateUniqueString(i))
	}
	for _, want := range wants {
		if !slices.Contains(results, want) {
			t.Errorf("CreateUniqueString did not create \"%s\" within range %d", want, len(wants))
		}
	}
}
