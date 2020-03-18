package main

import (
	"fmt"
	"math"

	set "github.com/pitakill/go-set/internal"
)

func main() {
	s0 := set.NewSet("Earth", 2020, math.MaxInt32)
	fmt.Println(s0.Set())

	// Add elements
	s0.Add("Venus", "Mars")
	fmt.Println(s0.Set())

	// Remove elements
	s0.Remove("Venus")
	fmt.Println(s0.Set())

	// Union with different Set
	s1 := set.NewSet("Sun")
	union := s0.Union(s1)
	fmt.Println(union.Set())

	// Intersection with other Set
	s2 := set.NewSet("Mercury", "Venus", "Earth", "Mars")
	intersection := s0.Intersection(s2)
	fmt.Println(intersection.Set())

	// Complement
	complement := s0.Complement(s2)
	fmt.Println(complement.Set())

	// Cartesian Product
	cp := s0.CartesianProduct(s1)
	fmt.Println(cp.Set())

	// Cardinality
	c := cp.Cardinality()
	fmt.Println(c)

	// Contains
	contains := s0.Contains("Mars")
	fmt.Println(contains)
}
