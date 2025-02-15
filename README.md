# BFQ - BlazinglyFastQueue

`BFQ` (BlazinglyFastQueue) is a high-performance, non-thread-safe queue implementation in Go. It leverages optimized memory management and efficient resizing to provide fast operations for typical use cases. While it isn't safe for concurrent use, it shines in single-threaded environments where performance is critical.

## Features

- **Fast Operations**: Designed for maximum speed in non-concurrent environments.
- **Efficient Resizing**: Dynamically adjusts the size of the internal array to avoid unnecessary allocations.
- **Non-Thread-Safe**: Ideal for single-threaded or controlled concurrency scenarios. No locking overhead.
- **Generic**: Written with Go generics, so you can use it with any data type.

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
go test -bench .
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
