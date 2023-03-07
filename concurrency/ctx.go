package main

import (
	"context"
	"fmt"
	"time"
)

func writeRedis(ctx context.Context) {
	fmt.Println("writeRedis start")
	for {
		select {
		case <-ctx.Done(): // 父协程通知自己结束
			fmt.Println("writeRedis done")
			return
		default:
			fmt.Println("writeRedis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writeDB(ctx context.Context) {
	fmt.Println("writeDB start")
	for {
		select {
		case <-ctx.Done(): // 父协程通知自己结束
			fmt.Println("writeDB done")
			return
		default:
			fmt.Println("writeDB running")
			time.Sleep(2 * time.Second)
		}
	}
}

func handlerReq(ctx context.Context) {
	fmt.Println("handlerReq start")
	go writeRedis(ctx)
	go writeDB(ctx)
	for {
		select {
		case <-ctx.Done(): // 父协程通知自己结束
			fmt.Println("handlerReq done")
			return
		default:
			fmt.Println("handlerReq running")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	cancelCtx, cancleFunc := context.WithCancel(context.Background())
	go handlerReq(cancelCtx)
	time.Sleep(2 * time.Second)
	cancleFunc() // 取消cancelCtx控制的所有协程
	time.Sleep(5 * time.Second)

	// 支持到时间后自动取消所有子协程
	timerCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// deadlineCtx, _ := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	go handlerReq(timerCtx)
	time.Sleep(10 * time.Second)
}
