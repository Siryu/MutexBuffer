package mutexBuffer

import (
	"log"
	"time"
)

// example using MutexBuffer with times and outputs in order to see how this behaves.
// runs through locks to Enable printing out the times for when each go routine is unlocking
// watch how it will only allow the amount of routines through that you specified.
func main() {
	// create reference to a MutexBuffer and set the size of the buffer
	mutex := New(3)
	locksToEnable := 100

	startTime := time.Now()
	done := make(chan bool, locksToEnable)
	finished := make(chan bool)

	for i := 0; i < locksToEnable; i++ {
		go func() {
			// start with the lock
			mutex.Lock()

			// code for the routine
			time.Sleep(1000 * time.Millisecond)
			log.Printf("routine %v finished sleeping after %v", i, time.Since(startTime))

			// unlock when you no longer need a lock
			err := mutex.Unlock()

			done <- true
			if err != nil {
				log.Printf("Got an erro: %v", err.Error())
			}
		}()
	}

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
