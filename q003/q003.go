package q003

import (
	"fmt"
	"time"
)

/**
package main

func main() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
	}()

	select {}
}

func proc() {
	panic("ok")
}
*/
func PrintTimerFunc() {
	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出
		t := time.NewTicker(time.Second * 1)
		for {
			select {
			case <-t.C:
				go func() {
					//defer func() {
					//	if err := recover(); err != nil {
					//		fmt.Println(err)
					//	}
					//}()
					proc()
				}()
			}
		}
	}()

	select {} // 不会提出
}

func proc() {
	//panic("ok")
	fmt.Println("ok")
}
