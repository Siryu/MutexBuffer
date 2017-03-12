package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	mutex := &MutexBuffer{}
	mutex.SetBuffer(10)

	locksToEnable := 100

	startTime := time.Now()
	done := make(chan bool, locksToEnable)
	finished := make(chan bool)

	go func(x int) {
		for i := 0; i < x; i++ {
			go func() {
				mutex.Lock()
				time.Sleep(1000 * time.Millisecond)
				log.Printf("routine %v finished sleeping after %v", i, time.Since(startTime))
				err := mutex.Unlock()
				done <- true
				if err != nil {
					log.Printf("Got an erro: %v", err.Error())
				}
			}()
		}
	}(locksToEnable)

	go func() {
		counter := 0
		for _ = range done {
			counter++
			if counter == locksToEnable {
				finished <- true
			}
		}
	}()

	<-finished
}

type MutexBuffer struct {
	Mutexes  []bool
	RealLock *sync.Mutex
}

func (m *MutexBuffer) SetBuffer(bufferSize int) {
	m.RealLock = new(sync.Mutex)
	m.RealLock.Lock()
	m.Mutexes = make([]bool, bufferSize)
	m.RealLock.Unlock()
	return
}

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
		if !found {
			m.RealLock.Unlock()

			runtime.Gosched()
		} else {
			m.RealLock.Unlock()
			break
		}
	}
	return
}

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
