
# Go Concurrency: Order Processing

A practical Go project demonstrating goroutines, synchronization,
shared memory, and data race detection through an order-processing
simulation.

## Learning Objectives

- Understand goroutines and concurrent execution.
- Coordinate goroutines using `sync.WaitGroup`.
- Protect shared data with `sync.Mutex`.
- Allow concurrent readers with `sync.RWMutex`.
- Detect data races using Go's race detector.
- Understand the difference between synchronization and task completion.

## Concepts Covered

### 1. Goroutines

Goroutines allow functions to execute concurrently.

```go
go processOrders(orders)
go updateOrderStatus(orders)
go reportOrderStatus(orders)
```

**Important:** The program does not automatically wait for goroutines
when `main()` returns.

### 2. WaitGroup

A `sync.WaitGroup` waits for a collection of goroutines to finish.

```go
var wg sync.WaitGroup

wg.Add(1)

go func() {
    defer wg.Done()
    processOrders(orders)
}()

wg.Wait()
```

- `Add()` increments the counter.
- `Done()` decrements the counter.
- `Wait()` blocks until the counter reaches zero.

### 3. Mutex

A `sync.Mutex` provides exclusive access to shared data.

```go
mu.Lock()
order.Status = "Shipped"
mu.Unlock()
```

Every conflicting access must follow the same synchronization
strategy to avoid data races.

### 4. RWMutex

A `sync.RWMutex` supports concurrent readers and exclusive writers.

Read shared data:

```go
mu.RLock()
status := order.Status
mu.RUnlock()
```

Write shared data:

```go
mu.Lock()
order.Status = "Delivered"
mu.Unlock()
```

Multiple readers can hold read locks simultaneously. A writer needs
exclusive access.

### 5. Race Detector

Go includes a built-in race detector.

```bash
go run -race main.go
```

The race detector reports data races encountered during execution.
A clean run does not prove that every possible execution is race-free.

## Requirements

- Go 1.22 or later
- Git

## Run the Project

Clone the repository:

```bash
git clone https://github.com/Ramadan-Lotfy/go-concurrency-order-processing.git
cd go-concurrency-order-processing
```

Run the application:

```bash
go run main.go
```

Run with race detection:

```bash
go run -race main.go
```

## Important Concurrency Considerations

### WaitGroup vs Mutex

A `WaitGroup` coordinates task completion.

A `Mutex` protects shared memory.

They solve different problems and may be needed together.

### Data Race vs Logical Race

A data race occurs when conflicting memory accesses happen without
proper synchronization.

A logical race occurs when concurrent operations produce an
unexpected result even though memory access is synchronized.

For example, a report may capture some order statuses before updates
and others after updates. This is not necessarily a data race, but
the report may not represent a single consistent snapshot.

### Safe Reporting

To produce a final report, wait for all status updates to finish
before generating it.

To produce an atomic snapshot while updates continue, protect the
snapshot operation with an appropriate shared synchronization
strategy.

## Project Structure

```text
go-concurrency-order-processing/
├── .gitignore
├── README.md
└── main.go
```

## Key Takeaways

- Goroutines enable concurrent execution.
- WaitGroups coordinate completion.
- Mutexes provide exclusive access to shared state.
- RWMutexes allow concurrent reads and exclusive writes.
- The race detector helps identify unsafe memory access.
- Correct synchronization is essential for reliable concurrent
  applications.

## License

This project is available under the MIT License.
Add a LICENSE file if you intend to distribute it under that license.
