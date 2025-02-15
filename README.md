# BFQ - BlazinglyFastQueue

`BFQ` (BlazinglyFastQueue) is a high-performance, **non-thread-safe** queue implementation in Go. It leverages optimized memory management and efficient resizing to provide fast operations for typical use cases. While it isn't safe for concurrent use, it shines in single-threaded environments where performance is critical.

## Features

- **Fast Operations**: Designed for maximum speed in non-concurrent environments.
- **Efficient Resizing**: Dynamically adjusts the size of the internal array to avoid unnecessary allocations.
- **Non-Thread-Safe**: Ideal for single-threaded or controlled concurrency scenarios. No locking overhead.
- **Generic**: Written with Go generics, so you can use it with any data type.

## Why It Works Fast

- **Circular Buffer Approach**: 
   The queue is backed by a **circular buffer**, which allows elements to be pushed and popped with constant time complexity (`O(1)`), regardless of the size of the queue. The wrap-around is efficiently handled using **bitwise AND** (`&`) with the capacity of the buffer, which is always a power of two. This avoids the need for costly modulus operations and provides a faster way to determine the correct indices for the front and back of the queue.

- **Optimized Resizing with Bitwise Operations**:  
   All operations that previously required division or multiplication (such as resizing the buffer) have been converted to bitwise operations. This reduces computational overhead and improves performance by utilizing faster bitwise shifts instead of more expensive arithmetic operations. This approach ensures efficient memory handling and quick resizing without unnecessary allocations.

- **Efficient Memory Management**:
   The queue minimizes memory allocations by resizing only when necessary (e.g., when the queue is full). Additionally, the queue shrinks the underlying array when it's much larger than needed, reducing memory usage.

- **Unsafe Pointer Manipulation**:
   The `indexUnsafe` function uses **unsafe pointers** to directly access elements in the underlying array without bounds checking. This significantly speeds up the access to elements, as it avoids the runtime checks that Go typically applies when indexing arrays.

- **Constant Time Operations**:
   Both `PushFront` and `PushBack` operations are designed to be performed in constant time (`O(1)`), as the front and back indices are updated using simple modulo arithmetic. The same is true for `PopFront` and `PopBack`, which operate in constant time without needing to shift elements.

- **Minimal Overhead**:
   Since `BFQ` is **not thread-safe**, it avoids the overhead of locking mechanisms. In single-threaded or non-concurrent environments, this reduces unnecessary computation, making `BFQ` faster than thread-safe alternatives.

## Usage

### Creating a New Queue

To create a new queue:

```go
package main

import "github.com/phtea/bfq"

func main() {
    q := bfq.NewQueue() // Create a new queue
    q.PushBack(42)      // Add an element to the queue
    q.PushFront(7)      // Add an element at the front
    q.PopFront()        // Remove an element from the front
}
```

In addition to creating an empty queue with `NewQueue`, you can also create a queue from an existing slice. The `FromSlice` function ensures the internal array size is a power of two for better performance. 

```go
q := bfq.FromSlice([]int{1, 2, 3, 4})
```

The `FromSlice` function does the following:
- Takes a slice as input and calculates the appropriate internal buffer size, which will always be a power of two.
- Initializes a new queue with that capacity, copying the elements of the slice into the queue.

### Pushing and Popping

```go
q := bfq.NewQueue()

q.PushBack(1)  // Adds an element to the back
q.PushFront(2) // Adds an element to the front

frontElem, _ := q.PopFront() // Removes and returns the element from the front
backElem, _ := q.PopBack()  // Removes and returns the element from the back
```

### Queue Operations

- **PushBack(v T)**: Adds an element to the back of the queue.
- **PushFront(v T)**: Adds an element to the front of the queue.
- **PopFront() (T, bool)**: Removes and returns the front element of the queue.
- **PopBack() (T, bool)**: Removes and returns the back element of the queue.
- **Front() (T, bool)**: Returns the front element without removing it.
- **Back() (T, bool)**: Returns the back element without removing it.
- **Len()**: Returns the current number of elements in the queue.
- **IsEmpty()**: Checks if the queue is empty.

## Important Notes

- **Not Thread-Safe**: `BFQ` is **not thread-safe** and should not be used concurrently without proper synchronization. If you need a thread-safe queue, consider using Go's built-in channels or adding your own synchronization mechanism.
  
- **Designed for Speed**: By avoiding the overhead of locks and other concurrency features, `BFQ` can outperform other queue implementations in single-threaded scenarios.

## Benchmark

Benchmarks show that `BFQ` is fast and efficient in scenarios where concurrency is not required. Here's an example of the benchmarking results:

```bash
$ go test -bench=. -benchtime=10000000x
goos: linux
goarch: amd64
pkg: bfq
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkQueueCreation-6                10000000              6333 ns/op           16256 B/op          7 allocs/op
BenchmarkQueuePushBack-6                10000000                 8.549 ns/op          26 B/op          0 allocs/op
BenchmarkQueuePushFront-6               10000000                 7.193 ns/op          26 B/op          0 allocs/op
BenchmarkQueueFromSlice-6               10000000                 1.698 ns/op          13 B/op          0 allocs/op
BenchmarkQueuePopBack-6                 10000000                 4.964 ns/op          13 B/op          0 allocs/op
BenchmarkQueuePopFront-6                10000000                 4.796 ns/op          13 B/op          0 allocs/op
BenchmarkQueueFront-6                   10000000                 0.8059 ns/op          0 B/op          0 allocs/op
BenchmarkQueueBack-6                    10000000                 0.5316 ns/op          0 B/op          0 allocs/op
BenchmarkQueueIsEmpty-6                 10000000                 0.2623 ns/op          0 B/op          0 allocs/op
BenchmarkQueueMixedOperations-6         10000000                 6.594 ns/op          13 B/op          0 allocs/op
BenchmarkQueueWithInts-6                10000000                 4.742 ns/op           0 B/op          0 allocs/op
BenchmarkQueueWithStruct-6              10000000               105.0 ns/op            24 B/op          2 allocs/op
BenchmarkQueueWithStructPointers-6      10000000               149.4 ns/op            48 B/op          3 allocs/op
BenchmarkQueueWithSlices-6              10000000                27.46 ns/op           24 B/op          1 allocs/op
PASS
ok      bfq     66.710s
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
