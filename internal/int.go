package set

func (s *set) intNewSet(input ...int) {
	if len(s.mi) == 0 {
		s.mi = make(map[int]struct{})
	}

	for _, value := range input {
		s.mi[value] = struct{}{}
	}
}

func (s *set) getIntSet() []int {
	output := make([]int, 0, len(s.mi))

	for value := range s.mi {
		output = append(output, value)
	}

	return output
}

func (s *set) intCardinality() int {
	return len(s.mi)
}
