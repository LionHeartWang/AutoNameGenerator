package internal

import (
	"sync"
)


type HashSet struct {
	Len int
	m   sync.Map
}

var value interface{}

func NewHashSet() *HashSet {
	set := &HashSet{
		Len: 0,
		m:   sync.Map{},
	}
	return set
}
func (hs *HashSet) Put(data interface{}) bool {
	if hs.IsExist(data) {
		return false
	}
	hs.m.Store(data, value)
	hs.Len++
	return true
}

func (hs *HashSet) Get() []interface{} {
	if hs.Len == 0 {
		return nil
	}
	res := make([]interface{}, 0, hs.Len)
	hs.m.Range(func(key, value interface{}) bool {
		res = append(res, key)
		return true
	})
	return res
}

func (hs *HashSet) Delete(data interface{}) bool {
	if !hs.IsExist(data) {
		return false
	}
	hs.m.Delete(data)
	hs.Len--
	return true
}

func (hs *HashSet) Clear() {
	hs.Len = 0
	hs.m = sync.Map{}
}

func (hs *HashSet) IsExist(data interface{}) bool {
	_, ok := hs.m.Load(data)
	return ok

}
