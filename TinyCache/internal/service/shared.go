package service

//这里基于数据分片和lru算法的结合进行实现

import (
	"container/list"
	"sync"
)

type entry struct {
	key string
	value interface{}
}

type cache struct {
	mutex sync.RWMutex

	maxEntries int // 这块cache 存放的最多条数
	onEvicted func(key string, value interface{})

	ll *list.List
	cache map[string]*list.Element
}

func newCache(maxEntries int, onEvicted func(key string, value interface{})) *cache {
	return &cache{
		maxEntries: maxEntries,
		onEvicted: onEvicted,
		ll: list.New(),
		cache: make(map[string]*list.Element),
	}
}

func (m *cache) remove(e *list.Element) {
	if e == nil {
		return
	}
	m.ll.Remove(e)
	val := e.Value.(*entry)
	delete(m.cache,val.key)

	if m.onEvicted != nil {
		m.onEvicted(val.key,val.value)
	}
}

func (m *cache) set(key string, value interface{}) {
	m.mutex.Lock()
	defer  m.mutex.Unlock()
	//放置首部
	if val, ok := m.cache[key] ; ok {
		m.ll.MoveToFront(val)
		en := val.Value.(*entry)
		en.value = value
		return
	}

	en := &entry{key,value}

	val := m.ll.PushFront(en)
	m.cache[key] = val

	if len(m.cache) > m.maxEntries {
		m.remove(m.ll.Back()) // 删除末尾元素
	}
}

func (m *cache) get(key string) interface{} {
	m.mutex.Lock()
	defer  m.mutex.Unlock()

	if val,ok := m.cache[key]; ok {
		m.ll.MoveToFront(val)
		return val.Value.(*entry).value
	}
	return nil
}

func (m *cache) del(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if val,ok := m.cache[key]; ok {
		delete(m.cache,key)
		m.ll.Remove(val)
	}
}

func (m *cache) len() int {
	m.mutex.RLock()
	defer  m.mutex.RUnlock()
	return m.ll.Len()
}