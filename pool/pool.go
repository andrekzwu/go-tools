// Copyright 2017, personal.andre. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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

// HandleEvent 处理事件
type HandleEvent func()

// ThreadInfo 协程结构定义
type ThreadInfo struct {
	f   HandleEvent
	sem chan int
	no  int
}

// ThreadPool 协程池定义
type ThreadPool struct {
	threadnums    int                // 携程数量
	threadChans   chan int           // 协程控制器
	mx            sync.Mutex         // 协程保护锁
	threadInfos   []*ThreadInfo      // 携程信息
	queuesmx      sync.Mutex         // 事件队列锁
	queues        []HandleEvent      // 事件队列
	queueInterval time.Duration      // queue handle interval default 5 millsecond
	cancel        context.CancelFunc // 协程控制上下文函数
	isover        bool               // 协程池是否结束
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
		threadChans:   make(chan int, threadnums),
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
	// close threadchans
	close(tp.threadChans)
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
		handle := tp.queues[0]
		tp.queues = tp.queues[1:]
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
	if len(tp.threadChans) == tp.threadnums {
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
	tp.threadChans <- 0
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
		<-tp.threadChans
	}
}
