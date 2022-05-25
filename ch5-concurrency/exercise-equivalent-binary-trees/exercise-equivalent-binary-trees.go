package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	// NOTE: The implementation leaks goroutines when trees are different.
	// 其原因是这样的，首先，两棵树所包含的数值全部被读取并分别发送到 ch1 和 ch2 中，当
	// 两棵树包含的数值不相同时，上面程序中的 for 循环被中止并返回 false, 但是，此时被
	// 发送到 ch1 或/及 ch2 的数值还没有全部被接收，因此，它们会一直处于等待被接收的状态
	// （在教程里有这样一句话：By default, sends and receives block until the
	// other side is ready），以这种状态留存在内存中（占用内存），相当于内存泄漏。（当
	// 然，当这个程序关闭时，即可释放这部分内存，因此对于非长期运行的小程序来说问题不大。）
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		// one of ch1 and ch2 finished
		if !ok1 || !ok2 {
			// they finished at the same time or not
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)
	fmt.Println(Same(t1, t2))

}
