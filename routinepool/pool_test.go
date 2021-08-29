package routinepool

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n

}

func genfunc(i int) Task {
	return func() error {
		fmt.Printf("this is the %dth task, the routine is %d\n", i, getGID())
		<-time.After(time.Second * 1)
		return nil
	}
}

func TestRoutinePool(t *testing.T) {
	pool := NewRoutinePool(10)
	for i := 0; i < 60; i++ {
		f := genfunc(i)
		err := pool.Submit(f)
		if err != nil {
			logrus.Errorf("submit error, err=%v", err)
		}
	}
	<-time.After(time.Minute)
}
