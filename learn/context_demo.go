package learn

import (
	"context"
	"fmt"
	"time"
)

func Demo1() {

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

func Demo2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("monitor end")
				return
			default:
				fmt.Println("执行中")
				time.Sleep(1 * time.Second)

			}
		}
	}(ctx)
	time.Sleep(5 * time.Second)
	// 通知结束了
	cancel()
	time.Sleep(10 * time.Second)
}

func Demo3() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("monitor end")
				return
			default:
				fmt.Println("执行中")
				ctx2, _ := context.WithCancel(ctx)
				go func(ctx context.Context) {
					time.Sleep(5 * time.Second)
					fmt.Println("test1")
					for {
						select {
						case <-ctx.Done():
							fmt.Println("monitor2 end")
							return
						default:
							time.Sleep(2 * time.Second)
							fmt.Println("======")
						}

					}
				}(ctx2)
				time.Sleep(1 * time.Second)

			}
		}
	}(ctx)

	time.Sleep(20 * time.Second)
}


func Demo4() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1100*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}