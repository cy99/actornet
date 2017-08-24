package util

import (
	"sync/atomic"
	"time"
)

var (
	globalSeq int64
)

func getSeq() int64 {

	return atomic.AddInt64(&globalSeq, 1)
}

// 生成可持久化的ID
func GenPersistantID(section int32) int64 {

	seq := getSeq() & 0xFFF

	time := (time.Now().Unix() & 0xFFFFFFFF) << 12

	part := (int64(section) & 0xFFFF) << 44

	return part | time | seq
}
