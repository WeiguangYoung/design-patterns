package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once

	instance map[string]string
)

// New 实例化对象
func New() map[string]string {
	once.Do(func() {
		instance = make(map[string]string)
	})
	return instance
}
func main() {
	s := New()
	s["this"] = "that"
	s2 := New()
	fmt.Println(s2["this"])
}
