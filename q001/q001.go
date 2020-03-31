package q001

import (
	"fmt"
	"time"
)

/**
交替打印数字1-100 的奇偶数
*/

var Numbers = 100

func printOdd(p chan bool) {
	for i := 1; i < Numbers; i++ {

		p <- true//发送信号

		if i%2 == 1 {
			fmt.Println("Odd goroutine: ", i)
		}
	}
}

func printEven(p chan bool) {
	for i := 1; i < Numbers; i++ {

		<-p //接受信号

		if i%2 == 0 {
			fmt.Println("Even goroutine: ", i)
		}
	}
}

func PrintTheOddAndEvenResult() {
	// 缓冲区的大小位0，意味着goroutine odd必须等上一个值被消费了 ，才可以塞下一个值
	// 改chan只用来同步两个协程的执行
	msg := make(chan bool)

	go printOdd(msg)
	go printEven(msg)

	time.Sleep(time.Second * 2)
}

// 总结：使用chan作为同步的工具，发送信号，让协程执行
//
