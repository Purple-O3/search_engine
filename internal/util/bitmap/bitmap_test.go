package bitmap

import "testing"

func TestAll(t *testing.T) {
	bm := NewBitmap()
	bm.Add(10)
	ret := bm.IsExist(10)
	t.Log(ret)
	bm.Add(888888888)
	ret = bm.IsExist(10)
	t.Log(ret)
	ret = bm.IsExist(888888888)
	t.Log(ret)
	bm.Add(2305843009213693952)
	ret = bm.IsExist(888888888)
	t.Log(ret)
	ret = bm.IsExist(9999)
	t.Log(ret)
}
