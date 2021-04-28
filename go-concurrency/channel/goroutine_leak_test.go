package channel

import (
	"fmt"
	"testing"
	"time"
)

// 这段代码的问题就在于，使用了一个 unbuffered channel，
// unbuffered channel 只有在 receiver 和 sender 都准备好的时候才不会发生阻塞。
// 当任务执行时间超过 2s 时，定时器到期后直接退出 select {} 多路复用，
// 我们就再也不可能接收到 ch 中的数据了。也就是说，执行耗时任务的那个 goroutine 永远都不会退出
func TestGoroutineLeak(t *testing.T) {
	ch := make(chan struct{})

	go func() {
		// 模拟耗时任务
		time.Sleep(5 * time.Second)
		// 任务执行完毕后向外发出通知
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		fmt.Println("Task Done")
	case <-time.After(2 * time.Second):
		fmt.Println("Time Limit Exceeded")
	}
}

// 这里有一个关于 GC 的问题，假如下面儿的 ch 中装的堆变量，那么当函数运行结束后，ch 会被 GC 回收掉吗?
// channel 源码: https://github.com/golang/go/blob/2ebe77a2fda1ee9ff6fd9a3e08933ad1ebaea039/src/runtime/chan.go#L355
// 一个比较有意思的地方在于，channel 的 close 方法并没有去清理环形队列 buf 中的内容，仅仅只是清理了两个阻塞队列中的内容，
// 这和 channel 的行为是一致的: 当我们关闭了一个 channel 以后，receiver 仍然可以继续读取 channel 中的内容，直到没有任何数据
// 如果 buf 中有内容的 ch 可以被回收，那么是什么时候被 GC 的呢? GC 又是通过什么条件去判断 channel 可以被回收的呢?
func TestTimeOutFunction(t *testing.T) {
	ch := make(chan struct{}, 1)

	go func() {
		// 模拟耗时任务
		time.Sleep(5 * time.Second)
		// 任务执行完毕后向外发出通知
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		fmt.Println("Task Done")
	case <-time.After(2 * time.Second):
		fmt.Println("Time Limit Exceeded")
	}
}
