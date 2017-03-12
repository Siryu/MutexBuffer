package MutexBuffer

import (
	"time"
    "sync"
	"log"
	"fmt"
)

type MutexList struct {
    Mutexes []bool
    Lock    *sync.Mutex
}

func (m *MutexList) SetBuffer(bufferSize int) {
    m.Lock := new(sync.Mutex)
    m.Lock.Lock()
    m.Mutexes = make([]bool, bufferSize)
    m.Lock.Unlock()
    return
}

func (m *MutexList) Lock() {
    m.Lock.Lock()
    found := false
    for {
        for key, locked := range m.Mutexes {
            if !locked {
                m.Mutexes[key] = true
                found = true
                break
            }
        } 
        if !found {
            time.Sleep(5 * time.Millisecond)
        } else {
            break
        }
    }
    m.Lock.Unlock()
    return
}

func (m *MutexList) Unlock()(err error) {
    m.Lock.Lock()
    found := false
    for key, locked := range m.Mutexes {
        if locked {
            m.Mutexes[key] = false
            found = true
            break
        }
    }
    if !found {
        err = fmt.Errorf("No Mutexes assigned to Unlock")
    } 
    m.Lock.Unlock()
    return
}