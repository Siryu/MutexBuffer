package mutexBuffer

import (
	"fmt"
	"runtime"
	"sync"
)

// MutexBuffer create reference and hold to make a buffered mutex
type MutexBuffer struct {
	Mutexes  []bool
	RealLock *sync.Mutex
}

// SetBuffer takes in an int that determines the amount of locks
func (m *MutexBuffer) SetBuffer(bufferSize int) {
	m.RealLock = new(sync.Mutex)
	m.RealLock.Lock()
	m.Mutexes = make([]bool, bufferSize)
	m.RealLock.Unlock()
	return
}

// Lock sets a lock on one of the mutexes
func (m *MutexBuffer) Lock() {
	found := false
	for {
		m.RealLock.Lock()
		for key, locked := range m.Mutexes {
			if !locked {
				m.Mutexes[key] = true
				found = true
				break
			}
		}
		m.RealLock.Unlock()
		if !found {
			runtime.Gosched()
		} else {
			break
		}
	}
	return
}

// Unlock unlocks a lock from the mutexes
func (m *MutexBuffer) Unlock() (err error) {
	m.RealLock.Lock()
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
	m.RealLock.Unlock()
	return
}
