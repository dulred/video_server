package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个 Ticker，每 2 秒触发一次
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() // 在程序结束时停止 Ticker

	done := make(chan bool)

	// 启动一个 goroutine 来处理 Ticker 的事件
	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			case <-done:
				return
			}
		}
	}()

	// 等待 10 秒，然后停止 Ticker
	time.Sleep(10 * time.Second)
	done <- true
	fmt.Println("Ticker stopped")
}
