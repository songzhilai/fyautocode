package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a uint64 = 3
	atomic.AddUint64(&a, 1)
	atomic.AddUint64(&a, ^uint64(1-1))
	fmt.Println(a)
}
