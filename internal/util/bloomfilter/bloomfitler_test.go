package bloomfilter

import (
	"math"
	"testing"
)

func Test(t *testing.T) {
	nub := math.Log(2) * math.Log(2)
	t.Log(nub)
	bf := NewBloomFilter(0.00001, 100)
	var docid uint64
	docid = 123
	bf.AddNub(docid)
	ret := bf.CheckNub(docid)
	t.Log(ret)
	docid = 120
	bf.AddNub(docid)
	ret = bf.CheckNub(docid)
	t.Log(ret)
	docid = 124
	ret = bf.CheckNub(docid)
	t.Log(ret)
	bf = NewBloomFilter(0.00001, 100000000)
	mbSize := bf.Size() / 8 / 1000 / 1000
	t.Log("MbSize:", mbSize)
	docid = 125
	bf.AddNub(docid)
	ret = bf.CheckNub(docid)
	t.Log(ret)
}
