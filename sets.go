package immutable

// Set represents a collection of unique values. The set uses a Hasher
// to generate hashes and check for equality of key values.
//
// Internally, the Set stores values as keys of a Map[T,struct{}]
type Set[T any] struct {
	m *Map[T, struct{}]
}

// NewSet returns a new instance of Set.
//
// If hasher is nil, a default hasher implementation will automatically be chosen based on the first key added.
// Default hasher implementations only exist for int, string, and byte slice types.
// NewSet can also take some initial values as varargs.
func NewSet[T any](hasher Hasher[T], values ...T) Set[T] { _ = "STUB: not implemented"; return nil }

// Add returns a set containing the new value.
//
// This function will return a new set even if the set already contains the value.
func (s Set[T]) Add(value T) Set[T] { _ = "STUB: not implemented"; return nil }

// Delete returns a set with the given key removed.
func (s Set[T]) Delete(value T) Set[T] { _ = "STUB: not implemented"; return nil }

// Has returns true when the set contains the given value
func (s Set[T]) Has(val T) bool { _ = "STUB: not implemented"; return false }

// Len returns the number of elements in the underlying map.
func (s Set[K]) Len() int {
	_ = "STUB: not implemented"

	// Items returns a slice of the items inside the set
	return 0
}

func (s Set[T]) Items() []T { _ = "STUB: not implemented"; return nil }

// Iterator returns a new iterator for this set positioned at the first value.
func (s Set[T]) Iterator() *SetIterator[T] { _ = "STUB: not implemented"; return nil }

// SetIterator represents an iterator over a set.
// Iteration can occur in natural or reverse order based on use of Next() or Prev().
type SetIterator[T any] struct {
	mi *MapIterator[T, struct{}]
}

// Done returns true if no more values remain in the iterator.
func (itr *SetIterator[T]) Done() bool { _ = "STUB: not implemented"; return false }

// First moves the iterator to the first value.
func (itr *SetIterator[T]) First() {
	_ = "STUB: not implemented"

	// Next moves the iterator to the next value.
	return
}

func (itr *SetIterator[T]) Next() (val T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

type SetBuilder[T any] struct {
	s Set[T]
}

func NewSetBuilder[T any](hasher Hasher[T]) *SetBuilder[T] { _ = "STUB: not implemented"; return nil }

func (s SetBuilder[T]) Set(val T) { _ = "STUB: not implemented"; return }

func (s SetBuilder[T]) Delete(val T) { _ = "STUB: not implemented"; return }

func (s SetBuilder[T]) Has(val T) bool { _ = "STUB: not implemented"; return false }

func (s SetBuilder[T]) Len() int { _ = "STUB: not implemented"; return 0 }

type SortedSet[T any] struct {
	m *SortedMap[T, struct{}]
}

// NewSortedSet returns a new instance of SortedSet.
//
// If comparer is nil then
// a default comparer is set after the first key is inserted. Default comparers
// exist for int, string, and byte slice keys.
// NewSortedSet can also take some initial values as varargs.
func NewSortedSet[T any](comparer Comparer[T], values ...T) SortedSet[T] {
	_ = "STUB: not implemented"
	return nil
}

// Add returns a set containing the new value.
//
// This function will return a new set even if the set already contains the value.
func (s SortedSet[T]) Add(value T) SortedSet[T] { _ = "STUB: not implemented"; return nil }

// Delete returns a set with the given key removed.
func (s SortedSet[T]) Delete(value T) SortedSet[T] { _ = "STUB: not implemented"; return nil }

// Has returns true when the set contains the given value
func (s SortedSet[T]) Has(val T) bool { _ = "STUB: not implemented"; return false }

// Len returns the number of elements in the underlying map.
func (s SortedSet[K]) Len() int {
	_ = "STUB: not implemented"

	// Items returns a slice of the items inside the set
	return 0
}

func (s SortedSet[T]) Items() []T { _ = "STUB: not implemented"; return nil }

// Iterator returns a new iterator for this set positioned at the first value.
func (s SortedSet[T]) Iterator() *SortedSetIterator[T] { _ = "STUB: not implemented"; return nil }

// SortedSetIterator represents an iterator over a sorted set.
// Iteration can occur in natural or reverse order based on use of Next() or Prev().
type SortedSetIterator[T any] struct {
	mi *SortedMapIterator[T, struct{}]
}

// Done returns true if no more values remain in the iterator.
func (itr *SortedSetIterator[T]) Done() bool { _ = "STUB: not implemented"; return false }

// First moves the iterator to the first value.
func (itr *SortedSetIterator[T]) First() {
	_ = "STUB: not implemented"

	// Last moves the iterator to the last value.
	return
}

func (itr *SortedSetIterator[T]) Last() {
	_ = "STUB: not implemented"

	// Next moves the iterator to the next value.
	return
}

func (itr *SortedSetIterator[T]) Next() (val T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Prev moves the iterator to the previous value.
func (itr *SortedSetIterator[T]) Prev() (val T, ok bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// Seek moves the iterator to the given value.
//
// If the value does not exist then the next value is used. If no more keys exist
// then the iterator is marked as done.
func (itr *SortedSetIterator[T]) Seek(val T) { _ = "STUB: not implemented"; return }

type SortedSetBuilder[T any] struct {
	s *SortedSet[T]
}

func NewSortedSetBuilder[T any](comparer Comparer[T]) *SortedSetBuilder[T] {
	_ = "STUB: not implemented"
	return nil
}

func (s SortedSetBuilder[T]) Set(val T) { _ = "STUB: not implemented"; return }

func (s SortedSetBuilder[T]) Delete(val T) { _ = "STUB: not implemented"; return }

func (s SortedSetBuilder[T]) Has(val T) bool { _ = "STUB: not implemented"; return false }

func (s SortedSetBuilder[T]) Len() int {
	_ = "STUB: not implemented"

	// SortedSet returns the current copy of the set.
	// The builder should not be used again after the list after this call.
	return 0
}

func (s SortedSetBuilder[T]) SortedSet() SortedSet[T] { _ = "STUB: not implemented"; return nil }
