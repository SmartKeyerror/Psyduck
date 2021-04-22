package mutex

import (
	"sync"
	"testing"
)

func foo(number int, mutex sync.Locker) {
	if number == 1 {
		return
	}
	mutex.Lock()

	foo(number-1, mutex)

	mutex.Unlock()
}

// fatal error: all goroutines are asleep - deadlock!
// 发生了死锁，也就是说，Mutex 是不可重入的。
// 因为 Mutex 在设计的时候，就没有把具体的 Goroutine 写入至 Mutex 的状态中，
// 因此所有的 Goroutine 的地位都是一样的，尽管当前 Goroutine 已经获得了锁，
func TestReentrantMutex(t *testing.T) {
	var mutex sync.Mutex
	foo(10, &mutex)
}
