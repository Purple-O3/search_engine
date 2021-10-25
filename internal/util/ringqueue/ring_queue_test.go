package ringqueue

import "testing"

func TestAll(t *testing.T) {
	rq := NewRingQueue(3)
	err := rq.Set(6)
	t.Log(err)
	err = rq.Set(5)
	t.Log(err)
	err = rq.Set(10)
	t.Log(err)
	err = rq.Set(8)
	t.Log(err)
	data, err := rq.Get()
	t.Log(data, err)
	data, err = rq.Get()
	t.Log(data, err)
	data, err = rq.Get()
	t.Log(data, err)
}
