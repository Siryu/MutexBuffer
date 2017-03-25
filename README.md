# MutexBuffer
Allow for a single mutex to have a buffer that allows for more than one lock

This allows for a certain number of go routines to access the code at a time

## How to Use
```go
package main

import (
  "github.com/Siryu/MutexBuffer"
)

func main() {
// create reference to mutex
// set the size you want the buffer to be
mutex := mutexBuffer.New(3)
// lock a mutex
mutex.Lock()

// insert the code here you want to be restricted to

// unlock a mutex
mutex.Unlock()
}
```

## License

MIT
