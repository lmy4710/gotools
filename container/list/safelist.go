package list

import (
	"container/list"
	"sync"
)

type SafeList struct {
	sync.RWMutex
	L *list.List
}

func NewSafeList() *SafeList {
	return &SafeList{L: list.New()}
}

func (this *SafeList) PushFront(v interface{}) {
	this.Lock()
	this.L.PushFront(v)
	this.Unlock()
}

func (this *SafeList) PopBack() interface{} {
	this.Lock()

	if elem := this.L.Back(); elem != nil {
		item := this.L.Remove(elem)
		this.Unlock()
		return item
	}

	this.Unlock()
	return nil
}

func (this *SafeList) PopBackBy(max int) []interface{} {
	this.Lock()

	count := this.len()
	if count == 0 {
		this.Unlock()
		return []interface{}{}
	}

	if count > max {
		count = max
	}

	items := make([]interface{}, 0, count)
	for i := 0; i < count; i++ {
		item := this.L.Remove(this.L.Back())
		items = append(items, item)
	}

	this.Unlock()
	return items
}

func (this *SafeList) RemoveAll() {
	this.Lock()
	this.L = list.New()
	this.Unlock()
}

func (this *SafeList) Front() interface{} {
	this.RLock()

	if f := this.L.Front(); f != nil {
		this.RUnlock()
		return f.Value
	}

	this.RUnlock()
	return nil
}

func (this *SafeList) Len() int {
	this.RLock()
	defer this.RUnlock()
	return this.len()
}

func (this *SafeList) len() int {
	return this.L.Len()
}

// SafeList with Limited Size
type SafeListLimited struct {
	maxSize int
	SL      *SafeList
}

func NewSafeListLimited(maxSize int) *SafeListLimited {
	return &SafeListLimited{SL: NewSafeList(), maxSize: maxSize}
}

func (this *SafeListLimited) PopBack() interface{} {
	return this.SL.PopBack()
}

func (this *SafeListLimited) PopBackBy(max int) []interface{} {
	return this.SL.PopBackBy(max)
}

func (this *SafeListLimited) PushFront(v interface{}) bool {
	if this.SL.Len() >= this.maxSize {
		return false
	}

	this.SL.PushFront(v)
	return true
}

func (this *SafeListLimited) RemoveAll() {
	this.SL.RemoveAll()
}

func (this *SafeListLimited) Front() interface{} {
	return this.SL.Front()
}

func (this *SafeListLimited) Len() int {
	return this.SL.Len()
}
