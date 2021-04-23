package mutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRWMutex(t *testing.T) {
	pool := make(map[string]string)

	var wg sync.WaitGroup
	wg.Add(10)

	var mutex sync.RWMutex
	for i := 0; i < 10; i++ {
		go func() {
			// 读操作
			for j := 0; j < 5; j++ {
				mutex.RLock()
				fmt.Println(pool["name"])
				mutex.RUnlock()
				time.Sleep(time.Millisecond * 500)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		mutex.Lock()
		pool["name"] = "Buz~~"
		mutex.Unlock()
		time.Sleep(time.Second * 1)
	}
}
