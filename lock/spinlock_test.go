package lock

import (
	"sync"
	"testing"
	"time"
)

func TestLock(t *testing.T) {
	var sl sync.Locker
	sl = NewSpinlock()
	for i := 0; i < 5; i++ {
		go func(i int) {
			sl.Lock()
			defer sl.Unlock()

			t.Logf("this is the %d goroutine", i)
		}(i)
	}

	time.Sleep(2)
}
