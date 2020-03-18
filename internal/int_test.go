package set

import (
	"reflect"
	"sort"
	"testing"
)

func Test_intNewSet(t *testing.T) {
	// Setup
	want := NewSet()
	want.mi = map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
	}

	t.Run("Correct mapping", func(t *testing.T) {
		got := NewSet()
		input := []int{1, 2, 3, 4, 5}
		got.intNewSet(input...)

		expect := reflect.DeepEqual(want.mi, got.mi)

		if !expect {
			t.Errorf("stringNewSet(%q) = %q; want %q", input, got.mi, want.mi)
		}
	})

	t.Run("Incorrect mapping", func(t *testing.T) {
		got := NewSet()
		input := []int{1, 2, 3, 4, 5, 6}
		got.intNewSet(input...)

		expect := reflect.DeepEqual(want.mi, got.mi)

		if expect {
			t.Errorf("intNewSet(%q) = %q; want %q", input, got.mi, want.mi)
		}
	})
}

func Test_getIntSet(t *testing.T) {
	// Setup
	want := []int{1, 2, 3, 4, 5}

	t.Run("Correct slice", func(t *testing.T) {
		input := NewSet()
		input.intNewSet(want...)

		got := input.getIntSet()
		// Sort the got slice to assure the same output every time
		sort.Ints(got)

		expect := reflect.DeepEqual(want, got)

		if !expect {
			t.Errorf("getIntSet() = %q; want %q", got, want)
		}
	})

	t.Run("Incorrect slice", func(t *testing.T) {
		input := NewSet()
		input.intNewSet([]int{1, 2, 3, 4, 5, 6}...)

		got := input.getIntSet()

		expect := reflect.DeepEqual(want, got)

		if expect {
			t.Errorf("getIntSet() = %q; want %q", got, want)
		}
	})
}

func Test_intCardinality(t *testing.T) {
	// Setup
	want := 5

	t.Run("Correct cardinality", func(t *testing.T) {
		input := NewSet()
		input.intNewSet([]int{1, 2, 3, 4, 5}...)

		got := input.intCardinality()

		if got != want {
			t.Errorf("intCardinality() = %d; want %d", got, want)
		}
	})

	t.Run("Incorrect cardinality", func(t *testing.T) {
		input := NewSet()
		input.intNewSet([]int{1, 2, 3, 4, 5, 6}...)

		got := input.intCardinality()

		if got == want {
			t.Errorf("intCardinality() = %d; want %d", got, want)
		}
	})
}
