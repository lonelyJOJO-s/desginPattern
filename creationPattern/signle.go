package creationpattern

import (
	"sync"
	"sync/atomic"
)

// 饿汉
type singleton1 struct{}

var instance1 *singleton1

func init() {
	instance1 = &singleton1{}
}

func GetSingleton1() *singleton1 {
	return instance1
}

// 懒汉
type singleton2 struct{}

var instance2 *singleton2
var lock sync.Mutex
var initialize uint32

func GetSingleton2() *singleton2 {
	if atomic.LoadUint32(&initialize) == 1 {
		return instance2
	}
	lock.Lock()
	defer lock.Unlock()
	// 线程安全
	if instance2 == nil {
		instance2 = new(singleton2)
		atomic.StoreUint32(&initialize, 1)
	}
	return instance2

}

// or just use once
type singleton struct{}

var instance *singleton
var once sync.Once

func GetSingleton() *singleton {
	once.Do(func() { // same as atomic
		instance = new(singleton)
	})
	return instance
}
