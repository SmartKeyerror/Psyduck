package memory

import (
	"fmt"
	"testing"
)

type Message struct {
	Msg string
}

var msg *Message
var Done bool

// 指令重排是指写在程序中的两条语句在 CPU 执行时可能会出现，执行顺序与编写顺序不一致的情况。
// 也就是说，可能会出现 Done 被赋值成为 true 之后，才会去执行 msg = new(Message) 这两条指令
// 那么在 TestMemoryReordering 执行时，很有可能出现跳出了无限循环，但是 msg 仍然为 nil 的情况

// 不过，这件事情不会在单个的 Goroutine 中发生，因为 Go 保证了 在一个 goroutine 内部，
// 程序的执行顺序和它们的代码指定的顺序是一样的，即使编译器或者 CPU 重排了读写顺序，从行为上来看，也和代码指定的顺序一样。
func MayBeReordering() {
	msg = new(Message)
	msg.Msg = "Hello Memory Reordering"
	Done = true
}

func TestMemoryReordering(t *testing.T) {
	go MayBeReordering()

	for Done == false {
	}

	fmt.Println(msg.Msg)
}
