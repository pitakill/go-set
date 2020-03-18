package set

func (s *set) stringNewSet(input ...string) {
	if len(s.ms) == 0 {
		s.ms = make(map[string]struct{})
	}

	for _, value := range input {
		s.ms[value] = struct{}{}
	}
}

func (s *set) getStringSet() []string {
	output := make([]string, 0, len(s.ms))

	for value := range s.ms {
		output = append(output, value)
	}

	return output
}

func (s *set) stringCardinality() int {
	return len(s.ms)
}
