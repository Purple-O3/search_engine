package index

import (
	"search_engine/internal/model/store"
	"search_engine/internal/util/log"
	"search_engine/internal/util/tools"
	"sync"
	"time"
)

type PositiveIndex struct {
	docDict [2]map[string]string
	use     int
	db      store.Store
	quit    chan quitGo
	rwlock  sync.RWMutex
}

func NewPositiveIndex(db store.Store) *PositiveIndex {
	pi := new(PositiveIndex)
	pi.docDict[0] = make(map[string]string)
	pi.docDict[1] = nil
	pi.use = 0
	pi.db = db
	pi.quit = make(chan quitGo, 1)
	go func() {
		now := time.Now().Unix()
		lastTime := now
		for {
			func() {
				defer func() {
					if err := recover(); err != nil {
						log.Errorf("%v", err)
					}
				}()

				select {
				case q := <-pi.quit:
					pi.flushDB()
					q.done <- struct{}{}
					return
				default:
					now = time.Now().Unix()
					if pi.Len() > 10000 || now-lastTime >= 1 {
						pi.flushDB()
						lastTime = now
					} else {
						time.Sleep(1 * time.Millisecond)
					}
				}
			}()
		}
	}()
	return pi
}

func (pi *PositiveIndex) FlushAll() {
	q := quitGo{}
	q.done = make(chan struct{}, 1)
	pi.quit <- q
	<-q.done
}

func (pi *PositiveIndex) Len() int {
	pi.rwlock.RLock()
	defer pi.rwlock.RUnlock()
	return len(pi.docDict[pi.use])
}

func (pi *PositiveIndex) Set(key string, value string) {
	log.Debugf("key:%v, value:%v", key, value)
	pi.rwlock.Lock()
	defer pi.rwlock.Unlock()
	pi.docDict[pi.use][key] = value
}

func (pi *PositiveIndex) Get(key string) (string, bool) {
	pi.rwlock.RLock()
	defer pi.rwlock.RUnlock()
	value, ok := pi.docDict[pi.use][key]
	log.Debugf("key:%v, value:%v", key, value)
	return value, ok
}

func (pi *PositiveIndex) flushDB() {
	defer func(cost func() time.Duration) {
		t := cost().Microseconds()
		if t > 1000 {
			log.Warnf("trackid:%v, cost: %.3f ms", 0, float64(t)/1000.0)
		}
	}(tools.TimeCost())

	pi.rwlock.Lock()
	dict := pi.docDict[pi.use]
	free := 1 - pi.use
	pi.docDict[free] = make(map[string]string)
	pi.docDict[pi.use] = nil
	pi.use = free
	pi.rwlock.Unlock()
	for k, v := range dict {
		key := tools.Str2Bytes(k)
		value := tools.Str2Bytes(v)
		pi.db.Set(key, value)
	}
}
