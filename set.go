package set

import "sync"

type Set[T comparable] struct {
	mu               sync.Mutex
	data             map[T]bool
	optFlattenedKeys []T
}

func (s *Set[T]) Add(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.data[item] {
		return
	}
	s.data[item] = true
	s.optFlattenedKeys = nil
}

func (s *Set[T]) Remove(item T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if !s.data[item] {
		return
	}
	delete(s.data, item)
	s.optFlattenedKeys = nil
}

func (s *Set[T]) Has(item T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[item]
}

func (s *Set[T]) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.data)
}

func (s *Set[T]) Flatten() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.optFlattenedKeys != nil {
		return s.optFlattenedKeys
	}
	keys := make([]T, len(s.data))
	keysIdx := 0
	for key, _ := range s.data {
		keys[keysIdx] = key
		keysIdx++
	}
	s.optFlattenedKeys = keys
	return keys
}

func (s *Set[T]) Union(s1 *Set[T]) *Set[T] {
	out := EmptySet[T]()
	for _, e := range s.Flatten() {
		out.Add(e)
	}
	for _, e := range s1.Flatten() {
		out.Add(e)
	}
	return out
}

func (s *Set[T]) Intersect(s1 *Set[T]) *Set[T] {
	out := EmptySet[T]()
	for _, e := range s.Flatten() {
		if s1.Has(e) {
			out.Add(e)
		}
	}
	return out
}

func (s *Set[T]) Copy() *Set[T] {
	dataCp := make(map[T]bool)
	for key, val := range s.data {
		dataCp[key] = val
	}

	return &Set[T]{
		mu:   sync.Mutex{},
		data: dataCp,
	}
}

func EmptySet[T comparable]() *Set[T] {
	return &Set[T]{
		mu:   sync.Mutex{},
		data: make(map[T]bool),
	}
}
