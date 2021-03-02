package service

type TinyCache struct {
	shared []*cache
	mask uint64 // 计算余数
	hash fnv64a  // 通过哈希值去进行计算
}

func NewTinyCache(maxEntries int , sharedNums int ,onEvicted func(key string, value interface{})) *TinyCache {
	tinyCache := &TinyCache{
		hash: newDefaultHash(),
		shared: make([]*cache,sharedNums),
		mask :uint64(sharedNums-1),
	}
	for i:= 0 ; i< sharedNums ;i++ {
		tinyCache.shared[i] = newCache(maxEntries,onEvicted)
	}
	return  tinyCache
}

func (m *TinyCache) GetShared(key string) *cache {
	return m.shared[m.hash.Sum64(key) & m.mask]
}

func (m *TinyCache) Get(key string)  interface{} {
	return m.GetShared(key).get(key)
}

func (m *TinyCache) Set(key string, value interface{}) {
	m.GetShared(key).set(key,value)
}

func (m *TinyCache) Del(key string) {
	m.GetShared(key).del(key)
}

