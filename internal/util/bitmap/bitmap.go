//realize resize function
package bitmap

import (
	"fmt"
	"sync"
)

const defaultNub = 1000 //初始化默认大小

type Bitmap struct {
	data   [2][]byte
	use    int //data in use
	maxNub uint64
	rwlock sync.RWMutex
}

func NewBitmap() *Bitmap {
	return NewBitmapNub(defaultNub)
}

func NewBitmapNub(nub uint64) *Bitmap {
	bm := new(Bitmap)
	bm.data[0] = make([]byte, nub/8+1)
	bm.data[1] = nil
	bm.use = 0
	bm.maxNub = nub
	return bm
}

func (bm *Bitmap) resizeBitmap(nub uint64) {
	free := 1 - bm.use
	bm.data[free] = make([]byte, nub/8+1)
	copy(bm.data[free], bm.data[bm.use])
	bm.data[bm.use] = nil
	bm.use = free
	bm.maxNub = nub
}

func (bm *Bitmap) Add(nub uint64) {
	//catch err: makeslice: len out of range
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	bm.rwlock.Lock()
	defer bm.rwlock.Unlock()
	byteIndex, bitIndex := nub/8, nub%8
	if bm.maxNub < nub {
		bm.resizeBitmap(nub)
	}
	bm.data[bm.use][byteIndex] |= 0x01 << bitIndex
}

func (bm *Bitmap) IsExist(nub uint64) bool {
	bm.rwlock.RLock()
	defer bm.rwlock.RUnlock()
	byteIndex, bitIndex := nub/8, nub%8
	if bm.maxNub < nub {
		return false
	}
	if bm.data[bm.use][byteIndex]&(0x01<<bitIndex) == 0 {
		return false
	} else {
		return true
	}
}
