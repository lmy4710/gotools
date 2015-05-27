package nmap

import (
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	M map[string]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		M: make(map[string]interface{}),
	}
}

func (this *SafeMap) Put(key string, val interface{}) {
	this.Lock()
	this.M[key] = val
	this.Unlock()
}

func (this *SafeMap) Get(key string) (interface{}, bool) {
	this.Lock()
	val, exists := this.M[key]
	this.Unlock()
	return val, exists
}

func (this *SafeMap) Remove(key string) {
	this.Lock()
	delete(this.M, key)
	this.Unlock()
}

func (this *SafeMap) Clear() {
	this.Lock()
	this.M = make(map[string]interface{})
	this.Unlock()
}

func (this *SafeMap) ContainsKey(key string) bool {
	this.RLock()
	_, exists := this.M[key]
	this.RUnlock()
	return exists
}

func (this *SafeMap) Size() int {
	this.RLock()
	len := len(this.M)
	this.RUnlock()
	return len
}

func (this *SafeMap) IsEmpty() bool {
	this.RLock()
	empty := (len(this.M) == 0)
	this.RUnlock()
	return empty
}