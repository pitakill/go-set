package set

import (
	"reflect"
	"sort"
	"testing"
)

func Test_stringNewSet(t *testing.T) {
	// Setup
	want := NewSet()
	want.ms = map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
		"d": {},
		"e": {},
	}

	t.Run("Correct mapping", func(t *testing.T) {
		got := NewSet()
		input := []string{"a", "b", "c", "d", "e"}
		got.stringNewSet(input...)

		expect := reflect.DeepEqual(want.ms, got.ms)

		if !expect {
			t.Errorf("stringNewSet(%q) = %q; want %q", input, got.ms, want.ms)
		}
	})

	t.Run("Incorrect mapping", func(t *testing.T) {
		got := NewSet()
		input := []string{"a", "b", "c", "d", "e", "f"}
		got.stringNewSet(input...)

		expect := reflect.DeepEqual(want.ms, got.ms)

		if expect {
			t.Errorf("stringNewSet(%q) = %q; want %q", input, got.ms, want.ms)
		}
	})
}

func Test_getStringSet(t *testing.T) {
	// Setup
	want := []string{"a", "b", "c", "d", "e"}

	t.Run("Correct slice", func(t *testing.T) {
		input := NewSet()
		input.stringNewSet(want...)

		got := input.getStringSet()
		// Sort the got slice to assure the same output every time
		sort.Strings(got)

		expect := reflect.DeepEqual(want, got)

		if !expect {
			t.Errorf("getStringSet() = %q; want %q", got, want)
		}
	})

	t.Run("Incorrect slice", func(t *testing.T) {
		input := NewSet()
		input.stringNewSet([]string{"a", "b", "c", "d", "e", "f"}...)

		got := input.getStringSet()

		expect := reflect.DeepEqual(want, got)

		if expect {
			t.Errorf("getStringSet() = %q; want %q", got, want)
		}
	})
}

func Test_stringCardinality(t *testing.T) {
	// Setup
	want := 5

	t.Run("Correct cardinality", func(t *testing.T) {
		input := NewSet()
		input.stringNewSet([]string{"a", "b", "c", "d", "e"}...)

		got := input.stringCardinality()

		if got != want {
			t.Errorf("stringCardinality() = %d; want %d", got, want)
		}
	})

	t.Run("Incorrect cardinality", func(t *testing.T) {
		input := NewSet()
		input.stringNewSet([]string{"a", "b", "c", "d", "e", "f"}...)

		got := input.stringCardinality()

		if got == want {
			t.Errorf("stringCardinality() = %d; want %d", got, want)
		}
	})
}
