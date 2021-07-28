package main

import (
	"fmt"
	"time"
)

// Object 资源对象
type Object struct {
}

// Do 模拟对象行为
func (obj *Object) Do(i int,p Pool) {
	fmt.Println("doing",i)
	time.Sleep(1 * time.Second)
	p <- obj
}

// Pool 资源池
type Pool chan *Object


// New 初始化资源池
func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		p <- new(Object)
	}

	return &p
}

func main() {
	p := New(2)

    // 消耗资源对象
    for i:=0;i<5;i++{
        for {
            select {
            case obj := <-*p:
                go obj.Do(i, *p)
            default:
                continue
            }
            break
        }
    }

    // 等待任务全部完成
    for {
        if len(*p) != 2{
            continue
        }
        break
    }
}
