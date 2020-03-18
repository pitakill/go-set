package set

import (
	"reflect"
	"testing"
)

func Test_getCartesianProductSet(t *testing.T) {
	// Setup
	s1 := NewSet()
	s1.stringNewSet([]string{"a", "b"}...)

	t.Run("Correct", func(t *testing.T) {
		s2 := NewSet()
		s2.intNewSet([]int{1, 2, 3}...)

		s3 := s1.CartesianProduct(s2)

		want := [][2]interface{}{
			{"a", 1},
			{"a", 2},
			{"a", 3},
			{"b", 1},
			{"b", 2},
			{"b", 3},
		}

		got := s3.getCartesianProductSet()

		expect := true
		for _, e := range got {
			if !find(got, e) {
				expect = false
				break
			}
		}

		if !expect {
			t.Errorf("getCartesianProductSet() = %v; want %v", got, want)
		}
	})

	t.Run("Incorrect", func(t *testing.T) {
		s2 := NewSet()
		s2.intNewSet([]int{1, 2, 3, 4}...)

		s3 := s1.CartesianProduct(s2)

		want := [][2]interface{}{
			{"a", 1},
			{"a", 2},
			{"a", 3},
			{"b", 1},
			{"b", 2},
			{"b", 3},
		}

		got := s3.getCartesianProductSet()

		expect := false
		for _, e := range got {
			if find(got, e) {
				expect = true
				break
			}
		}

		if !expect {
			t.Errorf("getCartesianProductSet() = %v; want %v", got, want)
		}
	})
}

func find(slice [][2]interface{}, v [2]interface{}) bool {
	for _, item := range slice {
		if reflect.DeepEqual(item, v) {
			return true
		}
	}

	return false
}

func Test_cpCardinality(t *testing.T) {
	// Setup
	want := 10
	s1 := NewSet()
	s1.stringNewSet([]string{"a", "b"}...)

	t.Run("Correct cardinality", func(t *testing.T) {
		s2 := NewSet()
		s2.intNewSet([]int{1, 2, 3, 4, 5}...)

		s3 := s1.CartesianProduct(s2)

		got := s3.cpCardinality()

		if got != want {
			t.Errorf("cpCardinality() = %d; want %d", got, want)
		}
	})

	t.Run("Incorrect cardinality", func(t *testing.T) {
		s2 := NewSet()
		s2.intNewSet([]int{1, 2, 3, 4, 5, 6}...)

		s3 := s1.CartesianProduct(s2)

		got := s3.cpCardinality()

		if got == want {
			t.Errorf("intCardinality() = %d; want %d", got, want)
		}
	})
}
