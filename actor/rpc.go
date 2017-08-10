package actor

import "sync/atomic"

var rpcSeq int64

func AllocRPCSeq() int64 {
	atomic.AddInt64(&rpcSeq, 1)
	return rpcSeq
}
