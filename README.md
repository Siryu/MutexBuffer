# MutexBuffer
Allow for a single mutex to have a buffer that allows for more than one lock


## How to Use
```go
package main

import (
  "github.com/Siryu/MutexBuffer"
)

func main() {
// create reference to mutex
mutex := &mutexBuffer.MutexBuffer{}
// set the size you want the buffer to be
mutex.SetBuffer(3)
// lock a mutex
mutex.Lock()
// unlock a mutex
mutex.Unlock()
```

## License

MIT
