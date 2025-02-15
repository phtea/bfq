package bfq

import (
	"fmt"
	"strings"
	"unsafe"
)

// Queue represents a double-ended queue using a circular buffer.
type Queue[T any] struct {
	buf    []T
	front  int
	back   int
	length int
}

const (
	minCapacity = 8
)

// NewQueue creates an empty queue with an initial capacity.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{buf: make([]T, minCapacity)}
}

// nextPowerOfTwo returns the smallest power of two greater than or equal to n.
func nextPowerOfTwo(n int) int {
	if n < minCapacity {
		return minCapacity
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32 // For 64-bit integers
	n++
	return n
}

// FromSlice creates a queue from a given slice, ensuring the buffer size is a power of two.
func FromSlice[T any](slice []T) *Queue[T] {
	size := nextPowerOfTwo(len(slice))
	q := &Queue[T]{buf: make([]T, size), front: 0, back: len(slice), length: len(slice)}
	copy(q.buf, slice)
	return q
}

// Len returns the number of elements in the queue.
func (q *Queue[T]) Len() int { return q.length }

// IsEmpty checks if the queue is empty.
func (q *Queue[T]) IsEmpty() bool { return q.length == 0 }

// resize resizes the queue when needed.
func (q *Queue[T]) resize(size int) {
	newBuf := make([]T, size)
	if q.front < q.back {
		copy(newBuf, q.buf[q.front:q.back])
	} else {
		n := copy(newBuf, q.buf[q.front:])
		copy(newBuf[n:], q.buf[:q.back])
	}
	q.buf = newBuf
	q.front = 0
	q.back = q.length
}

// grow expands the queue when full.
func (q *Queue[T]) grow() {
	if q.length == len(q.buf) {
		q.resize(len(q.buf) << 1)
	}
}

// shrink reduces memory usage when necessary.
func (q *Queue[T]) shrink() {
	if q.length > minCapacity && q.length == len(q.buf) >> 2 {
		q.resize(len(q.buf) >> 1)
	}
}

// indexUnsafe gets the pointer to an element without bounds checks.
func (q *Queue[T]) indexUnsafe(index int) *T {
	base := unsafe.Pointer(&q.buf[0]) // Base address of buffer
	size := unsafe.Sizeof(q.buf[0])   // Size of one element
	return (*T)(unsafe.Pointer(uintptr(base) + uintptr(index)*size))
}

// PushFront inserts an element at the front.
func (q *Queue[T]) PushFront(v T) {
	q.grow()
	q.front = (q.front - 1 + len(q.buf)) & (len(q.buf) - 1)
	*(*T)(unsafe.Pointer(q.indexUnsafe(q.front))) = v
	q.length++
}

// PushBack inserts an element at the back.
func (q *Queue[T]) PushBack(v T) {
	q.grow()
	*(*T)(unsafe.Pointer(q.indexUnsafe(q.back))) = v
	q.back = (q.back + 1) & (len(q.buf) - 1)
	q.length++
}

// PopFront removes and returns the front element.
func (q *Queue[T]) PopFront() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	v := *q.indexUnsafe(q.front)
	q.front = (q.front + 1) & (len(q.buf) - 1)
	q.length--
	q.shrink()
	return v, true
}

// PopBack removes and returns the back element.
func (q *Queue[T]) PopBack() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	q.back = (q.back - 1 + len(q.buf)) & (len(q.buf) - 1)
	v := *q.indexUnsafe(q.back)
	q.length--
	q.shrink()
	return v, true
}

// Front returns the first element without removing it.
func (q *Queue[T]) Front() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return *q.indexUnsafe(q.front), true
}

// Back returns the last element without removing it.
func (q *Queue[T]) Back() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return *q.indexUnsafe((q.back - 1 + len(q.buf)) & (len(q.buf) - 1)), true
}

// String returns a string representation of the queue.
func (q *Queue[T]) String() string {
	var sb strings.Builder
	sb.WriteByte('[')
	for i, idx := 0, q.front; i < q.length; i++ {
		if i > 0 {
			sb.WriteByte(' ')
		}
		sb.WriteString(fmt.Sprintf("%v", *q.indexUnsafe(idx)))
		idx = (idx + 1) & (len(q.buf) - 1)
	}
	sb.WriteByte(']')
	return sb.String()
}
