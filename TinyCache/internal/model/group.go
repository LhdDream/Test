package model

import (
	"errors"
	"sync"
	"tinycache/internal/service"
)

type Getter interface {
	Get(key string) (string,error)
}

type GetterFunc func(key string) (string, error)

func (m GetterFunc) Get(key string) (string, error) {
	return m(key)
}
//从远处获取数据

type Group struct {
	name string
	getter Getter
	m  service.TinyCache
}

var (
	groups sync.Map
	keyErr error = errors.New("key does not exit")
)



func NewGroup(name string,maxEntries int , sharedNums int ,onEvicted func(key string, value interface{}), getter Getter) *Group {
	if sharedNums == 0 || maxEntries == 0 || getter == nil {
		panic("params error")
	}
	group := &Group{
		name: name,
		getter: getter,
		m: *service.NewTinyCache(maxEntries,sharedNums,onEvicted),
	}
	groups.Store(name,group)
	return group
}

func GetGroup(name string) *Group {
	if value,err := groups.Load(name) ; !err {
		return value.(*Group)
	}
	return nil
}

func (g *Group) Get(key string) (string,error) {
	if key == "" {
		return "",keyErr
	}
	if v := g.m.Get(key) ; v != nil {
		return  v.(string)  ,nil
	}
	// 本地获取
	//如果没有获取到从远端获取
	//通过用户已定义数据获取
	return g.getFromPeer(key)
}

func (g *Group) getFromPeer(key string) (string, error){
	if value , err := g.getter.Get(key) ; err == nil {
			g.m.Set(key,value)
			return value, nil
	}
	return "",keyErr
}
