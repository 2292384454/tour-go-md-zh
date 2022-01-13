# 并发

Go 将并发结构作为核心语言的一部分提供。本节课程通过一些示例介绍并展示了它们的用法。

Go 作者组编写，Go-zh 小组翻译。
https://go-zh.org

## 1.Go 程

Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

	go f(x, y, z)

会启动一个新的 Go 程并执行

	f(x, y, z)

`f`, `x`, `y` 和 `z` 的求值发生在当前的 Go 程中，而 `f` 的执行发生在新的 Go 程中。

Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。[sync](https://go-zh.org/pkg/sync/ ) 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法（见下一页）。

[goroutines.go](ch5-concurrency/goroutines/goroutines.go)
```go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```


## 2.信道

信道是带有类型的管道，你可以通过它用信道操作符 `<-` 来发送或者接收值。

	ch <- v    // 将 v 发送至信道 ch。
	v := <-ch  // 从 ch 接收值并赋予 v。

（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

	ch := make(chan int)

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。

[channels.go](ch5-concurrency/channels/channels.go)
```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
```


## 3.带缓冲的信道

信道可以是 _带缓冲的_。将缓冲长度作为第二个参数提供给 `make` 来初始化一个带缓冲的信道：

	ch := make(chan int, 100)

仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

修改示例填满缓冲区，然后看看会发生什么。

[buffered-channels.go](ch5-concurrency/buffered-channels/buffered-channels.go)
```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
```


## 4.range 和 close

发送者可通过 `close` 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

	v, ok := <-ch

此时 `ok` 会被设置为 `false`。

循环 `for`i`:=`range`c` 会不断从信道接收值，直到它被关闭。

**注意**： 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。

**还要注意**： 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 `range` 循环。

[range-and-close.go](ch5-concurrency/range-and-close/range-and-close.go)
```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```


## 5.select 语句

`select` 语句使一个 Go 程可以等待多个通信操作。

`select` 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

[select.go](ch5-concurrency/select/select.go)
```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```


## 6.默认选择

当 `select` 中的其它分支都没有准备好时，`default` 分支就会执行。

为了在尝试发送或者接收时不发生阻塞，可使用 `default` 分支：

	select {
	case i := <-c:
		// 使用 i
	default:
		// 从 c 中接收会阻塞时执行
	}

[default-selection.go](ch5-concurrency/default-selection/default-selection.go)
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```


## 7.练习：等价二叉查找树

不同二叉树的叶节点上可以保存相同的值序列。例如，以下两个二叉树都保存了序列 `1，1，2，3，5，8，13`。

![tree.png](img/tree.png)

在大多数语言中，检查两个二叉树是否保存了相同序列的函数都相当复杂。 我们将使用 Go 的并发和信道来编写一个简单的解法。

本例使用了 `tree` 包，它定义了类型：

	type Tree struct {
		Left  *Tree
		Value int
		Right *Tree
	}

点击`下一页`继续。

## 8.练习：等价二叉查找树

1. 实现 `Walk` 函数。
2. 测试 `Walk` 函数。

函数 `tree.New(k)` 用于构造一个随机结构的已排序二叉查找树，它保存了值 `k`, `2k`, `3k`, ..., `10k`。

创建一个新的信道 `ch` 并且对其进行步进：

	go Walk(tree.New(1), ch)

然后从信道中读取并打印 10 个值。应当是数字 `1,`2,`3,`...,`10`。

3. 用 `Walk` 实现 `Same` 函数来检测 `t1` 和 `t2` 是否存储了相同的值。
4. 测试 `Same` 函数。

`Same(tree.New(1),`tree.New(1))` 应当返回 `true`，而 `Same(tree.New(1),`tree.New(2))` 应当返回 `false`。

`Tree` 的文档可在[这里](https://godoc.org/golang.org/x/tour/tree#Tree )找到。

[exercise-equivalent-binary-trees.go](ch5-concurrency/exercise-equivalent-binary-trees/exercise-equivalent-binary-trees.go)
```go
package main

import "golang.org/x/tour/tree"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int)

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool

func main() {
}
```


## 9.sync.Mutex

我们已经看到信道非常适合在各个 Go 程间进行通信。

但是如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？

这里涉及的概念叫做 *互斥（mutual*exclusion）## ，我们通常使用 *互斥锁（Mutex）## 这一数据结构来提供这种机制。

Go 标准库中提供了 [sync.Mutex](https://go-zh.org/pkg/sync/#Mutex ) 互斥锁类型及其两个方法：

- `Lock`
- `Unlock`

我们可以通过在代码前调用 `Lock` 方法，在代码后调用 `Unlock` 方法来保证一段代码的互斥执行。参见 `Inc` 方法。

我们也可以用 `defer` 语句来保证互斥锁一定会被解锁。参见 `Value` 方法。

[mutex-counter.go](ch5-concurrency/mutex-counter/mutex-counter.go)
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```


## 10.练习：Web 爬虫

在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 `Crawl` 函数来并行地抓取 URL，并且保证不重复。

_提示_：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！

[exercise-web-crawler.go](ch5-concurrency/exercise-web-crawler/exercise-web-crawler.go)
```go
package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
        // 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
```


## 11.接下来去哪？

你可以从[安装 Go](https://go-zh.org/doc/install/ ) 开始。

一旦安装了 Go，Go [文档](https://go-zh.org/doc/ )是一个极好的应当继续阅读的内容。 它包含了参考、指南、视频等等更多资料。 了解如何组织 Go
代码并在其上工作，参阅[此视频](https://www.youtube.com/watch?v=XCsL89YtqCs )
，或者阅读[如何编写 Go 代码](https://go-zh.org/doc/code.html )。 如果你需要标准库方面的帮助，请参考[包手册](https://go-zh.org/pkg/ )
。如果是语言本身的帮助，阅读[语言规范](https://go-zh.org/ref/spec )是件令人愉快的事情。

进一步探索 Go 的并发模型，参阅 [Go 并发模型](https://www.youtube.com/watch?v=f6kdp27TYZs )
([幻灯片](https://talks.go-zh.org/2012/ch5-concurrency.slide ))
以及[深入 Go 并发模型](https://www.youtube.com/watch?v=QDDwwePbDtw )([幻灯片](https://talks.go-zh.org/2013/advconc.slide ))
并阅读[通过通信共享内存](https://go-zh.org/doc/codewalk/sharemem/ )的代码之旅。

想要开始编写 Web 应用，请参阅[一个简单的编程环境](https://vimeo.com/53221558 )([幻灯片](https://talks.go-zh.org/2012/simple.slide ))
并阅读[编写 Web 应用](https://go-zh.org/doc/articles/wiki/ )的指南。

[函数：Go 中的一等公民](https://go-zh.org/doc/codewalk/functions/ )展示了有趣的函数类型。

[Go 博客](https://blog.go-zh.org/ )有着众多关于 Go 的文章和信息。

[mikespook 的博客](https://www.mikespook.com/tag/golang/ )中有大量中文的关于 Go 的文章和翻译。

开源电子书 [Go Web 编程](https://github.com/astaxie/build-web-application-with-golang )
和 [Go 入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN )能够帮助你更加深入的了解和学习 Go 语言。

访问 [go-zh.org](https://go-zh.org ) 了解更多内容。
