package scheduler

import (
	"sync"
	"time"
)

// Bucket flow control
type Bucket struct {
	updateInterval uint
	rate           float32
	burst          float32
	avaliableToken float32
	lock           *sync.Mutex
	lastUpdate     float32
}

// NewBucket is defalut constructor
func NewBucket(rateP float32, burstP float32) *Bucket {
	now := float32(time.Now().Second())
	return &Bucket{
		// Refresh interval
		updateInterval: 30,
		// Regenerate token rate
		rate: rateP,
		// Limit size of bucket
		burst: burstP,
		// Lock
		lock: new(sync.Mutex),
		// Avaliable tokens in bucket now
		avaliableToken: burstP,
		// The last update time
		lastUpdate: now,
	}
}

// Get the avaliable token in Bucket
func (buc *Bucket) Get() float32 {
	now := float32(time.Now().Second())
	if buc.avaliableToken >= buc.burst {
		buc.lastUpdate = now
		return buc.avaliableToken
	}

	// Generate now token
	bucket := buc.rate * (now - buc.lastUpdate)
	buc.lock.Lock()
	if bucket > 1 {
		buc.avaliableToken += bucket
		if buc.avaliableToken > buc.burst {
			buc.avaliableToken = buc.burst
		}
		buc.lastUpdate = now
	}
	buc.lock.Unlock()

	// Give avaliable token
	return buc.avaliableToken
}

// Set numbers of token in bucket
func (buc *Bucket) Set(size float32) {
	buc.avaliableToken = size
}

// Desc to decrease the avaliable token
func (buc *Bucket) Desc(n float32) {
	buc.avaliableToken -= n
}
