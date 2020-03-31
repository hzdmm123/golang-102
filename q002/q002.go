package q002

import (
	"fmt"
	"math/rand"
	"sync"
)

/**
写代码实现两个 goroutine，其中一个产生随机数并写入到 go channel 中，
另外一个从 channel 中读取数字并打印到标准输出。最终输出五个随机数。
*/

func PrintRandomIntWithChan() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	msg := make(chan int)
	go func() {

		defer func() {
			wg.Done()
		}()

		for i := 0; i < 5; i++ {
			msg <- rand.Intn(5)
		}

		close(msg)
	}()

	go func() {

		defer wg.Done()
		for i := range msg {
			fmt.Println(i)
		}

	}()

	wg.Wait()
}
