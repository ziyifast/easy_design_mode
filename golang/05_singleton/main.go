package main

import (
	"fmt"
	"sync"
)

// 懒汉式：用到才加载【饿汉式：直接放在init方法里，程序一启动就创建好】
var (
	instance *Singleton
	once     = sync.Once{}
)

type Singleton struct {
}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	one := GetInstance()
	two := GetInstance()
	//one=0x100f54088
	//two=0x100f54088
	fmt.Printf("one=%p\n", one)
	fmt.Printf("two=%p\n", two)
}
