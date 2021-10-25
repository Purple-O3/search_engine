package bloomfilter

import (
	"hash"
	"hash/fnv"
	"math"
	"search_engine/internal/util/tools"
	"strconv"
	"sync"
)

type BloomFilter struct {
	bitmap  []byte
	used    uint64 //n
	size    uint64 //m
	hashFn  hash.Hash64
	hashCnt uint64 //k
	rwlock  sync.RWMutex
}

func NewBloomFilter(miscalRate float64, addSize uint64) *BloomFilter {
	bf := new(BloomFilter)
	bf.used = 0
	bf.size = uint64(math.Ceil(float64(addSize) * math.Log(miscalRate) / (-0.48))) //m=n*ln(p)/-0.48, math.Log() equals ln()
	bf.bitmap = make([]byte, bf.size/8+1)
	bf.hashFn = fnv.New64()
	bf.hashCnt = uint64(math.Ceil(0.7 * float64(bf.size/addSize))) //k=0.7*m/n
	return bf
}

func (bf *BloomFilter) getHash(Byte []byte) (uint64, uint64) {
	bf.hashFn.Reset()
	bf.hashFn.Write(Byte)
	h1 := bf.hashFn.Sum64()

	h1String := strconv.FormatUint(h1, 10)
	h1Byte := tools.Str2Bytes(h1String)
	bf.hashFn.Reset()
	bf.hashFn.Write(h1Byte)
	h2 := bf.hashFn.Sum64()
	return h1, h2
}

func (bf *BloomFilter) Add(Byte []byte) {
	var i uint64
	h1, h2 := bf.getHash(Byte)
	bf.rwlock.Lock()
	defer bf.rwlock.Unlock()
	for i = 0; i < bf.hashCnt; i++ {
		nub := (h1 + i*h2) % bf.size
		byteIndex, bitIndex := nub/8, nub%8
		bf.bitmap[byteIndex] |= 0x01 << bitIndex
	}
	bf.used++
}

func (bf *BloomFilter) AddNub(id uint64) {
	idString := strconv.FormatUint(id, 10)
	idByte := tools.Str2Bytes(idString)
	bf.Add(idByte)
}

func (bf *BloomFilter) Check(Byte []byte) bool {
	var i uint64
	h1, h2 := bf.getHash(Byte)
	bf.rwlock.RLock()
	defer bf.rwlock.RUnlock()
	for i = 0; i < bf.hashCnt; i++ {
		nub := (h1 + i*h2) % bf.size
		byteIndex, bitIndex := nub/8, nub%8
		if bf.bitmap[byteIndex]&(0x01<<bitIndex) == 0 {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) CheckNub(id uint64) bool {
	idString := strconv.FormatUint(id, 10)
	idByte := tools.Str2Bytes(idString)
	return bf.Check(idByte)
}

func (bf *BloomFilter) Used() uint64 {
	return bf.used
}

func (bf *BloomFilter) Size() uint64 {
	return bf.size
}

//Pow(x, y float64) float64  // x 的幂函数
//Exp(x float64) float64 // x的base-e指数函数
func (bf *BloomFilter) FalsePositiveRate() float64 {
	return math.Pow((1 - math.Exp(-float64(bf.used*bf.hashCnt)/float64(bf.size))), float64(bf.hashCnt))
}
