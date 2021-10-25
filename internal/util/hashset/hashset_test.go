package hashset

import "testing"

func TestAll(t *testing.T) {
	set := NewSet()
	set.Add(8)
	set.Add(10)
	len := set.Size()
	t.Log("len:", len)
	set.Del(8)
	set.Del(9)
	v := set.GetAll()
	t.Log("all:", v)
}
