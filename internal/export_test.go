package set

import (
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {
	// Setup
	input1 := []interface{}{"a", "b", "c"}
	input2 := []interface{}{1, 2, 3}

	want1 := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
	}
	want2 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}

	t.Run("String correct", func(t *testing.T) {
		got := NewSet(input1...)

		expect := reflect.DeepEqual(got.ms, want1)

		if !expect {
			t.Errorf("NewSet(%q...) = %q; want %q", input1, got, want1)
		}
	})

	t.Run("Int correct", func(t *testing.T) {
		got := NewSet(input2...)

		expect := reflect.DeepEqual(got.mi, want2)

		if !expect {
			t.Errorf("NewSet(%v...) = %v; want %v", input2, got, want2)
		}
	})

	t.Run("String & int correct", func(t *testing.T) {
		input3 := append(input1, input2...)
		got := NewSet(input3...)

		expect := reflect.DeepEqual(got.ms, want1) && reflect.DeepEqual(got.mi, want2)

		if !expect {
			t.Errorf("NewSet(%v...) = %v; want %v and %v", input3, got, want1, want2)
		}
	})

	t.Run("String incorrect", func(t *testing.T) {
		input1 = append(input1, "d")
		got := NewSet(input1...)

		expect := reflect.DeepEqual(got.ms, want1)

		if expect {
			t.Errorf("NewSet(%q...) = %q; want %q", input1, got, want1)
		}
	})

	t.Run("Int incorrect", func(t *testing.T) {
		input2 = append(input2, 4)
		got := NewSet(input2...)

		expect := reflect.DeepEqual(got.mi, want2)

		if expect {
			t.Errorf("NewSet(%v...) = %v; want %v", input2, got, want2)
		}
	})

	t.Run("String & int correct", func(t *testing.T) {
		input3 := append(input1, input2...)
		got := NewSet(input3...)

		expect := reflect.DeepEqual(got.ms, want1) && reflect.DeepEqual(got.mi, want2)

		if expect {
			t.Errorf("NewSet(%v...) = %v; want %v and %v", input3, got, want1, want2)
		}
	})
}

func TestAdd(t *testing.T) {
	// Setup
	input1 := []interface{}{"a", "b", "c"}
	input2 := []interface{}{1, 2, 3}

	want1 := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
		"d": {},
	}
	want2 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
	}

	t.Run("Add string correct", func(t *testing.T) {
		got := NewSet(input1...)
		got.Add("d", "a")

		expect := reflect.DeepEqual(got.ms, want1)

		if !expect {
			t.Errorf("Add(%q) = %q; want %q", input1, got, want1)
		}
	})

	t.Run("Add int correct", func(t *testing.T) {
		got := NewSet(input2...)
		got.Add(4, 2)

		expect := reflect.DeepEqual(got.mi, want2)

		if !expect {
			t.Errorf("Add(%v) = %v; want %v", input2, got, want2)
		}
	})

	t.Run("Add string & int correct", func(t *testing.T) {
		input3 := append(input1, input2...)
		got := NewSet(input3...)
		got.Add("d", 4, "c", 1)

		expect := reflect.DeepEqual(got.ms, want1) && reflect.DeepEqual(got.mi, want2)

		if !expect {
			t.Errorf("Add(%v) = %v; want %v and %v", input3, got, want1, want2)
		}
	})
}

func TestRemove(t *testing.T) {
	// Setup
	input1 := []interface{}{"a", "b", "c"}
	input2 := []interface{}{1, 2, 3}

	want1 := map[string]struct{}{
		"a": {},
		"b": {},
	}
	want2 := map[int]struct{}{
		1: {},
		2: {},
	}

	t.Run("Remove string correct", func(t *testing.T) {
		got := NewSet(input1...)
		got.Remove("c", "d")

		expect := reflect.DeepEqual(got.ms, want1)

		if !expect {
			t.Errorf("Remove(%q) = %q; want %q", input1, got, want1)
		}
	})

	t.Run("Remove int correct", func(t *testing.T) {
		got := NewSet(input2...)
		got.Remove(4, 3)

		expect := reflect.DeepEqual(got.mi, want2)

		if !expect {
			t.Errorf("Remove(%v) = %v; want %v", input2, got, want2)
		}
	})

	t.Run("Remove string & int correct", func(t *testing.T) {
		input3 := append(input1, input2...)
		got := NewSet(input3...)
		got.Remove("d", 4, "c", 3)

		expect := reflect.DeepEqual(got.ms, want1) && reflect.DeepEqual(got.mi, want2)

		if !expect {
			t.Errorf("Remove(%v) = %v; want %v and %v", input3, got, want1, want2)
		}
	})
}

func TestContains(t *testing.T) {
	// Setup
	input1 := []interface{}{"a", "b", "c"}
	input2 := []interface{}{1, 2, 3}

	t.Run("Contains string", func(t *testing.T) {
		s := NewSet(input1...)
		got := s.Contains("c")

		if !got {
			t.Errorf("Contains(%q) = %v; want %v", input1, got, true)
		}
	})

	t.Run("Contains int", func(t *testing.T) {
		s := NewSet(input2...)
		got := s.Contains(3)

		if !got {
			t.Errorf("Contains(%v) = %v; want %v", input2, got, true)
		}
	})

	t.Run("Does not contains string", func(t *testing.T) {
		s := NewSet(input2...)
		got := s.Contains(4)

		if got {
			t.Errorf("Contains(%v) = %v; want %v", input2, got, false)
		}
	})

	t.Run("Does not contains int", func(t *testing.T) {
		s := NewSet(input2...)
		got := s.Contains(4)

		if got {
			t.Errorf("Contains(%v) = %v; want %v", input2, got, false)
		}
	})
}

func TestUnion(t *testing.T) {
	t.Run("Union string", func(t *testing.T) {
		// Setup
		input1 := []interface{}{"a", "b", "c"}
		input2 := []interface{}{"a", "d", "e", "f"}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := append(input1, input2...)
		want := NewSet(input...)
		s3 := s1.Union(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Union(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("Union int", func(t *testing.T) {
		// Setup
		input1 := []interface{}{1, 2, 3}
		input2 := []interface{}{1, 4, 5, 6}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := append(input1, input2...)
		want := NewSet(input...)
		s3 := s1.Union(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Union(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("Union string & int", func(t *testing.T) {
		// Setup
		input1 := []interface{}{"a", "b", "c", 2}
		input2 := []interface{}{"a", 1, 2, 3}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := append(input1, input2...)
		want := NewSet(input...)
		s3 := s1.Union(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Union(%v) = %v; want %v", input, got, want)
		}
	})
}

func TestIntersection(t *testing.T) {
	t.Run("Intersection string", func(t *testing.T) {
		// Setup
		input1 := []interface{}{"a", "b", "c"}
		input2 := []interface{}{"a", "d", "e", "f"}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := []interface{}{"a"}
		want := NewSet(input...)
		s3 := s1.Intersection(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Intersection(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("Intersection int", func(t *testing.T) {
		// Setup
		input1 := []interface{}{1, 2, 3}
		input2 := []interface{}{1, 4, 5, 6}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := []interface{}{1}
		want := NewSet(input...)
		s3 := s1.Intersection(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Intersection(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("Intersection string & int", func(t *testing.T) {
		// Setup
		input1 := []interface{}{"a", "b", "c", 2}
		input2 := []interface{}{"a", 1, 2, 3}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := []interface{}{"a", 2}
		want := NewSet(input...)
		s3 := s1.Intersection(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Intersection(%v) = %v; want %v", input, got, want)
		}
	})
}

func TestComplement(t *testing.T) {
	t.Run("Complement string", func(t *testing.T) {
		// Setup
		input1 := []interface{}{"a", "b", "c"}
		input2 := []interface{}{"a", "d", "e", "f"}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := []interface{}{"d", "e", "f"}
		want := NewSet(input...)
		s3 := s1.Complement(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Complement(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("Complement int", func(t *testing.T) {
		// Setup
		input1 := []interface{}{1, 2, 3}
		input2 := []interface{}{1, 4, 5, 6}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := []interface{}{4, 5, 6}
		want := NewSet(input...)
		s3 := s1.Complement(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Complement(%v) = %v; want %v", input, got, want)
		}
	})

	t.Run("Complement string & int", func(t *testing.T) {
		// Setup
		input1 := []interface{}{"a", "b", "c", 2}
		input2 := []interface{}{"a", 1, 2, 3}
		s1 := NewSet(input1...)
		s2 := NewSet(input2...)

		input := []interface{}{1, 3}
		want := NewSet(input...)
		s3 := s1.Complement(s2)

		got := reflect.DeepEqual(s3, want)

		if !got {
			t.Errorf("Intersection(%v) = %v; want %v", input, got, want)
		}
	})
}

func TestCardinality(t *testing.T) {
	// Setup
	input1 := []interface{}{"a", "b", "c"}
	input2 := []interface{}{1, 2, 3}
	want := 3

	t.Run("Cardinality string", func(t *testing.T) {
		s := NewSet(input1...)
		got := s.Cardinality()

		if got != want {
			t.Errorf("Cardinality() = %v; want %v", got, want)
		}
	})

	t.Run("Cardinality int", func(t *testing.T) {
		s := NewSet(input2...)
		got := s.Cardinality()

		if got != want {
			t.Errorf("Cardinality() = %v; want %v", got, true)
		}
	})

	t.Run("Cardinality string & int", func(t *testing.T) {
		s := NewSet(input1...).Union(NewSet(input2...))
		got := s.Cardinality()

		if got != 6 {
			t.Errorf("Cardinality() = %v; want %v", got, 6)
		}
	})
}

// There is a problem with the order of the elements, currently the output has
// the same elements but not in the same order
//func TestSet(t *testing.T) {
//// Setup
//input1 := []interface{}{"a", "b", "c"}
//input2 := []interface{}{1, 2, 3}

//t.Run("Set string", func(t *testing.T) {
//s := NewSet(input1...)
//want := "[a, b, c]"

//got := s.Set()

//if got != want {
//t.Errorf("Set() = %v; want %v", got, want)
//}
//})

//t.Run("Set int", func(t *testing.T) {
//s := NewSet(input2...)
//want := "[1, 2, 3]"
//got := s.Set()

//if got != want {
//t.Errorf("Set() = %v; want %v", got, want)
//}
//})

//t.Run("Set string & int", func(t *testing.T) {
//s := NewSet(input1...).CartesianProduct(NewSet(input2...))
//want := "[[a, 1], [a, 2], [a, 3], [b, 1], [b, 2], [b, 3], [c, 1], [c, 2], [c, 3]]"
//got := s.Set()

//if got != want {
//t.Errorf("Set() = %v; want %v", got, want)
//}
//})
//}

func TestCartesianProduct(t *testing.T) {
	// Setup
	input1_1 := []interface{}{"a", "b"}
	input1_2 := []interface{}{"d", "e"}
	input2_1 := []interface{}{1, 2}
	input2_2 := []interface{}{4, 5}

	s1_1 := NewSet(input1_1...)
	s1_2 := NewSet(input1_2...)
	s2_1 := NewSet(input2_1...)
	s2_2 := NewSet(input2_2...)

	want1 := map[[2]interface{}]struct{}{
		{
			"a", "d",
		}: {},
		{
			"a", "e",
		}: {},
		{
			"b", "d",
		}: {},
		{
			"b", "e",
		}: {},
	}

	want2 := map[[2]interface{}]struct{}{
		{
			1, 4,
		}: {},
		{
			1, 5,
		}: {},
		{
			2, 4,
		}: {},
		{
			2, 5,
		}: {},
	}

	t.Run("Cartesian Product string", func(t *testing.T) {
		s := s1_1.CartesianProduct(s1_2)
		got := reflect.DeepEqual(s.mcp, want1)

		if !got {
			t.Errorf("%v.CartesianProduct(%v) = %v; want %v", s1_1, s1_2, s, want1)
		}
	})

	t.Run("Cartesian Product int", func(t *testing.T) {
		s := s2_1.CartesianProduct(s2_2)
		got := reflect.DeepEqual(s.mcp, want2)

		if !got {
			t.Errorf("%v.CartesianProduct(%v) = %v; want %v", s2_1, s2_2, s, want2)
		}
	})

	t.Run("Cartesian Product string & int", func(t *testing.T) {
		s := s1_1.CartesianProduct(s2_1)
		want := map[[2]interface{}]struct{}{
			{
				"a", 1,
			}: {},
			{
				"a", 2,
			}: {},
			{
				"b", 1,
			}: {},
			{
				"b", 2,
			}: {},
		}
		got := reflect.DeepEqual(s.mcp, want)

		if !got {
			t.Errorf("%v.CartesianProduct(%v) = %v; want %v", s2_1, s2_2, s, want2)
		}
	})
}
