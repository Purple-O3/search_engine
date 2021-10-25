package hashset

import "sync"

type Set struct {
	items  map[interface{}]struct{}
	rwlock sync.RWMutex
}

func NewSet() *Set {
	set := &Set{items: make(map[interface{}]struct{})}
	return set
}

func (s *Set) Add(item interface{}) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	s.items[item] = struct{}{}
}

func (s *Set) Del(item interface{}) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	delete(s.items, item)
}

func (s *Set) Contains(item interface{}) bool {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	if _, ok := s.items[item]; !ok {
		return false
	} else {
		return true
	}
}

func (s *Set) Size() int {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	return len(s.items)
}

func (s *Set) GetAll() []interface{} {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()
	values := make([]interface{}, s.Size())
	count := 0
	for item := range s.items {
		values[count] = item
		count++
	}
	return values
}

func (s *Set) Clear() {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	s.items = make(map[interface{}]struct{})
}
