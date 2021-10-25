package ringqueue

import (
	"errors"
	"sync"
)

type RingQueue struct {
	queue []interface{}
	free  int
	used  int
	len   int
	lock  sync.Mutex
}

func NewRingQueue(len int) *RingQueue {
	rq := new(RingQueue)
	rq.queue = make([]interface{}, len)
	rq.len = len
	rq.free = 0
	rq.used = 0
	return rq
}

func (rq *RingQueue) Set(data interface{}) error {
	rq.lock.Lock()
	defer rq.lock.Unlock()
	if rq.full() {
		return errors.New("ring_queue is full")
	}
	rq.queue[rq.free] = data
	rq.free = (rq.free + 1) % rq.len
	return nil
}

func (rq *RingQueue) Get() (interface{}, error) {
	rq.lock.Lock()
	defer rq.lock.Unlock()
	if rq.empty() {
		return nil, errors.New("ring_queue is free")
	}
	data := rq.queue[rq.used]
	rq.used = (rq.used + 1) % rq.len
	return data, nil
}

func (rq *RingQueue) empty() bool {
	if rq.used == rq.free {
		return true
	} else {
		return false
	}
}

func (rq *RingQueue) full() bool {
	if rq.used == (rq.free+1)%rq.len {
		return true
	} else {
		return false
	}
}
