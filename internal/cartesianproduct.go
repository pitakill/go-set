package set

func (s *set) getCartesianProductSet() [][2]interface{} {
	output := make([][2]interface{}, 0, len(s.mcp))

	for value := range s.mcp {
		output = append(output, value)
	}

	return output
}

func (s *set) cpCardinality() int {
	return len(s.mcp)
}
