package set

import "fmt"

type set struct {
	ms  map[string]struct{}
	mi  map[int]struct{}
	mcp map[[2]interface{}]struct{}
}

// NewSet creates a new set from the input
// Returns a pointer to a set
func NewSet(input ...interface{}) *set {
	s := new(set)

	for _, value := range input {
		switch value.(type) {
		case string:
			s.stringNewSet(value.(string))
		case int:
			s.intNewSet(value.(int))
		}
	}

	return s
}

// Add adds new elements to the set s
func (s *set) Add(input ...interface{}) {
	for _, value := range input {
		switch value.(type) {
		case string:
			s.stringNewSet(value.(string))
		case int:
			s.intNewSet(value.(int))
		}
	}
}

// Remove removes elements from the set s
func (s *set) Remove(input ...interface{}) {
	for _, value := range input {
		switch value.(type) {
		case string:
			delete(s.ms, value.(string))
		case int:
			delete(s.mi, value.(int))
		}
	}
}

// Contains verifies if the input elements exists in s
func (s *set) Contains(input interface{}) bool {
	ok := false

	switch input.(type) {
	case string:
		_, ok = s.ms[input.(string)]
	case int:
		_, ok = s.mi[input.(int)]
	}

	return ok
}

// Union unifies two sets, s and input
// Returns a new set
func (s *set) Union(input *set) *set {
	output := NewSet()

	for value := range s.ms {
		output.Add(value)
	}

	for value := range s.mi {
		output.Add(value)
	}

	for value := range input.ms {
		output.Add(value)
	}

	for value := range input.mi {
		output.Add(value)
	}

	return output
}

// Intersection selects the elements that are present in s and input
// Returns a new set
func (s *set) Intersection(input *set) *set {
	output := NewSet()

	for value := range input.ms {
		if s.Contains(value) {
			output.Add(value)
		}
	}

	for value := range input.mi {
		if s.Contains(value) {
			output.Add(value)
		}
	}

	return output
}

// Complement selects the elements presents in input that are not present in s
// Returns a new set
func (s *set) Complement(input *set) *set {
	output := NewSet()

	for value := range input.ms {
		if !s.Contains(value) {
			output.Add(value)
		}
	}

	for value := range input.mi {
		if !s.Contains(value) {
			output.Add(value)
		}
	}

	return output
}

// Cardinality shows the number of elements of a set s
// Returns an int
func (s *set) Cardinality() int {
	return s.stringCardinality() + s.intCardinality() + s.cpCardinality()
}

// Set shows the string representation of the set
// Returns a string
func (s *set) Set() string {
	output := make([]interface{}, 0)

	for _, value := range s.getStringSet() {
		output = append(output, value)
	}

	for _, value := range s.getIntSet() {
		output = append(output, value)
	}

	for _, value := range s.getCartesianProductSet() {
		output = append(output, value)
	}

	str := ""

	for i, value := range output {
		if i == 0 {
			str += fmt.Sprintf("[%+v, ", value)
			continue
		}

		if i == len(output)-1 {
			str += fmt.Sprintf("%+v]", value)
			break
		}

		str += fmt.Sprintf("%+v, ", value)
	}

	return str
}

// CartesianProduct gets the ordered pairs of every element of s, with every
// element of input
// Returns a new set
func (s *set) CartesianProduct(input *set) *set {
	output := NewSet()

	cp := make(map[[2]interface{}]struct{})

	for left := range s.ms {
		for right := range input.ms {
			key := [2]interface{}{
				left,
				right,
			}

			cp[key] = struct{}{}
		}

		for right := range input.mi {
			key := [2]interface{}{
				left,
				right,
			}

			cp[key] = struct{}{}
		}
	}

	for left := range s.mi {
		for right := range input.mi {
			key := [2]interface{}{
				left,
				right,
			}

			cp[key] = struct{}{}
		}

		for right := range input.ms {
			key := [2]interface{}{
				left,
				right,
			}

			cp[key] = struct{}{}
		}
	}

	output.mcp = cp

	return output
}
