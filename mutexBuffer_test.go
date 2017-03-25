package mutexBuffer

import (
	"strconv"
	"testing"
	"time"
)

// TestLock tests to make sure the mutexes are being locked
func TestLock(t *testing.T) {
	amountOfLocks := 2
	bufferSize := 10
	mutex := New(bufferSize)
	for i := 0; i < amountOfLocks; i++ {
		mutex.Lock()
	}
	amountLocked, _ := setLocks(mutex)
	if amountLocked != amountOfLocks {
		t.Errorf("The amount of locks set: %v, does not equal the actual: %v", amountOfLocks, amountLocked)
	}
}

// TestUnlock test to make sure the mutexes are being released
func TestUnlock(t *testing.T) {
	amountOfLocks := 6
	amountOfUnlocks := 2
	bufferSize := 10
	mutex := New(bufferSize)
	for i := 0; i < amountOfLocks; i++ {
		mutex.Lock()
	}
	for i := 0; i < amountOfUnlocks; i++ {
		mutex.Unlock()
	}
	_, amountUnlocked := setLocks(mutex)
	if amountUnlocked != bufferSize-amountOfLocks+amountOfUnlocks {
		t.Errorf("The amount of unlocks: %v, does not set the locks correctly to: %v", amountUnlocked, bufferSize-amountOfLocks+amountOfUnlocks)
	}
}

// TestLockLocking tests that the lock is locking the go routine
func TestLockLocking(t *testing.T) {
	testListSize := 5
	bufferSize := 2
	times := make(chan time.Duration, testListSize)
	mutex := New(bufferSize)

	// start go routines that sleep for a time, in order to check times upon completion
	for i := 0; i < testListSize; i++ {
		startTime := time.Now().Round(time.Millisecond)
		go func() {
			mutex.Lock()
			time.Sleep(100 * time.Millisecond)
			mutex.Unlock()
			times <- time.Since(startTime)
		}()
	}

	// pulls out times to make sure that the mutex is locking when it's supposed to
	<-times
	<-times
	timeFirst := <-times
	timeFirstCheck := convertToWholeMS(timeFirst)
	<-times
	timeLast := <-times
	timeLastCheck := convertToWholeMS(timeLast)
	if timeFirstCheck < 199 || timeFirstCheck > 300 {
		t.Errorf("time for getting through the mutex doesn't match got: %v, expected: between %v and %v", timeFirst, 200*time.Millisecond, 300*time.Millisecond)
	}
	if timeLastCheck < 299 || timeLastCheck > 400 {
		t.Errorf("time for getting through the mutex doesn't match got: %v, expected: between %v and %v", timeLast, 300*time.Millisecond, 400*time.Millisecond)
	}
}

func convertToWholeMS(duration time.Duration) (durationFixed int) {
	timeToString := ((duration / time.Millisecond) * time.Millisecond).String()
	timeToString = timeToString[:3]
	durationFixed, _ = strconv.Atoi(timeToString)
	return
}

func setLocks(mutex *MutexBuffer) (amountLocked int, amountUnlocked int) {
	for _, value := range mutex.Mutexes {
		if value {
			amountLocked++
		} else {
			amountUnlocked++
		}
	}
	return
}
