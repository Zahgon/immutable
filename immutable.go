// Package immutable provides immutable collection types.
//
// # Introduction
//
// Immutable collections provide an efficient, safe way to share collections
// of data while minimizing locks. The collections in this package provide
// List, Map, and SortedMap implementations. These act similarly to slices
// and maps, respectively, except that altering a collection returns a new
// copy of the collection with that change.
//
// Because collections are unable to change, they are safe for multiple
// goroutines to read from at the same time without a mutex. However, these
// types of collections come with increased CPU & memory usage as compared
// with Go's built-in collection types so please evaluate for your specific
// use.
//
// # Collection Types
//
// The List type provides an API similar to Go slices. They allow appending,
// prepending, and updating of elements. Elements can also be fetched by index
// or iterated over using a ListIterator.
//
// The Map & SortedMap types provide an API similar to Go maps. They allow
// values to be assigned to unique keys and allow for the deletion of keys.
// Values can be fetched by key and key/value pairs can be iterated over using
// the appropriate iterator type. Both map types provide the same API. The
// SortedMap, however, provides iteration over sorted keys while the Map
// provides iteration over unsorted keys. Maps improved performance and memory
// usage as compared to SortedMaps.
//
// # Hashing and Sorting
//
// Map types require the use of a Hasher implementation to calculate hashes for
// their keys and check for key equality. SortedMaps require the use of a
// Comparer implementation to sort keys in the map.
//
// These collection types automatically provide built-in hasher and comparers
// for int, string, and byte slice keys. If you are using one of these key types
// then simply pass a nil into the constructor. Otherwise you will need to
// implement a custom Hasher or Comparer type. Please see the provided
// implementations for reference.
package immutable

import (
	"golang.org/x/exp/constraints"
)

// List is a dense, ordered, indexed collections. They are analogous to slices
// in Go. They can be updated by appending to the end of the list, prepending
// values to the beginning of the list, or updating existing indexes in the
// list.
type List[T any] struct {
	root   listNode[T] // root node
	origin int         // offset to zero index element
	size   int         // total number of elements in use
}

// NewList returns a new empty instance of List.
func NewList[T any](values ...T) *List[T] { _ = "STUB: not implemented"; return nil }

// clone returns a copy of the list.
func (l *List[T]) clone() *List[T] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the list.
func (l *List[T]) Len() int {
	_ = "STUB: not implemented"

	// cap returns the total number of possible elements for the current depth.
	return 0
}

func (l *List[T]) cap() int { _ = "STUB: not implemented"; return 0 }

// Get returns the value at the given index. Similar to slices, this method will
// panic if index is below zero or is greater than or equal to the list size.
func (l *List[T]) Get(index int) T { _ = "STUB: not implemented"; return *new(T) }

// Set returns a new list with value set at index. Similar to slices, this
// method will panic if index is below zero or if the index is greater than
// or equal to the list size.
func (l *List[T]) Set(index int, value T) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) set(index int, value T, mutable bool) *List[T] {
	_ = "STUB: not implemented"
	return nil
}

// Append returns a new list with value added to the end of the list.
func (l *List[T]) Append(value T) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) append(value T, mutable bool) *List[T] { _ = "STUB: not implemented"; return nil }

// Expand list to the right if no slots remain.

// Increase size and set the last element to the new value.

// Prepend returns a new list with value(s) added to the beginning of the list.
func (l *List[T]) Prepend(value T) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) prepend(value T, mutable bool) *List[T] { _ = "STUB: not implemented"; return nil }

// Expand list to the left if no slots remain.

// Increase size and move origin back. Update first element to value.

// Slice returns a new list of elements between start index and end index.
// Similar to slices, this method will panic if start or end are below zero or
// greater than the list size. A panic will also occur if start is greater than
// end.
//
// Unlike Go slices, references to inaccessible elements will be automatically
// removed so they can be garbage collected.
func (l *List[T]) Slice(start, end int) *List[T] { _ = "STUB: not implemented"; return nil }

func (l *List[T]) slice(start, end int, mutable bool) *List[T] {
	_ = "STUB: not implemented"
	// Panics similar to Go slices.
	return nil
}

// Return the same list if the start and end are the entire range.

// Create copy, if immutable.

// Update origin/size.

// Contract tree while the start & end are in the same child node.

// branch contains at least two nodes, exit

// Replace the current root with the single child & update origin offset.

// Ensure all references are removed before start & after end.

// Iterator returns a new iterator for this list positioned at the first index.
func (l *List[T]) Iterator() *ListIterator[T] { _ = "STUB: not implemented"; return nil }

// ListBuilder represents an efficient builder for creating new Lists.
type ListBuilder[T any] struct {
	list *List[T] // current state
}

// NewListBuilder returns a new instance of ListBuilder.
func NewListBuilder[T any]() *ListBuilder[T] { _ = "STUB: not implemented"; return nil }

// List returns the current copy of the list.
// The builder should not be used again after the list after this call.
func (b *ListBuilder[T]) List() *List[T] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the underlying list.
func (b *ListBuilder[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// Get returns the value at the given index. Similar to slices, this method will
// panic if index is below zero or is greater than or equal to the list size.
func (b *ListBuilder[T]) Get(index int) T { _ = "STUB: not implemented"; return *new(T) }

// Set updates the value at the given index. Similar to slices, this method will
// panic if index is below zero or if the index is greater than or equal to the
// list size.
func (b *ListBuilder[T]) Set(index int, value T) { _ = "STUB: not implemented"; return }

// Append adds value to the end of the list.
func (b *ListBuilder[T]) Append(value T) { _ = "STUB: not implemented"; return }

// Prepend adds value to the beginning of the list.
func (b *ListBuilder[T]) Prepend(value T) { _ = "STUB: not implemented"; return }

// Slice updates the list with a sublist of elements between start and end index.
// See List.Slice() for more details.
func (b *ListBuilder[T]) Slice(start, end int) { _ = "STUB: not implemented"; return }

// Iterator returns a new iterator for the underlying list.
func (b *ListBuilder[T]) Iterator() *ListIterator[T] { _ = "STUB: not implemented"; return nil }

// Constants for bit shifts used for levels in the List trie.
const (
	listNodeBits = 5
	listNodeSize = 1 << listNodeBits
	listNodeMask = listNodeSize - 1
)

// listNode represents either a branch or leaf node in a List.
type listNode[T any] interface {
	depth() uint
	get(index int) T
	set(index int, v T, mutable bool) listNode[T]

	containsBefore(index int) bool
	containsAfter(index int) bool

	deleteBefore(index int, mutable bool) listNode[T]
	deleteAfter(index int, mutable bool) listNode[T]
}

// newListNode returns a leaf node for depth zero, otherwise returns a branch node.
func newListNode[T any](depth uint) listNode[T] { _ = "STUB: not implemented"; return nil }

// listBranchNode represents a branch of a List tree at a given depth.
type listBranchNode[T any] struct {
	d        uint // depth
	children [listNodeSize]listNode[T]
}

// depth returns the depth of this branch node from the leaf.
func (n *listBranchNode[T]) depth() uint {
	_ = "STUB: not implemented"

	// get returns the child node at the segment of the index for this depth.
	return 0
}

func (n *listBranchNode[T]) get(index int) T { _ = "STUB: not implemented"; return *new(T) }

// set recursively updates the value at index for each lower depth from the node.
func (n *listBranchNode[T]) set(index int, v T, mutable bool) listNode[T] {
	_ = "STUB: not implemented"
	return nil
}

// Find child for the given value in the branch. Create new if it doesn't exist.

// Return a copy of this branch with the new child.

// containsBefore returns true if non-nil values exists between [0,index).
func (n *listBranchNode[T]) containsBefore(index int) bool { _ = "STUB: not implemented"; return false }

// Quickly check if any direct children exist before this segment of the index.

// Recursively check for children directly at the given index at this segment.

// containsAfter returns true if non-nil values exists between (index,listNodeSize).
func (n *listBranchNode[T]) containsAfter(index int) bool { _ = "STUB: not implemented"; return false }

// Quickly check if any direct children exist after this segment of the index.

// Recursively check for children directly at the given index at this segment.

// deleteBefore returns a new node with all elements before index removed.
func (n *listBranchNode[T]) deleteBefore(index int, mutable bool) listNode[T] {
	_ = "STUB: not implemented"
	// Ignore if no nodes exist before the given index.
	return nil
}

// Return a copy with any nodes prior to the index removed.

// deleteBefore returns a new node with all elements before index removed.
func (n *listBranchNode[T]) deleteAfter(index int, mutable bool) listNode[T] {
	_ = "STUB: not implemented"
	// Ignore if no nodes exist after the given index.
	return nil
}

// Return a copy with any nodes after the index removed.

// listLeafNode represents a leaf node in a List.
type listLeafNode[T any] struct {
	children [listNodeSize]T
	// bitset with ones at occupied positions, position 0 is the LSB
	occupied uint32
}

// depth always returns 0 for leaf nodes.
func (n *listLeafNode[T]) depth() uint {
	_ = "STUB: not implemented"

	// get returns the value at the given index.
	return 0
}

func (n *listLeafNode[T]) get(index int) T { _ = "STUB: not implemented"; return *new(T) }

// set returns a copy of the node with the value at the index updated to v.
func (n *listLeafNode[T]) set(index int, v T, mutable bool) listNode[T] {
	_ = "STUB: not implemented"
	return nil
}

// containsBefore returns true if non-nil values exists between [0,index).
func (n *listLeafNode[T]) containsBefore(index int) bool { _ = "STUB: not implemented"; return false }

// containsAfter returns true if non-nil values exists between (index,listNodeSize).
func (n *listLeafNode[T]) containsAfter(index int) bool { _ = "STUB: not implemented"; return false }

// deleteBefore returns a new node with all elements before index removed.
func (n *listLeafNode[T]) deleteBefore(index int, mutable bool) listNode[T] {
	_ = "STUB: not implemented"
	return nil
}

// Set the first idx bits to 0.

// deleteAfter returns a new node with all elements after index removed.
func (n *listLeafNode[T]) deleteAfter(index int, mutable bool) listNode[T] {
	_ = "STUB: not implemented"
	return nil
}

// Set bits after idx to 0. idx < 31 because n.containsAfter(index) == true.

// ListIterator represents an ordered iterator over a list.
type ListIterator[T any] struct {
	list  *List[T] // source list
	index int      // current index position

	stack [32]listIteratorElem[T] // search stack
	depth int                     // stack depth
}

// Done returns true if no more elements remain in the iterator.
func (itr *ListIterator[T]) Done() bool { _ = "STUB: not implemented"; return false }

// First positions the iterator on the first index.
// If source list is empty then no change is made.
func (itr *ListIterator[T]) First() { _ = "STUB: not implemented"; return }

// Last positions the iterator on the last index.
// If source list is empty then no change is made.
func (itr *ListIterator[T]) Last() { _ = "STUB: not implemented"; return }

// Seek moves the iterator position to the given index in the list.
// Similar to Go slices, this method will panic if index is below zero or if
// the index is greater than or equal to the list size.
func (itr *ListIterator[T]) Seek(index int) {
	_ = "STUB: not implemented"
	// Panic similar to Go slices.
	return
}

// Reset to the bottom of the stack at seek to the correct position.

// Next returns the current index and its value & moves the iterator forward.
// Returns an index of -1 if the there are no more elements to return.
func (itr *ListIterator[T]) Next() (index int, value T) {
	_ = "STUB: not implemented"
	// Exit immediately if there are no elements remaining.
	return 0, *new(T)
}

// Retrieve current index & value.

// Increase index. If index is at the end then return immediately.

// Move up stack until we find a node that has remaining position ahead.

// Seek to correct position from current depth.

// Prev returns the current index and value and moves the iterator backward.
// Returns an index of -1 if the there are no more elements to return.
func (itr *ListIterator[T]) Prev() (index int, value T) {
	_ = "STUB: not implemented"
	// Exit immediately if there are no elements remaining.
	return 0, *new(T)
}

// Retrieve current index & value.

// Decrease index. If index is past the beginning then return immediately.

// Move up stack until we find a node that has remaining position behind.

// Seek to correct position from current depth.

// seek positions the stack to the given index from the current depth.
// Elements and indexes below the current depth are assumed to be correct.
func (itr *ListIterator[T]) seek(index int) {
	_ = "STUB: not implemented"
	// Iterate over each level until we reach a leaf node.
	return
}

// listIteratorElem represents the node and it's child index within the stack.
type listIteratorElem[T any] struct {
	node  listNode[T]
	index int
}

// Size thresholds for each type of branch node.
const (
	maxArrayMapSize      = 8
	maxBitmapIndexedSize = 16
)

// Segment bit shifts within the map tree.
const (
	mapNodeBits = 5
	mapNodeSize = 1 << mapNodeBits
	mapNodeMask = mapNodeSize - 1
)

// Map represents an immutable hash map implementation. The map uses a Hasher
// to generate hashes and check for equality of key values.
//
// It is implemented as an Hash Array Mapped Trie.
type Map[K, V any] struct {
	size   int           // total number of key/value pairs
	root   mapNode[K, V] // root node of trie
	hasher Hasher[K]     // hasher implementation
}

// NewMap returns a new instance of Map. If hasher is nil, a default hasher
// implementation will automatically be chosen based on the first key added.
// Default hasher implementations only exist for int, string, and byte slice types.
func NewMap[K, V any](hasher Hasher[K]) *Map[K, V] { _ = "STUB: not implemented"; return nil }

// NewMapOf returns a new instance of Map, containing a map of provided entries.
//
// If hasher is nil, a default hasher implementation will automatically be chosen based on the first key added.
// Default hasher implementations only exist for int, string, and byte slice types.
func NewMapOf[K comparable, V any](hasher Hasher[K], entries map[K]V) *Map[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Len returns the number of elements in the map.
func (m *Map[K, V]) Len() int {
	_ = "STUB: not implemented"

	// clone returns a shallow copy of m.
	return 0
}

func (m *Map[K, V]) clone() *Map[K, V] { _ = "STUB: not implemented"; return nil }

// Get returns the value for a given key and a flag indicating whether the
// key exists. This flag distinguishes a nil value set on a key versus a
// non-existent key in the map.
func (m *Map[K, V]) Get(key K) (value V, ok bool) { _ = "STUB: not implemented"; return *new(V), false }

// Set returns a map with the key set to the new value. A nil value is allowed.
//
// This function will return a new map even if the updated value is the same as
// the existing value because Map does not track value equality.
func (m *Map[K, V]) Set(key K, value V) *Map[K, V] { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) set(key K, value V, mutable bool) *Map[K, V] {
	_ = "STUB: not implemented"
	// Set a hasher on the first value if one does not already exist.
	return nil
}

// Generate copy if necessary.

// If the map is empty, initialize with a simple array node.

// Otherwise copy the map and delegate insertion to the root.
// Resized will return true if the key does not currently exist.

// Delete returns a map with the given key removed.
// Removing a non-existent key will cause this method to return the same map.
func (m *Map[K, V]) Delete(key K) *Map[K, V] { _ = "STUB: not implemented"; return nil }

func (m *Map[K, V]) delete(key K, mutable bool) *Map[K, V] {
	_ = "STUB: not implemented"
	// Return original map if no keys exist.
	return nil
}

// If the delete did not change the node then return the original map.

// Generate copy if necessary.

// Return copy of map with new root and decreased size.

// Iterator returns a new iterator for the map.
func (m *Map[K, V]) Iterator() *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// MapBuilder represents an efficient builder for creating Maps.
type MapBuilder[K, V any] struct {
	m *Map[K, V] // current state
}

// NewMapBuilder returns a new instance of MapBuilder.
func NewMapBuilder[K, V any](hasher Hasher[K]) *MapBuilder[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Map returns the underlying map. Only call once.
// Builder is invalid after call. Will panic on second invocation.
func (b *MapBuilder[K, V]) Map() *Map[K, V] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the underlying map.
func (b *MapBuilder[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

// Get returns the value for the given key.
func (b *MapBuilder[K, V]) Get(key K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Set sets the value of the given key. See Map.Set() for additional details.
func (b *MapBuilder[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

// Delete removes the given key. See Map.Delete() for additional details.
func (b *MapBuilder[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

// Iterator returns a new iterator for the underlying map.
func (b *MapBuilder[K, V]) Iterator() *MapIterator[K, V] { _ = "STUB: not implemented"; return nil }

// mapNode represents any node in the map tree.
type mapNode[K, V any] interface {
	get(key K, shift uint, keyHash uint32, h Hasher[K]) (value V, ok bool)
	set(key K, value V, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V]
	delete(key K, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V]
}

var _ mapNode[string, any] = (*mapArrayNode[string, any])(nil)
var _ mapNode[string, any] = (*mapBitmapIndexedNode[string, any])(nil)
var _ mapNode[string, any] = (*mapHashArrayNode[string, any])(nil)
var _ mapNode[string, any] = (*mapValueNode[string, any])(nil)
var _ mapNode[string, any] = (*mapHashCollisionNode[string, any])(nil)

// mapLeafNode represents a node that stores a single key hash at the leaf of the map tree.
type mapLeafNode[K, V any] interface {
	mapNode[K, V]
	keyHashValue() uint32
}

var _ mapLeafNode[string, any] = (*mapValueNode[string, any])(nil)
var _ mapLeafNode[string, any] = (*mapHashCollisionNode[string, any])(nil)

// mapArrayNode is a map node that stores key/value pairs in a slice.
// Entries are stored in insertion order. An array node expands into a bitmap
// indexed node once a given threshold size is crossed.
type mapArrayNode[K, V any] struct {
	entries []mapEntry[K, V]
}

// indexOf returns the entry index of the given key. Returns -1 if key not found.
func (n *mapArrayNode[K, V]) indexOf(key K, h Hasher[K]) int { _ = "STUB: not implemented"; return 0 }

// get returns the value for the given key.
func (n *mapArrayNode[K, V]) get(key K, shift uint, keyHash uint32, h Hasher[K]) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// set inserts or updates the value for a given key. If the key is inserted and
// the new size crosses the max size threshold, a bitmap indexed node is returned.
func (n *mapArrayNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	return nil

	// Mark as resized if the key doesn't exist.
}

// If we are adding and it crosses the max size threshold, expand the node.
// We do this by continually setting the entries to a value node and expanding.

// Update in-place if mutable.

// Update existing entry if a match is found.
// Otherwise append to the end of the element list if it doesn't exist.

// delete removes the given key from the node. Returns the same node if key does
// not exist. Returns a nil node when removing the last entry.
func (n *mapArrayNode[K, V]) delete(key K, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	return nil

	// Return original node if key does not exist.
}

// Return nil if this node will contain no nodes.

// Update in-place, if mutable.

// Otherwise create a copy with the given entry removed.

// mapBitmapIndexedNode represents a map branch node with a variable number of
// node slots and indexed using a bitmap. Indexes for the node slots are
// calculated by counting the number of set bits before the target bit using popcount.
type mapBitmapIndexedNode[K, V any] struct {
	bitmap uint32
	nodes  []mapNode[K, V]
}

// get returns the value for the given key.
func (n *mapBitmapIndexedNode[K, V]) get(key K, shift uint, keyHash uint32, h Hasher[K]) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// set inserts or updates the value for the given key. If a new key is inserted
// and the size crosses the max size threshold then a hash array node is returned.
func (n *mapBitmapIndexedNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	// Extract the index for the bit segment of the key hash.
	return nil
}

// Determine the bit based on the hash index.

// Mark as resized if the key doesn't exist.

// Find index of node based on popcount of bits before it.

// If the node already exists, delegate set operation to it.
// If the node doesn't exist then create a simple value leaf node.

// Convert to a hash-array node once we exceed the max bitmap size.
// Copy each node based on their bit position within the bitmap.

// Update in-place if mutable.

// If node exists at given slot then overwrite it with new node.
// Otherwise expand the node list and insert new node into appropriate position.

// delete removes the key from the tree. If the key does not exist then the
// original node is returned. If removing the last child node then a nil is
// returned. Note that shrinking the node will not convert it to an array node.
func (n *mapBitmapIndexedNode[K, V]) delete(key K, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Return original node if key does not exist.

// Find index of node based on popcount of bits before it.

// Delegate delete to child node.

// Return original node if key doesn't exist in child.

// Remove if returned child has been deleted.

// If we won't have any children then return nil.

// Update in-place if mutable.

// Return copy with bit removed from bitmap and node removed from node list.

// Generate copy, if necessary.

// Update child.

// mapHashArrayNode is a map branch node that stores nodes in a fixed length
// array. Child nodes are indexed by their index bit segment for the current depth.
type mapHashArrayNode[K, V any] struct {
	count uint                       // number of set nodes
	nodes [mapNodeSize]mapNode[K, V] // child node slots, may contain empties
}

// clone returns a shallow copy of n.
func (n *mapHashArrayNode[K, V]) clone() *mapHashArrayNode[K, V] {
	_ = "STUB: not implemented"
	return nil

	// get returns the value for the given key.
}

func (n *mapHashArrayNode[K, V]) get(key K, shift uint, keyHash uint32, h Hasher[K]) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// set returns a node with the value set for the given key.
func (n *mapHashArrayNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// If node at index doesn't exist, create a simple value leaf node.
// Otherwise delegate set to child node.

// Generate copy, if necessary.

// Update child node (and update size, if new).

// delete returns a node with the given key removed. Returns the same node if
// the key does not exist. If node shrinks to within bitmap-indexed size then
// converts to a bitmap-indexed node.
func (n *mapHashArrayNode[K, V]) delete(key K, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Return original node if child is not found.

// Return original node if child is unchanged.

// If we remove a node and drop below a threshold, convert back to bitmap indexed node.

// Generate copy, if necessary.

// Return copy of node with child updated.

// mapValueNode represents a leaf node with a single key/value pair.
// A value node can be converted to a hash collision leaf node if a different
// key with the same keyHash is inserted.
type mapValueNode[K, V any] struct {
	keyHash uint32
	key     K
	value   V
}

// newMapValueNode returns a new instance of mapValueNode.
func newMapValueNode[K, V any](keyHash uint32, key K, value V) *mapValueNode[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// keyHashValue returns the key hash for this node.
func (n *mapValueNode[K, V]) keyHashValue() uint32 {
	_ = "STUB: not implemented"

	// get returns the value for the given key.
	return 0
}

func (n *mapValueNode[K, V]) get(key K, shift uint, keyHash uint32, h Hasher[K]) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// set returns a new node with the new value set for the key. If the key equals
// the node's key then a new value node is returned. If key is not equal to the
// node's key but has the same hash then a hash collision node is returned.
// Otherwise the nodes are merged into a branch node.
func (n *mapValueNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	// If the keys match then return a new value node overwriting the value.
	return nil
}

// Update in-place if mutable.

// Otherwise return a new copy.

// Recursively merge nodes together if key hashes are different.

// Merge into collision node if hash matches.

// delete returns nil if the key matches the node's key. Otherwise returns the original node.
func (n *mapValueNode[K, V]) delete(key K, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	// Return original node if the keys do not match.
	return nil
}

// Otherwise remove the node if keys do match.

// mapHashCollisionNode represents a leaf node that contains two or more key/value
// pairs with the same key hash. Single pairs for a hash are stored as value nodes.
type mapHashCollisionNode[K, V any] struct {
	keyHash uint32 // key hash for all entries
	entries []mapEntry[K, V]
}

// keyHashValue returns the key hash for all entries on the node.
func (n *mapHashCollisionNode[K, V]) keyHashValue() uint32 {
	_ = "STUB: not implemented"

	// indexOf returns the index of the entry for the given key.
	// Returns -1 if the key does not exist in the node.
	return 0
}

func (n *mapHashCollisionNode[K, V]) indexOf(key K, h Hasher[K]) int {
	_ = "STUB: not implemented"
	return 0
}

// get returns the value for the given key.
func (n *mapHashCollisionNode[K, V]) get(key K, shift uint, keyHash uint32, h Hasher[K]) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// set returns a copy of the node with key set to the given value.
func (n *mapHashCollisionNode[K, V]) set(key K, value V, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	// Merge node with key/value pair if this is not a hash collision.
	return nil
}

// Update in-place if mutable.

// Append to end of node if key doesn't exist & mark resized.
// Otherwise copy nodes and overwrite at matching key index.

// delete returns a node with the given key deleted. Returns the same node if
// the key does not exist. If removing the key would shrink the node to a single
// entry then a value node is returned.
func (n *mapHashCollisionNode[K, V]) delete(key K, shift uint, keyHash uint32, h Hasher[K], mutable bool, resized *bool) mapNode[K, V] {
	_ = "STUB: not implemented"
	return nil

	// Return original node if key is not found.
}

// Mark as resized if key exists.

// Convert to value node if we move to one entry.

// Remove entry in-place if mutable.

// Return copy without entry if immutable.

// mergeIntoNode merges a key/value pair into an existing node.
// Caller must verify that node's keyHash is not equal to keyHash.
func mergeIntoNode[K, V any](node mapLeafNode[K, V], shift uint, keyHash uint32, key K, value V) mapNode[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Recursively build branch nodes to combine the node and its key.

// mapEntry represents a single key/value pair.
type mapEntry[K, V any] struct {
	key   K
	value V
}

// MapIterator represents an iterator over a map's key/value pairs. Although
// map keys are not sorted, the iterator's order is deterministic.
type MapIterator[K, V any] struct {
	m *Map[K, V] // source map

	stack [32]mapIteratorElem[K, V] // search stack
	depth int                       // stack depth
}

// Done returns true if no more elements remain in the iterator.
func (itr *MapIterator[K, V]) Done() bool { _ = "STUB: not implemented"; return false }

// First resets the iterator to the first key/value pair.
func (itr *MapIterator[K, V]) First() {
	_ = "STUB: not implemented"
	// Exit immediately if the map is empty.
	return
}

// Initialize the stack to the left most element.

// Next returns the next key/value pair. Returns a nil key when no elements remain.
func (itr *MapIterator[K, V]) Next() (key K, value V, ok bool) {
	_ = "STUB: not implemented"
	// Return nil key if iteration is done.
	return *new(K), *new(V), false
}

// Retrieve current index & value. Current node is always a leaf.

// Move up stack until we find a node that has remaining position ahead
// and move that element forward by one.

// next moves to the next available key.
func (itr *MapIterator[K, V]) next() { _ = "STUB: not implemented"; return }

// always the last value, traverse up

// first positions the stack left most index.
// Elements and indexes at and below the current depth are assumed to be correct.
func (itr *MapIterator[K, V]) first() { _ = "STUB: not implemented"; return }

// find first node

// *mapArrayNode, mapLeafNode

// mapIteratorElem represents a node/index pair in the MapIterator stack.
type mapIteratorElem[K, V any] struct {
	node  mapNode[K, V]
	index int
}

// Sorted map child node limit size.
const (
	sortedMapNodeSize = 32
)

// SortedMap represents a map of key/value pairs sorted by key. The sort order
// is determined by the Comparer used by the map.
//
// This map is implemented as a B+tree.
type SortedMap[K, V any] struct {
	size     int                 // total number of key/value pairs
	root     sortedMapNode[K, V] // root of b+tree
	comparer Comparer[K]
}

// NewSortedMap returns a new instance of SortedMap. If comparer is nil then
// a default comparer is set after the first key is inserted. Default comparers
// exist for int, string, and byte slice keys.
func NewSortedMap[K, V any](comparer Comparer[K]) *SortedMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// NewSortedMapOf returns a new instance of SortedMap, containing a map of provided entries.
//
// If comparer is nil then a default comparer is set after the first key is inserted. Default comparers
// exist for int, string, and byte slice keys.
func NewSortedMapOf[K comparable, V any](comparer Comparer[K], entries map[K]V) *SortedMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Len returns the number of elements in the sorted map.
func (m *SortedMap[K, V]) Len() int {
	_ = "STUB: not implemented"

	// Get returns the value for a given key and a flag indicating if the key is set.
	// The flag can be used to distinguish between a nil-set key versus an unset key.
	return 0
}

func (m *SortedMap[K, V]) Get(key K) (V, bool) { _ = "STUB: not implemented"; return *new(V), false }

// Set returns a copy of the map with the key set to the given value.
func (m *SortedMap[K, V]) Set(key K, value V) *SortedMap[K, V] {
	_ = "STUB: not implemented"
	return nil
}

func (m *SortedMap[K, V]) set(key K, value V, mutable bool) *SortedMap[K, V] {
	_ = "STUB: not implemented"
	// Set a comparer on the first value if one does not already exist.
	return nil
}

// Create copy, if necessary.

// If no values are set then initialize with a leaf node.

// Otherwise delegate to root node.
// If a split occurs then grow the tree from the root.

// Update root and size (if resized).

// Delete returns a copy of the map with the key removed.
// Returns the original map if key does not exist.
func (m *SortedMap[K, V]) Delete(key K) *SortedMap[K, V] { _ = "STUB: not implemented"; return nil }

func (m *SortedMap[K, V]) delete(key K, mutable bool) *SortedMap[K, V] {
	_ = "STUB: not implemented"
	// Return original map if no keys exist.
	return nil
}

// If the delete did not change the node then return the original map.

// Create copy, if necessary.

// Update root and size.

// clone returns a shallow copy of m.
func (m *SortedMap[K, V]) clone() *SortedMap[K, V] { _ = "STUB: not implemented"; return nil }

// Iterator returns a new iterator for this map positioned at the first key.
func (m *SortedMap[K, V]) Iterator() *SortedMapIterator[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// SortedMapBuilder represents an efficient builder for creating sorted maps.
type SortedMapBuilder[K, V any] struct {
	m *SortedMap[K, V] // current state
}

// NewSortedMapBuilder returns a new instance of SortedMapBuilder.
func NewSortedMapBuilder[K, V any](comparer Comparer[K]) *SortedMapBuilder[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// SortedMap returns the current copy of the map.
// The returned map is safe to use even if after the builder continues to be used.
func (b *SortedMapBuilder[K, V]) Map() *SortedMap[K, V] { _ = "STUB: not implemented"; return nil }

// Len returns the number of elements in the underlying map.
func (b *SortedMapBuilder[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

// Get returns the value for the given key.
func (b *SortedMapBuilder[K, V]) Get(key K) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// Set sets the value of the given key. See SortedMap.Set() for additional details.
func (b *SortedMapBuilder[K, V]) Set(key K, value V) { _ = "STUB: not implemented"; return }

// Delete removes the given key. See SortedMap.Delete() for additional details.
func (b *SortedMapBuilder[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

// Iterator returns a new iterator for the underlying map positioned at the first key.
func (b *SortedMapBuilder[K, V]) Iterator() *SortedMapIterator[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// sortedMapNode represents a branch or leaf node in the sorted map.
type sortedMapNode[K, V any] interface {
	minKey() K
	indexOf(key K, c Comparer[K]) int
	get(key K, c Comparer[K]) (value V, ok bool)
	set(key K, value V, c Comparer[K], mutable bool, resized *bool) (sortedMapNode[K, V], sortedMapNode[K, V])
	delete(key K, c Comparer[K], mutable bool, resized *bool) sortedMapNode[K, V]
}

var _ sortedMapNode[string, any] = (*sortedMapBranchNode[string, any])(nil)
var _ sortedMapNode[string, any] = (*sortedMapLeafNode[string, any])(nil)

// sortedMapBranchNode represents a branch in the sorted map.
type sortedMapBranchNode[K, V any] struct {
	elems []sortedMapBranchElem[K, V]
}

// newSortedMapBranchNode returns a new branch node with the given child nodes.
func newSortedMapBranchNode[K, V any](children ...sortedMapNode[K, V]) *sortedMapBranchNode[K, V] {
	_ = "STUB: not implemented"
	// Fetch min keys for every child.
	return nil
}

// minKey returns the lowest key stored in this node's tree.
func (n *sortedMapBranchNode[K, V]) minKey() K { _ = "STUB: not implemented"; return *new(K) }

// indexOf returns the index of the key within the child nodes.
func (n *sortedMapBranchNode[K, V]) indexOf(key K, c Comparer[K]) int {
	_ = "STUB: not implemented"
	return 0
}

// get returns the value for the given key.
func (n *sortedMapBranchNode[K, V]) get(key K, c Comparer[K]) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *new(V), false
}

// set returns a copy of the node with the key set to the given value.
func (n *sortedMapBranchNode[K, V]) set(key K, value V, c Comparer[K], mutable bool, resized *bool) (sortedMapNode[K, V], sortedMapNode[K, V]) {
	_ = "STUB: not implemented"
	return nil,

		// Delegate insert to child node.
		nil
}

// Update in-place, if mutable.

// If the child splits and we have no more room then we split too.

// If no split occurs, copy branch and update keys.
// If the child splits, insert new key/child into copy of branch.

// If the child splits and we have no more room then we split too.

// Otherwise return the new branch node with the updated entry.

// delete returns a node with the key removed. Returns the same node if the key
// does not exist. Returns nil if all child nodes are removed.
func (n *sortedMapBranchNode[K, V]) delete(key K, c Comparer[K], mutable bool, resized *bool) sortedMapNode[K, V] {
	_ = "STUB: not implemented"
	return nil

	// Return original node if child has not changed.
}

// Remove child if it is now nil.

// If this node will become empty then simply return nil.

// If mutable, update in-place.

// Return a copy without the given node.

// If mutable, update in-place.

// Return a copy with the updated node.

type sortedMapBranchElem[K, V any] struct {
	key  K
	node sortedMapNode[K, V]
}

// sortedMapLeafNode represents a leaf node in the sorted map.
type sortedMapLeafNode[K, V any] struct {
	entries []mapEntry[K, V]
}

// minKey returns the first key stored in this node.
func (n *sortedMapLeafNode[K, V]) minKey() K {
	_ = "STUB: not implemented"
	return *

	// indexOf returns the index of the given key.
	new(K)
}

func (n *sortedMapLeafNode[K, V]) indexOf(key K, c Comparer[K]) int {
	_ = "STUB: not implemented"
	return 0
}

// GTE

// get returns the value of the given key.
func (n *sortedMapLeafNode[K, V]) get(key K, c Comparer[K]) (value V, ok bool) {
	_ = "STUB: not implemented"
	return *

	// If the index is beyond the entry count or the key is not equal then return 'not found'.
	new(V), false
}

// If the key matches then return its value.

// set returns a copy of node with the key set to the given value. If the update
// causes the node to grow beyond the maximum size then it is split in two.
func (n *sortedMapLeafNode[K, V]) set(key K, value V, c Comparer[K], mutable bool, resized *bool) (sortedMapNode[K, V], sortedMapNode[K, V]) {
	_ = "STUB: not implemented"
	// Find the insertion index for the key.
	return nil, nil
}

// Update in-place, if mutable.

// If the key doesn't exist and we exceed our max allowed values then split.

// If the key matches then simply return a copy with the entry overridden.
// If there is no match then insert new entry and mark as resized.

// If the key doesn't exist and we exceed our max allowed values then split.

// Otherwise return the new leaf node with the updated entry.

// delete returns a copy of node with key removed. Returns the original node if
// the key does not exist. Returns nil if the removed key is the last remaining key.
func (n *sortedMapLeafNode[K, V]) delete(key K, c Comparer[K], mutable bool, resized *bool) sortedMapNode[K, V] {
	_ = "STUB: not implemented"
	return nil

	// Return original node if key is not found.
}

// If this is the last entry then return nil.

// Update in-place, if mutable.

// Return copy of node with entry removed.

// SortedMapIterator represents an iterator over a sorted map.
// Iteration can occur in natural or reverse order based on use of Next() or Prev().
type SortedMapIterator[K, V any] struct {
	m *SortedMap[K, V] // source map

	stack [32]sortedMapIteratorElem[K, V] // search stack
	depth int                             // stack depth
}

// Done returns true if no more key/value pairs remain in the iterator.
func (itr *SortedMapIterator[K, V]) Done() bool { _ = "STUB: not implemented"; return false }

// First moves the iterator to the first key/value pair.
func (itr *SortedMapIterator[K, V]) First() { _ = "STUB: not implemented"; return }

// Last moves the iterator to the last key/value pair.
func (itr *SortedMapIterator[K, V]) Last() { _ = "STUB: not implemented"; return }

// Seek moves the iterator position to the given key in the map.
// If the key does not exist then the next key is used. If no more keys exist
// then the iteartor is marked as done.
func (itr *SortedMapIterator[K, V]) Seek(key K) { _ = "STUB: not implemented"; return }

// Next returns the current key/value pair and moves the iterator forward.
// Returns a nil key if the there are no more elements to return.
func (itr *SortedMapIterator[K, V]) Next() (key K, value V, ok bool) {
	_ = "STUB: not implemented"
	// Return nil key if iteration is complete.
	return *new(K), *new(V), false
}

// Retrieve current key/value pair.

// Move to the next available key/value pair.

// Only occurs when iterator is done.

// next moves to the next key. If no keys are after then depth is set to -1.
func (itr *SortedMapIterator[K, V]) next() { _ = "STUB: not implemented"; return }

// Prev returns the current key/value pair and moves the iterator backward.
// Returns a nil key if the there are no more elements to return.
func (itr *SortedMapIterator[K, V]) Prev() (key K, value V, ok bool) {
	_ = "STUB: not implemented"
	// Return nil key if iteration is complete.
	return *new(K), *new(V), false
}

// Retrieve current key/value pair.

// prev moves to the previous key. If no keys are before then depth is set to -1.
func (itr *SortedMapIterator[K, V]) prev() { _ = "STUB: not implemented"; return }

// first positions the stack to the leftmost key from the current depth.
// Elements and indexes below the current depth are assumed to be correct.
func (itr *SortedMapIterator[K, V]) first() { _ = "STUB: not implemented"; return }

// last positions the stack to the rightmost key from the current depth.
// Elements and indexes below the current depth are assumed to be correct.
func (itr *SortedMapIterator[K, V]) last() { _ = "STUB: not implemented"; return }

// seek positions the stack to the given key from the current depth.
// Elements and indexes below the current depth are assumed to be correct.
func (itr *SortedMapIterator[K, V]) seek(key K) { _ = "STUB: not implemented"; return }

// sortedMapIteratorElem represents node/index pair in the SortedMapIterator stack.
type sortedMapIteratorElem[K, V any] struct {
	node  sortedMapNode[K, V]
	index int
}

// Hasher hashes keys and checks them for equality.
type Hasher[K any] interface {
	// Computes a hash for key.
	Hash(key K) uint32

	// Returns true if a and b are equal.
	Equal(a, b K) bool
}

// NewHasher returns the built-in hasher for a given key type.
func NewHasher[K any](key K) Hasher[K] {
	_ = "STUB: not implemented"
	// Attempt to use non-reflection based hasher first.
	return nil
}

// Fallback to reflection-based hasher otherwise.
// This is used when caller wraps a type around a primitive type.

// If no hashers match then panic.
// This is a compile time issue so it should not return an error.

// Hash returns a hash for value.
func hashString(value string) uint32 { _ = "STUB: not implemented"; return 0 }

// reflectIntHasher implements a reflection-based Hasher for keys.
type reflectHasher[K any] struct{}

// Hash returns a hash for key.
func (h *reflectHasher[K]) Hash(key K) uint32 { _ = "STUB: not implemented"; return 0 }

// Equal returns true if a is equal to b. Otherwise returns false.
// Panics if a and b are not int-ish or string-ish.
func (h *reflectHasher[K]) Equal(a, b K) bool { _ = "STUB: not implemented"; return false }

// hashUint64 returns a 32-bit hash for a 64-bit value.
func hashUint64(value uint64) uint32 { _ = "STUB: not implemented"; return 0 }

// defaultHasher implements Hasher.
type defaultHasher[K any] struct{}

// Hash returns a hash for key.
func (h *defaultHasher[K]) Hash(key K) uint32 { _ = "STUB: not implemented"; return 0 }

// Equal returns true if a is equal to b. Otherwise returns false.
// Panics if a and b are not comparable.
func (h *defaultHasher[K]) Equal(a, b K) bool { _ = "STUB: not implemented"; return false }

// Comparer allows the comparison of two keys for the purpose of sorting.
type Comparer[K any] interface {
	// Returns -1 if a is less than b, returns 1 if a is greater than b,
	// and returns 0 if a is equal to b.
	Compare(a, b K) int
}

// NewComparer returns the built-in comparer for a given key type.
// Note that only int-ish and string-ish types are supported, despite the 'comparable' constraint.
// Attempts to use other types will result in a panic - users should define their own Comparers for these cases.
func NewComparer[K any](key K) Comparer[K] {
	_ = "STUB: not implemented"
	// Attempt to use non-reflection based comparer first.
	return nil
}

// Fallback to reflection-based comparer otherwise.
// This is used when caller wraps a type around a primitive type.

// If no comparers match then panic.
// This is a compile time issue so it should not return an error.

// defaultComparer compares two values (int-ish and string-ish types are supported). Implements Comparer.
type defaultComparer[K any] struct{}

// Compare returns -1 if a is less than b, returns 1 if a is greater than b, and
// returns 0 if a is equal to b. Panic if a or b is not a string or int* type
func (c *defaultComparer[K]) Compare(i K, j K) int { _ = "STUB: not implemented"; return 0 }

// defaultCompare only operates on constraints.Ordered.
// For other types, users should bring their own comparers
func defaultCompare[K constraints.Ordered](i, j K) int { _ = "STUB: not implemented"; return 0 }

// reflectIntComparer compares two values using reflection. Implements Comparer.
type reflectComparer[K any] struct{}

// Compare returns -1 if a is less than b, returns 1 if a is greater than b, and
// returns 0 if a is equal to b. Panic if a or b is not an int-ish or string-ish type.
func (c *reflectComparer[K]) Compare(a, b K) int { _ = "STUB: not implemented"; return 0 }

func assert(condition bool, message string) { _ = "STUB: not implemented"; return }
