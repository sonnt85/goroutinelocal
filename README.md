# goroutinelocal

[![Go Reference](https://pkg.go.dev/badge/github.com/sonnt85/goroutinelocal.svg)](https://pkg.go.dev/github.com/sonnt85/goroutinelocal)

Generic goroutine-local storage for Go — per-goroutine values isolated by goroutine ID.

## Installation

```bash
go get github.com/sonnt85/goroutinelocal
```

## Features

- Generic `goroutineLocal[T]` — stores a value of any type per goroutine
- Optional initializer function called when a goroutine has no stored value
- `Get` / `Set` / `Remove` per the current goroutine's ID
- `GetMap` to retrieve all stored values across all goroutines

## Usage

```go
// Create goroutine-local storage with a default initializer
local := goroutinelocal.NewGoroutineLocal(func() int { return 0 })

// In goroutine A
local.Set(42)
fmt.Println(local.Get()) // 42

// In goroutine B (different goroutine ID)
fmt.Println(local.Get()) // 0 (from initializer)

// Iterate all stored values
m := local.GetMap() // map[goroutineID]value
```

## API

- `NewGoroutineLocal[T](initfun func() T) *goroutineLocal[T]` — create a new goroutine-local store; `initfun` may be nil
- `Get() T` — return the value for the current goroutine (calls `initfun` if not set)
- `Set(v T)` — store a value for the current goroutine
- `Remove()` — delete the value for the current goroutine
- `GetMap() map[int64]T` — return a snapshot of all goroutine values

## Author

**sonnt85** — [thanhson.rf@gmail.com](mailto:thanhson.rf@gmail.com)

## License

MIT License - see [LICENSE](LICENSE) for details.
