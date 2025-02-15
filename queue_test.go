package bfq

import (
	"fmt"
	"testing"
)

func BenchmarkNewQueueCreation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q := NewQueue[int]()
		// Push enough elements to trigger a resize and allocation
        for j := 0; j < 1000; j++ { 
            q.PushBack(j)
        }
	}
}

func BenchmarkNewQueuePushBack(b *testing.B) {
	b.ReportAllocs()
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
}

func BenchmarkNewQueuePushFront(b *testing.B) {
	b.ReportAllocs()
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.PushFront(i)
	}
}

func BenchmarkNewQueueFromSlice(b *testing.B) {
	arr := make([]int, b.N)
	for i:=0; i<b.N; i++ {
		arr[i]=i
	}
	b.ResetTimer()
	b.ReportAllocs()
	_ = FromSlice(arr)
}

func BenchmarkNewQueuePopBack(b *testing.B) {
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q.PopBack()
	}
}

func BenchmarkNewQueuePopFront(b *testing.B) {
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q.PopFront()
	}
}

func BenchmarkNewQueueFront(b *testing.B) {
	q := NewQueue[int]()
	q.PushBack(42)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q.Front()
	}
}

func BenchmarkNewQueueBack(b *testing.B) {
	q := NewQueue[int]()
	q.PushBack(42)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q.Back()
	}
}

func BenchmarkNewQueueIsEmpty(b *testing.B) {
	q := NewQueue[int]()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		q.IsEmpty()
	}
}

func BenchmarkNewQueueMixedOperations(b *testing.B) {
	b.ReportAllocs()
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
		if i%2 == 0 {
			q.PopFront()
		}
	}
}

// Different data types
func BenchmarkQueueWithInts(b *testing.B) {
	b.ReportAllocs()
	q := NewQueue[int]()
	for i := 0; i < b.N; i++ {
		q.PushBack(42)
		q.PopFront()
	}
}

type Data struct {
	ID   int
	Name string
}

func BenchmarkQueueWithStruct(b *testing.B) {
	b.ReportAllocs()
	q := NewQueue[Data]()
	for i := 0; i < b.N; i++ {
		q.PushBack(Data{ID: i, Name: fmt.Sprintf("Item %d", i)})
		q.PopFront()
	}
}

func BenchmarkQueueWithStructPointers(b *testing.B) {
	b.ReportAllocs()
	q := NewQueue[*Data]()
	for i := 0; i < b.N; i++ {
		item := &Data{ID: i, Name: fmt.Sprintf("Item %d", i)}
		q.PushBack(item)
		q.PopFront()
	}
}

func BenchmarkQueueWithSlices(b *testing.B) {
	b.ReportAllocs()
	q := NewQueue[[]int]()
	for i := 0; i < b.N; i++ {
		q.PushBack([]int{i, i + 1, i + 2})
		q.PopFront()
	}
}

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	if q.Len() != 0 {
		t.Errorf("Expected queue length 0, got %d", q.Len())
	}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}
}

func TestPushFrontAndPopFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(10)
	q.PushFront(20)

	// Ensure the front is correct
	if front, ok := q.Front(); !ok || front != 20 {
		t.Errorf("Expected front element 20, got %v", front)
	}

	// Pop the front and check the result
	if val, ok := q.PopFront(); !ok || val != 20 {
		t.Errorf("Expected popped value 20, got %v", val)
	}

	// Ensure the front is now the next element
	if front, ok := q.Front(); !ok || front != 10 {
		t.Errorf("Expected front element 10, got %v", front)
	}

	// Test popping last element
	if val, ok := q.PopFront(); !ok || val != 10 {
		t.Errorf("Expected popped value 10, got %v", val)
	}

	// Ensure the queue is empty after popping
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}
}

func TestPushBackAndPopBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(10)
	q.PushBack(20)

	// Ensure the back is correct
	if back, ok := q.Back(); !ok || back != 20 {
		t.Errorf("Expected back element 20, got %v", back)
	}

	// Pop the back and check the result
	if val, ok := q.PopBack(); !ok || val != 20 {
		t.Errorf("Expected popped value 20, got %v", val)
	}

	// Ensure the back is now the next element
	if back, ok := q.Back(); !ok || back != 10 {
		t.Errorf("Expected back element 10, got %v", back)
	}

	// Test popping last element
	if val, ok := q.PopBack(); !ok || val != 10 {
		t.Errorf("Expected popped value 10, got %v", val)
	}

	// Ensure the queue is empty after popping
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}
}

func TestFrontAndBackOnEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	// Test Front on empty queue
	if _, ok := q.Front(); ok {
		t.Errorf("Expected Front to return false on empty queue")
	}

	// Test Back on empty queue
	if _, ok := q.Back(); ok {
		t.Errorf("Expected Back to return false on empty queue")
	}
}

func TestPushPopMix(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(10)
	q.PushFront(20)

	// Test mixed operations: PushFront, PushBack, PopFront, PopBack
	if front, ok := q.Front(); !ok || front != 20 {
		t.Errorf("Expected front element 20, got %v", front)
	}
	if back, ok := q.Back(); !ok || back != 10 {
		t.Errorf("Expected back element 10, got %v", back)
	}

	// Pop front
	if val, ok := q.PopFront(); !ok || val != 20 {
		t.Errorf("Expected popped value 20, got %v", val)
	}

	// Pop back
	if val, ok := q.PopBack(); !ok || val != 10 {
		t.Errorf("Expected popped value 10, got %v", val)
	}

	// Ensure the queue is empty
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}
}

func TestResize(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 100; i++ {
		q.PushBack(i)
	}

	// Test the queue length after inserting 100 elements
	if q.Len() != 100 {
		t.Errorf("Expected queue length 100, got %d", q.Len())
	}

	// Pop all elements to check if everything is in order
	for i := 0; i < 100; i++ {
		if val, ok := q.PopFront(); !ok || val != i {
			t.Errorf("Expected popped value %d, got %v", i, val)
		}
	}

	// Ensure the queue is empty
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}
}

func TestPushFrontOverflow(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 1000; i++ {
		q.PushFront(i)
	}

	// Test the queue length after inserting 1000 elements
	if q.Len() != 1000 {
		t.Errorf("Expected queue length 1000, got %d", q.Len())
	}

	// Pop all elements to check if everything is in order (LIFO)
	for i := 999; i >= 0; i-- {
		if val, ok := q.PopFront(); !ok || val != i {
			t.Errorf("Expected popped value %d, got %v", i, val)
		}
	}

	// Ensure the queue is empty
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}
}

func TestPopFromEmptyQueue(t *testing.T) {
	q := NewQueue[int]()
	if val, ok := q.PopFront(); ok {
		t.Errorf("Expected PopFront to return false, got %v", val)
	}
	if val, ok := q.PopBack(); ok {
		t.Errorf("Expected PopBack to return false, got %v", val)
	}
}

func TestQueueWithNilPointers(t *testing.T) {
	q := NewQueue[*Data]()
	q.PushBack(nil)

	// Test popping a nil value
	if val, ok := q.PopFront(); !ok || val != nil {
		t.Errorf("Expected nil value, got %v", val)
	}

	// Ensure the queue is empty after popping
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}
}

func TestLargeQueueOperations(t *testing.T) {
	q := NewQueue[int]()
	for i := 0; i < 100000; i++ {
		q.PushBack(i)
	}

	// Ensure the queue size is correct
	if q.Len() != 100000 {
		t.Errorf("Expected queue size 100000, got %d", q.Len())
	}

	// Pop all elements and check if it's in order
	for i := 0; i < 100000; i++ {
		if val, ok := q.PopFront(); !ok || val != i {
			t.Errorf("Expected popped value %d, got %v", i, val)
		}
	}

	// Ensure the queue is empty
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it's not")
	}
}
