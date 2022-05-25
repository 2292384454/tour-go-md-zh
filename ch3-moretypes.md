# 更多类型：struct、slice 和映射

学习如何基于现有类型定义新的类型：本节课涵盖了结构体、数组、切片和映射。

Go 作者组编写，Go-zh 小组翻译。
https://go-zh.org

## 1.指针

Go 拥有指针。指针保存了值的内存地址。

类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil`。

	var p *int

`&` 操作符会生成一个指向其操作数的指针。

	i := 42
	p = &i

`*` 操作符表示指针指向的底层值。

	fmt.Println(*p) // 通过指针 p 读取 i
	*p = 21         // 通过指针 p 设置 i

这也就是通常所说的“间接引用”或“重定向”。

与 C 不同，Go 没有指针运算。

[pointers.go](ch3-moretypes/pointers/pointers.go)
```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
```


## 2.结构体

一个结构体（`struct`）就是一组字段（field）。

[structs.go](ch3-moretypes/structs/structs.go)
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```


## 3.结构体字段

结构体字段使用点号来访问。

[struct-fields.go](ch3-moretypes/struct-fields/struct-fields.go)
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```


## 4.结构体指针

结构体字段可以通过结构体指针来访问。

如果我们有一个指向结构体的指针 `p`，那么可以通过 `(*p).X` 来访问其字段 `X`。不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 `p.X` 就可以。

[struct-pointers.go](ch3-moretypes/struct-pointers/struct-pointers.go)
```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```


## 5.结构体文法

结构体文法通过直接列出字段的值来新分配一个结构体。

使用 `Name:` 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 `&` 返回一个指向结构体的指针。

[struct-literals.go](ch3-moretypes/struct-literals/struct-literals.go)
```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```


## 6.数组

类型 `[n]T` 表示拥有 `n` 个 `T` 类型的值的数组。

表达式

```
var a [10]int
```

会将变量 `a` 声明为拥有 10 个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。

[array.go](ch3-moretypes/array/array.go)
```go
package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```


## 7.切片

每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。

类型 `[]T` 表示一个元素类型为 `T` 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

```
a[low : high]
```

它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 `a` 中下标从 1 到 3 的元素：

```
a[1:4]
```

[slices.go](ch3-moretypes/slices/slices.go)
```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
```


## 8.切片就像数组的引用

切片并不存储任何数据，它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。

[slices-pointers.go](ch3-moretypes/slices-pointers/slices-pointers.go)
```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```


## 9.切片文法

切片文法类似于没有长度的数组文法。

这是一个数组文法：

	[3]bool{true, true, false}

下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

	[]bool{true, true, false}

[slice-literals.go](ch3-moretypes/slice-literals/slice-literals.go)
```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```


## 10.切片的默认行为

在进行切片时，你可以利用它的默认行为来忽略上下界。

切片下界的默认值为 `0`，上界则是该切片的长度。

对于数组

	var a [10]int

来说，以下切片是等价的：

	a[0:10]
	a[:10]
	a[0:]
	a[:]

[slice-bounds.go](ch3-moretypes/slice-bounds/slice-bounds.go)
```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```


## 11.切片的长度与容量

切片拥有 *长度* 和 *容量*。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 `s` 的长度和容量可通过表达式 `len(s)` 和 `cap(s)` 来获取。

只要具有足够的容量，你就可以通过重新切片来扩展一个切片。请试着修改示例程序中的某个切片操作，使其长度超过容量（即将它扩展到超出其容量范围），看看会发生什么。

[slice-len-cap.go](ch3-moretypes/slice-len-cap/slice-len-cap.go)
```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```


## 12.nil 切片

切片的零值是 `nil`。

nil 切片的长度和容量为 0 且没有底层数组。

[nil-slices.go](ch3-moretypes/nil-slices/nil-slices.go)
```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```


## 13.用 make 创建切片

切片可以用内建函数 `make` 来创建，这也是你创建动态数组的方式。

`make` 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

	a := make([]int, 5)  // len(a)=5

要指定它的容量，需向 `make` 传入第三个参数：

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4

[making-slices.go](ch3-moretypes/making-slices/making-slices.go)
```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
```


## 14.切片的切片

切片可包含任何类型，甚至包括其它的切片。

[slices-of-slice.go](ch3-moretypes/slices-of-slice/slices-of-slice.go)
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```


## 15.向切片追加元素

为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 `append` 函数。内建函数的[文档](https://go-zh.org/pkg/builtin/#append )对此函数有详细的介绍。

	func append(s []T, vs ...T) []T

`append` 的第一个参数 `s` 是一个元素类型为 `T` 的切片，其余类型为 `T` 的值将会追加到该切片的末尾。

`append` 的结果是一个包含原切片所有元素加上新添加元素的切片。

当 `s` 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

（要了解关于切片的更多内容，请阅读文章 [Go 切片：用法和本质](https://blog.go-zh.org/go-slices-usage-and-internals )。）

[append.go](ch3-moretypes/append/append.go)
```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```


## 16.Range

`for` 循环的 `range` 形式可遍历切片或映射。

当使用 `for` 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。

[range.go](ch3-moretypes/range/range.go)
```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
```


## 17.range（续）

可以将下标或值赋予 `_` 来忽略它。

    for i, _ := range pow
    for _, value := range pow

若你只需要索引，忽略第二个变量即可。

    for i := range pow

[range-continued.go](ch3-moretypes/range-continued/range-continued.go)
```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```


## 18.练习：切片

实现 `Pic`。它应当返回一个长度为 `dy` 的切片，其中每个元素是一个长度为 `dx`，元素类型为 `uint8` 的切片。当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。

图像的选择由你来定。几个有趣的函数包括 `(x+y)/2`, `x*y`, `x^y`, `x*log(y)` 和 `x%(y+1)`。

（提示：需要使用循环来分配 `[][]uint8` 中的每个 `[]uint8`；请使用 `uint8(intValue)` 在类型之间转换；你可能会用到 `math` 包中的函数。）

[exercise-slices.go](ch3-moretypes/exercise-slices/exercise-slices.go)
```go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
}

func main() {
	pic.Show(Pic)
}
```


## 19.映射 (map)

映射将键映射到值。

映射的零值为 `nil` 。`nil` 映射既没有键，也不能添加键。

`make` 函数会返回给定类型的映射，并将其初始化备用。

[maps.go](ch3-moretypes/maps/maps.go)
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```


## 20.映射的文法

映射的文法与结构体相似，不过必须有键名。

[map-literals.go](ch3-moretypes/map-literals/map-literals.go)
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```


## 21.映射的文法（续）

若顶级类型只是一个类型名，你可以在文法的元素中省略它。

[map-literals-continued.go](ch3-moretypes/map-literals-continued/map-literals-continued.go)
```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```


## 22.修改映射

在映射 `m` 中插入或修改元素：

	m[key] = elem

获取元素：

	elem = m[key]

删除元素：

	delete(m, key)

通过双赋值检测某个键是否存在：

	elem, ok = m[key]

若 `key` 在 `m` 中，`ok` 为 `true` ；否则，`ok` 为 `false`。

若 `key` 不在映射中，那么 `elem` 是该映射元素类型的零值。

同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。

*注## ：若 `elem` 或 `ok` 还未声明，你可以使用短变量声明：

	elem, ok := m[key]

[mutating-maps.go](ch3-moretypes/mutating-maps/mutating-maps.go)
```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```


## 23.练习：映射

实现 `WordCount`。它应当返回一个映射，其中包含字符串 `s` 中每个“单词”的个数。函数 `wc.Test` 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 [strings.Fields](https://go-zh.org/pkg/strings/#Fields) 很有帮助。

[exercise-maps.go](ch3-moretypes/exercise-maps/exercise-maps.go)
```go
package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
```


## 24.函数值

函数也是值。它们可以像其它值一样传递。

函数值可以用作函数的参数或返回值。

[function-values.go](ch3-moretypes/function-values/function-values.go)
```go
package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
```


## 25.函数的闭包

Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

例如，函数 `adder` 返回一个闭包。每个闭包都被绑定在其各自的 `sum` 变量上。

[function-closures.go](ch3-moretypes/function-closures/function-closures.go)
```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```


## 26.练习：斐波纳契闭包

让我们用函数做些好玩的事情。

实现一个 `fibonacci`
函数，它返回一个函数（闭包），该闭包返回一个[斐波纳契数列](https://zh.wikipedia.org/wiki/%E6%96%90%E6%B3%A2%E9%82%A3%E5%A5%91%E6%95%B0%E5%88%97) `(0, 1, 1, 2, 3, 5, ...)`
。

[exercise-fibonacci-closure.go](ch3-moretypes/exercise-fibonacci-closure/exercise-fibonacci-closure.go)
```go
package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```


## 27.恭喜！

你已经完成了本课程！

你可以返回[模块](list.md)列表看看接下来要学什么，或者继续[后面的课程](ch4-methods.md)。
