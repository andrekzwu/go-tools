package pool

import (
	"fmt"
	"testing"
	"time"
)

// PrintNums
func PrintNums(i int) {
	fmt.Println(i)
}

// TestThreadPools
func TestThreadPool(t *testing.T) {
	threadPool, err := CreateThreadPool(20)
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}
	// set queue hanle interval
	threadPool.SetQueueInterval(500 * time.Millisecond)
	for i := 0; i < 100; i++ {
		n := i
		threadPool.GetFreeThreadInfo(func() {
			PrintNums(n)
		})
	}
	select {
	case <-time.After(10 * time.Second):
	}
	threadPool.DistoryPool()
}
