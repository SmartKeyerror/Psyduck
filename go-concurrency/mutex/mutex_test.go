package mutex

import (
	"fmt"
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

// 我们也可以使用 Mutex 来实现单例模式，比如配置文件的生成

type Config struct {
	Host    string
	Port    uint16
	Workers uint8
}

var config *Config

// 传统的双重校验锁方式，不过这种写法在其它一些语言中，例如 C++ 中，会因为指令重排的原因导致创建多个实例
// 所以其实现在很少有人会这么用，支持静态变量的语言都会利用“静态变量只会初始化一次”的特性来实现
// 而 Go 中则提供了 sync.Once 来实现单例
func GetConfig() *Config {
	var mutex sync.Mutex

	if config == nil {
		mutex.Lock()
		if config == nil {
			config = &Config{
				Host:    "127.0.0.1",
				Port:    8080,
				Workers: 4,
			}
		}
		mutex.Unlock()
	}
	return config
}

func TestSingletonConfig(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			config := GetConfig()
			fmt.Printf("config address : %p \n", config)
		}()
	}

	wg.Wait()

	fmt.Printf("config address : %p", GetConfig())
}
