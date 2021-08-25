// Copyright 2018, maizuo.andre. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package util

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/andrezhz/go-tools/redis"
)

const (
	MAX_UUID_LOCK_NUM = 3
	UUID_LOCK_TIME    = 3 * time.Second
	UUID_CREATE_LOCK  = "xxxxx:uuid:"
)

var (
	uuid = newUUID()
)

type UUID struct {
	lastTimestamp int64
	mutex         sync.Mutex
}

func newUUID() *UUID {
	uuid := &UUID{}
	uuid.lastTimestamp = -1
	uuid.mutex = sync.Mutex{}
	return uuid
}

// timeGen generate a unix millsecond
func timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// tilNextMillis
func tilNextMillis(lastTimestamp int64) int64 {
	timestamp := timeGen()
	for timestamp <= lastTimestamp {
		timestamp = timeGen()
	}
	return timestamp
}

func formatuuid(timestamp int64) string {
	return fmt.Sprintf("%d", timestamp)
}

func (u *UUID) uuid() string {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	timestamp := timeGen()
	if u.lastTimestamp == timestamp {
		timestamp = tilNextMillis(u.lastTimestamp)
	}
	u.lastTimestamp = timestamp
	return formatuuid(timestamp)
}

// getUUID
func getUUID() string {
	return uuid.uuid()
}

// EXPORT
//
// GenerateUUID
func GenerateUUID() (string, error) {
	for i := uint(0); i < MAX_UUID_LOCK_NUM; i++ {
		uuid := getUUID()
		lock := UUID_CREATE_LOCK + uuid
		if redis.RedisSetNx(lock, uuid, UUID_LOCK_TIME) {
			return uuid, nil
		}
	}
	return "", errors.New("generate uuid fail,please check the reids lock")
}
