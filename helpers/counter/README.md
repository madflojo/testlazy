# Counter (helpers/counter)

A tiny, test-friendly, thread-safe counter for Go. It supports atomic
increments/decrements, direct reads/sets, and simple wait helpers to block
until the value crosses a threshold.

## Module

- Module path: `github.com/madflojo/testlazy/helpers/counter`
- Install: `go get github.com/madflojo/testlazy/helpers/counter`

## Why

- Minimal API tailored for tests.
- Safe for concurrent use by multiple goroutines.
- Non-blocking wait helpers that integrate naturally with `select` and timeouts.

## Usage

Basic operations:

```go
c := counter.New()
c.Increment()
c.Add(9)
fmt.Println(c.Value()) // 10
c.Reset()
```

Waiting for thresholds:

```go
// Wait until the counter is at least 10 or time out after 1s.
if err := <-c.WaitAbove(10, time.Second); err != nil {
    // err is counter.ErrTimeout on timeout
    log.Fatal(err)
}

// Wait until the counter drops back to 0.
<-c.WaitBelow(0, time.Second)
```

Behavior notes:

- All operations are atomic on an `int64`.
- Waiters poll at a small fixed interval (10ms) to avoid busy spinning.
- WaitAbove/WaitBelow send a single value on their returned channel: `nil` on
  success, or `ErrTimeout` on timeout. The goroutine exits after sending.
