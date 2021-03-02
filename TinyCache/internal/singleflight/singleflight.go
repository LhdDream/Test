package singleflight

import "sync"

type call struct {
	 wg sync.WaitGroup
	 val interface{}
	 err error
}

type Group struct {
	m sync.Map
}

func (g * Group) Do(key string, fn func()(interface{},error)) (interface{},error) {

	if c, ok := g.m.Load(key) ; ok {
		value := c.(*call)
		value.wg.Wait()
		return value.val , value.err
	}
	value := new(call)
	value.wg.Add(1)
	g.m.Store(key,value)

	value.val, value.err = fn()
	value.wg.Done()

	g.m.Delete(key)
	return  value.val,value.err
}
