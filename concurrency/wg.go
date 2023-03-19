package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(ii int) {
			time.Sleep(time.Second)
			fmt.Printf("G %d finished", ii)
			// wg.Done()
			wg.Add(-1) // wg.Done()就是调用了Add(-1)
		}(i)
	}

	// 计数器counter归零时通知所有waiter，实际上是最后一个完成的G唤醒了所有waiter
	wg.Wait()
	fmt.Println("waiter 1 finished")
	wg.Wait()
	fmt.Println("waiter 2 finished")
	wg.Wait()
	fmt.Println("waiter 3 finished")
}
