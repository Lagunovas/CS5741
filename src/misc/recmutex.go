package recmutex

import (
	"sync"
	"time"

	gls "github.com/v2pro/plz/gls"
)

type RecursiveMutex struct {
	internalMutex    sync.Mutex
	currentGoRoutine int64
	lockCount        uint64
}

func (rm *RecursiveMutex) Lock() {
	goRoutineID := gls.GoID()

	for {
		rm.internalMutex.Lock()
		if rm.currentGoRoutine == 0 {
			rm.currentGoRoutine = goRoutineID
			break
		} else if rm.currentGoRoutine == goRoutineID {
			break
		} else {
			rm.internalMutex.Unlock()
			time.Sleep(time.Millisecond)
			continue
		}
	}
	rm.lockCount++
	rm.internalMutex.Unlock()
}

func (rm *RecursiveMutex) Unlock() {
	rm.internalMutex.Lock()
	rm.lockCount--
	if rm.lockCount == 0 {
		rm.currentGoRoutine = 0
	}
	rm.internalMutex.Unlock()
}
