package main

import "sync"

var once sync.Once

func main() {
	once.Do(func() {

	})
}
