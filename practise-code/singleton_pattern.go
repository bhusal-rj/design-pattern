package main

import (
	"sync"
)

var name string
var once sync.Once

func LearnSingletonPattern(value string) string {
	//learning singleTon pattern
	//There will be only one value of name
	once.Do(func() {
		name = value
	})
	return name
}
