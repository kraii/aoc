package aoc

type Set[T comparable] map[T]bool

func (s Set[T]) Add(item T) {
	s[item] = true
}

func (s Set[T]) Contains(item T) bool {
	return s[item]
}
