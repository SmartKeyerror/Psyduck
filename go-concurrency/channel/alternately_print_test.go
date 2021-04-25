package channel

import (
	"fmt"
	"testing"
	"time"
)

// 有四个 goroutine，编号为 1、2、3、4。
// 每秒钟会有一个 goroutine 打印出它自己的编号，要求你编写一个程序，
// 让输出的编号总是按照 1、2、3、4、1、2、3、4、……的顺序打印出来。

func AlternatelyPrint() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	ch3 := make(chan bool)
	ch4 := make(chan bool)

	go func() {
		for {
			fmt.Println("Goroutine 1")
			time.Sleep(time.Second)
			ch2 <- true
			<-ch1
		}
	}()

	go func() {
		for {
			<-ch2
			fmt.Println("Goroutine 2")
			time.Sleep(time.Second)
			ch3 <- true
		}
	}()

	go func() {
		for {
			<-ch3
			fmt.Println("Goroutine 3")
			time.Sleep(time.Second)
			ch4 <- true
		}
	}()

	go func() {
		for {
			<-ch4
			fmt.Println("Goroutine 4")
			time.Sleep(time.Second)
			ch1 <- true
		}
	}()

	select {}
}

// 另一种实现方式
func AlternatelyPrint2() {
	ch := make(chan struct{})

	for i := 1; i < 5; i++ {
		go func(index int) {
			time.Sleep(time.Duration(index) * time.Millisecond)
			for {
				<-ch
				fmt.Printf("Goroutine %d \n", index)
				time.Sleep(time.Second)
				ch <- struct{}{}
			}
		}(i)
	}

	ch <- struct{}{}
	select {}
}

type Token struct{}

// 第三种方式，其实就是把第一种方式抽象了出来
func AlternatelyPrint3(total int) {

	// 初始化一堆 channel，用于多个 Goroutine 间通信
	channels := make([]chan Token, total)
	for i := 0; i < total; i++ {
		channels[i] = make(chan Token)
	}

	for i := 0; i < total; i++ {
		go func(index int, current chan Token, nextChan chan Token) {
			for {
				<-current
				fmt.Printf("Goroutine %d \n", index)
				time.Sleep(time.Second)
				nextChan <- Token{}
			}
		}(i+1, channels[i], channels[(i+1)%total])
	}
	channels[0] <- Token{}
	select {}
}

func TestAlternatelyPrint(t *testing.T) {
	//AlternatelyPrint()
	//AlternatelyPrint2()
	AlternatelyPrint3(4)
}
