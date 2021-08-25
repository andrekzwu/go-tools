package pool

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	MAX_THREAD_NUMS = 100
)

var ErrMaxThreadLimit = errors.New("input thread nums more than Max nums")

type HandleEvent func()

// ThreadInfo
type ThreadInfo struct {
	f   HandleEvent
	sem chan int
	no  int
}

// ThreadPool
type ThreadPool struct {
	threadnums    int
	freeChans     chan int
	mx            sync.Mutex
	threadInfos   []*ThreadInfo
	queuesmx      sync.Mutex
	queues        []HandleEvent
	queueInterval time.Duration // queue handle interval default 5 millsecond
	cancel        context.CancelFunc
	isover        bool
}

// CreateThreadPool
func CreateThreadPool(threadnums int) (*ThreadPool, error) {
	if threadnums > MAX_THREAD_NUMS {
		return nil, ErrMaxThreadLimit
	}
	//
	ctx, cancel := context.WithCancel(context.Background())
	//
	pool := &ThreadPool{
		threadnums:    threadnums,
		freeChans:     make(chan int, threadnums),
		mx:            sync.Mutex{},
		queuesmx:      sync.Mutex{},
		threadInfos:   make([]*ThreadInfo, 0),
		queues:        make([]HandleEvent, 0),
		queueInterval: time.Millisecond * 5,
		cancel:        cancel,
		isover:        false,
	}
	for i := 0; i < threadnums; i++ {
		pool.threadInfos = append(pool.threadInfos, &ThreadInfo{
			sem: make(chan int),
			no:  i,
		})
		go pool.threadFunc(pool.threadInfos[i])
	}
	//
	go pool.queueMonitor(ctx)
	return pool, nil
}

// DistoryPool
func (tp *ThreadPool) DistoryPool() {
	tp.isover = true
	// stop monitor
	tp.cancel()
	// close freechans
	close(tp.freeChans)
	// close thread sem
	for _, v := range tp.threadInfos {
		close(v.sem)
	}
	tp.threadInfos = nil
	// queue
	tp.queues = nil
	//
	return
}

// queueMonitor
func (tp *ThreadPool) queueMonitor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.Tick(tp.queueInterval):
		}
		//
		if len(tp.queues) == 0 {
			continue
		}
		tp.queuesmx.Lock()
		handle := tp.queues[len(tp.queues)-1]
		tp.queues = tp.queues[:len(tp.queues)-1]
		tp.queuesmx.Unlock()
		//
		tp.dispachTask2Thread(handle)
	}
}

// GetFreeThreadInfo
func (tp *ThreadPool) GetFreeThreadInfo(handle HandleEvent) {
	if tp == nil || tp.isover {
		return
	}
	if len(tp.freeChans) == tp.threadnums {
		tp.queuesmx.Lock()
		tp.queues = append(tp.queues, handle)
		tp.queuesmx.Unlock()
		return
	}
	tp.dispachTask2Thread(handle)
}

// SetQueueInterval
func (tp *ThreadPool) SetQueueInterval(interval time.Duration) {
	tp.queueInterval = interval
}

// dispachFreeThread
func (tp *ThreadPool) dispachTask2Thread(handle HandleEvent) {
	if tp.isover {
		return
	}
	// free sem
	tp.freeChans <- 0
	//
	if len(tp.threadInfos) == 0 {
		return
	}
	//
	tp.mx.Lock()
	threadInfo := tp.threadInfos[len(tp.threadInfos)-1]
	tp.threadInfos = tp.threadInfos[:len(tp.threadInfos)-1]
	tp.mx.Unlock()
	//
	threadInfo.f = handle
	threadInfo.sem <- 0
	//
}

// threadFunc
func (tp *ThreadPool) threadFunc(threadInfo *ThreadInfo) {
	for {
		<-threadInfo.sem
		// call handle
		if threadInfo.f == nil {
			continue
		}
		threadInfo.f()
		//
		threadInfo.f = nil
		//
		tp.mx.Lock()
		tp.threadInfos = append(tp.threadInfos, threadInfo)
		tp.mx.Unlock()
		//
		<-tp.freeChans
	}
}
