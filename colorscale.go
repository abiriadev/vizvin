package main

import (
	"slices"
)

type Colorscale interface {
	Scale(value int) float64
}

type LinearScale struct {
	max float64
}

func (this LinearScale) Scale(value int) float64 {
	return float64(value) / this.max
}

type DistinctScale struct {
	valueset []int
}

func dedup[T comparable](s []T) []T {
	if len(s) < 2 {
		return s
	}

	var j int = 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			continue
		}
		s[j] = s[i]
		j++
	}

	return s[:j]
}

func NewDistinctScale(plots [][]int, n int) DistinctScale {
	lin := make([]int, n*n)

	for _, row := range plots {
		lin = append(lin, row...)
	}

	slices.Sort(lin)

	valueset := dedup(lin)

	return DistinctScale{valueset}
}

func (this DistinctScale) Scale(value int) float64 {
	i, _ := slices.BinarySearch(this.valueset, value)
	l := float64(len(this.valueset))
	return float64(i) * l / (l - 1)
}
