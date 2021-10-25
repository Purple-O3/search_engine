package index

import (
	"encoding/json"
	"search_engine/internal/model/store"
	"search_engine/internal/service/objs"
	"search_engine/internal/util/log"
	"search_engine/internal/util/tools"
	"sync"
	"time"
)

type InvertedIndex struct {
	termDict [2]map[string]objs.PostingList
	use      int
	db       store.Store
	quit     chan quitGo
	rwlock   sync.RWMutex
}

func NewInvertedIndex(db store.Store) *InvertedIndex {
	ii := new(InvertedIndex)
	ii.termDict[0] = make(map[string]objs.PostingList)
	ii.termDict[1] = nil
	ii.use = 0
	ii.db = db
	go func() {
		now := time.Now().Second()
		lastTime := now
		for {
			select {
			case q := <-ii.quit:
				ii.flushDB()
				q.done <- struct{}{}
				return
			default:
				now = time.Now().Second()
				if ii.Len() > 10000 || now-lastTime > 1 {
					ii.flushDB()
					lastTime = now
				} else {
					time.Sleep(1 * time.Millisecond)
				}
			}
		}
	}()
	return ii
}

type quitGo struct {
	done chan struct{}
}

func (ii *InvertedIndex) FlushAll() {
	q := quitGo{}
	q.done = make(chan struct{})
	ii.quit <- q
	<-q.done
}

func (ii *InvertedIndex) Len() int {
	ii.rwlock.RLock()
	defer ii.rwlock.RUnlock()
	return len(ii.termDict[ii.use])
}

func (ii *InvertedIndex) Set(term string, posting objs.Posting) {
	log.Debugf("key:%v, value:%v", term, posting)
	ii.rwlock.Lock()
	defer ii.rwlock.Unlock()
	if _, ok := ii.termDict[ii.use][term]; !ok {
		ii.termDict[ii.use][term] = make(objs.PostingList, 0)
	}
	ii.termDict[ii.use][term] = append(ii.termDict[ii.use][term], posting)
}

func (ii *InvertedIndex) Get(term string) (objs.PostingList, bool) {
	ii.rwlock.RLock()
	defer ii.rwlock.RUnlock()
	pl, ok := ii.termDict[ii.use][term]
	log.Debugf("key:%v, len_value:%v", term, len(pl))
	return pl, ok
}

func (ii *InvertedIndex) flushDB() {
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%d, cost: %.3f ms", 0, float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	ii.rwlock.RLock()
	dict := ii.termDict[ii.use]
	free := 1 - ii.use
	ii.termDict[free] = make(map[string]objs.PostingList)
	ii.termDict[ii.use] = nil
	ii.use = free
	ii.rwlock.RUnlock()
	for k, v := range dict {
		key := tools.Str2Bytes(k)
		valueStored, err := ii.db.Get(key)
		if err == nil {
			pl := make(objs.PostingList, 0)
			err = json.Unmarshal(valueStored, &pl)
			if err == nil {
				v = append(v, pl...)
			}
		}
		value, err := json.Marshal(&v)
		if err == nil {
			ii.db.Set(key, value)
		}
	}
}
