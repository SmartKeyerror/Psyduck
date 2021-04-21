package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 一个经典的 Race Condition 的入门例子，count 的结果取决于程序运行时的环境，也就是
// 多次运行，得到的 count 都会不一样
// 使用 go test -race race_test.go 可发现出该测试文件中存在的竞态问题
func TestRaceCondition(t *testing.T) {
	count := 0

	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				mutex.Lock()
				count ++
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}

// Goroutine g1 通过 Lock() 方法获取了锁，在 g1 持有期间，另一个 Goroutine g2 调用
// Unlock() 想要释放这个锁，那么这个调用是否会成功? g1 还能不能正常运行?
// g2 能够成功调用 Unlock，但是后续 g1 调用 Unlock 时将会 panic。这操作真的挺骚的：走你的路，让你无路可走
func TestUnlockOtherCorLock(t *testing.T) {
	var mutex sync.Mutex

	// 假设该 Goroutine 叫 A，主流程的 Goroutine 叫 B
	go func() {
		mutex.Lock()
		time.Sleep(5 * time.Second)
		mutex.Unlock()
	}()

	// A 此时可以 unlock mutex 吗?
	time.Sleep(time.Second)
	mutex.Unlock()

	select {}
}
