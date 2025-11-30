package utils

// Set is a simple set implementation using a map
type Set[T comparable] map[T]struct{}

// NewSet creates a new empty set
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// NewSetFrom creates a new set from a slice
func NewSetFrom[T comparable](items []T) Set[T] {
	s := NewSet[T]()
	for _, item := range items {
		s.Add(item)
	}
	return s
}

// Add adds an item to the set
func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

// Remove removes an item from the set
func (s Set[T]) Remove(item T) {
	delete(s, item)
}

// Contains checks if an item is in the set
func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Size returns the number of items in the set
func (s Set[T]) Size() int {
	return len(s)
}

// ToSlice converts the set to a slice
func (s Set[T]) ToSlice() []T {
	result := make([]T, 0, len(s))
	for item := range s {
		result = append(result, item)
	}
	return result
}

// Queue is a simple FIFO queue
type Queue[T any] struct {
	items []T
}

// NewQueue creates a new empty queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{items: make([]T, 0)}
}

// Enqueue adds an item to the queue
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the first item
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

// IsEmpty returns true if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Stack is a simple LIFO stack
type Stack[T any] struct {
	items []T
}

// NewStack creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

// Peek returns the top item without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Keys returns all keys from a map
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values from a map
func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Reverse reverses a slice in place
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Contains checks if a slice contains an item
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// IndexOf returns the index of an item in a slice, or -1 if not found
func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

// Filter returns a new slice containing only items that satisfy the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Map applies a function to each element of a slice
func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}
