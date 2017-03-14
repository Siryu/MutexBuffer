package mutexBuffer

import "testing"

// TestLock tests to make sure the mutexes are being locked
func TestLock(t *testing.T) {
	amountOfLocks := 2
	bufferSize := 10
	mutex := &MutexBuffer{}
	mutex.SetBuffer(bufferSize)
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
	mutex := &MutexBuffer{}
	mutex.SetBuffer(bufferSize)
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
