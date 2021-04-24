package mutex

import (
	"fmt"
	"sync"
	"testing"
)

type OtherConfig struct {
	Host    string
	Port    uint16
	Workers uint8
}

var once sync.Once
var otherConfig *OtherConfig

func GetOtherConfig() *OtherConfig {

	// 使用 Once 的一个最常见的错误就是把 once 的声明放在了函数里面儿
	// 如果用下面儿的代码运行的话，就会发现每一次 config 的地址都不同
	// var once sync.Once

	once.Do(func() {
		otherConfig = &OtherConfig{
			Host:    "127.0.0.1",
			Port:    9090,
			Workers: 4,
		}
	})

	return otherConfig
}

func TestOnceSingleton(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			config := GetOtherConfig()
			fmt.Printf("other config address : %p \n", config)
		}()
	}

	wg.Wait()

	fmt.Printf("other config address : %p", GetOtherConfig())
}
