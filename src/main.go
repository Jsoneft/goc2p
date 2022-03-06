package main

import (
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(5 * time.Second)
		}()
		time.Sleep(time.Second)
	}
	time.Sleep(time.Second)
	runtime.Gosched()
}
