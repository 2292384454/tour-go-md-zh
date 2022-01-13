# 方法和接口

本节课包含了方法和接口，可以用这种构造来定义对象及其行为。

Go 作者组编写，Go-zh 小组翻译。
https://go-zh.org

## 1.方法

Go 没有类。不过你可以为结构体类型定义方法。

方法就是一类带特殊的 **接收者**  参数的函数。

方法接收者在它自己的参数列表内，位于 `func` 关键字和方法名之间。

在此例中，`Abs` 方法拥有一个名为 `v`，类型为 `Vertex` 的接收者。

[methods.go](ch4-methods/methods/methods.go)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

## 2.方法即函数

记住：方法只是个带接收者参数的函数。

现在这个 `Abs` 的写法就是个正常的函数，功能并没有什么变化。

[methods-funcs.go](ch4-methods/methods-funcs/methods-funcs.go)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

## 3.方法（续）

你也可以为非结构体类型声明方法。

在此例中，我们看到了一个带 `Abs` 方法的数值类型 `MyFloat`。

你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 `int` 之类的内建类型）的接收者声明方法。

（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）

[methods-continued.go](ch4-methods/methods-continued/methods-continued.go)

```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

## 4.指针接收者

你可以为指针接收者声明方法。

这意味着对于某类型 `T`，接收者的类型可以用 `*T` 的文法。（此外，`T` 不能是像 `*int` 这样的指针。）

例如，这里为 `*Vertex` 定义了 `Scale` 方法。

指针接收者的方法可以修改接收者指向的值（就像 `Scale` 在这做的）。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。

试着移除第 16 行 `Scale` 函数声明中的 `*`，观察此程序的行为如何变化。

若使用值接收者，那么 `Scale` 方法会对原始 `Vertex` 值的副本进行操作。（对于函数的其它参数也是如此。）`Scale` 方法必须用指针接受者来更改 `main` 函数中声明的 `Vertex` 的值。

[methods-pointers.go](ch4-methods/methods-pointers/methods-pointers.go)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

## 5.指针与函数

现在我们要把 `Abs` 和 `Scale` 方法重写为函数。

同样，我们先试着移除掉第 16 的 `*`。你能看出为什么程序的行为改变了吗？要怎样做才能让该示例顺利通过编译？

（若你不确定，继续往下看。）

[methods-pointers-explained.go](ch4-methods/methods-pointers-explained/methods-pointers-explained.go)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
```

## 6.方法与指针重定向

比较前两个程序，你大概会注意到带指针参数的函数必须接受一个指针：

	var v Vertex
	ScaleFunc(v, 5)  // 编译错误！
	ScaleFunc(&v, 5) // OK

而以指针为接收者的方法被调用时，接收者既能为值又能为指针：

	var v Vertex
	v.Scale(5)  // OK
	p := &v
	p.Scale(10) // OK

对于语句 `v.Scale(5)`，即便 `v` 是个值而非指针，带指针接收者的方法也能被直接调用。 也就是说，由于 `Scale` 方法有一个指针接收者，为方便起见，Go 会将语句 `v.Scale(5)`
解释为 `(&v).Scale(5)`。

[indirection.go](ch4-methods/indirection/indirection.go)

```go
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

## 7.方法与指针重定向（续）

同样的事情也发生在相反的方向。

接受一个值作为参数的函数必须接受一个指定类型的值：

	var v Vertex
	fmt.Println(AbsFunc(v))  // OK
	fmt.Println(AbsFunc(&v)) // 编译错误！

而以值为接收者的方法被调用时，接收者既能为值又能为指针：

	var v Vertex
	fmt.Println(v.Abs()) // OK
	p := &v
	fmt.Println(p.Abs()) // OK

这种情况下，方法调用 `p.Abs()` 会被解释为 `(*p).Abs()`。

[indirection-values.go](ch4-methods/indirection-values/indirection-values.go)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
```

## 8.选择值或指针作为接收者

使用指针接收者的原因有二：

首先，方法能够修改其接收者指向的值。

其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

在本例中，`Scale` 和 `Abs` 接收者的类型为 `*Vertex`，即便 `Abs` 并不需要修改其接收者。

通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。（我们会在接下来几页中明白为什么。）

[methods-with-pointer-receivers.go](ch4-methods/methods-with-pointer-receivers/methods-with-pointer-receivers.go)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

## 9.接口

**接口类型**是由一组方法签名定义的集合。

接口类型的变量可以保存任何实现了这些方法的值。

**注意:** 示例代码的 22 行存在一个错误。由于 `Abs` 方法只为 `## Vertex` （指针类型）定义，因此 `Vertex`（值类型）并未实现 `Abser`。

[interfaces.go](ch4-methods/interfaces/interfaces.go)

```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

## 10.接口与隐式实现

类型通过实现一个接口的所有方法来实现该接口。既然无需专门显式声明，也就没有“implements”关键字。

隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。

因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

[interfaces-are-satisfied-implicitly.go](ch4-methods/interfaces-are-satisfied-implicitly/interfaces-are-satisfied-implicitly.go)

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

## 11.接口值

接口也是值。它们可以像其它值一样传递。

接口值可以用作函数的参数或返回值。

在内部，接口值可以看做包含值和具体类型的元组：

	(value, type)

接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。

[interface-values.go](ch4-methods/interface-values/interface-values.go)

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## 12.底层值为 nil 的接口值

即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 `M` 方法）。

**注意:** 保存了 nil 具体值的接口其自身并不为 nil。

[interface-values-with-nil.go](ch4-methods/interface-values-with-nil/interface-values-with-nil.go)

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## 13.nil 接口值

nil 接口值既不保存值也不保存具体类型。

为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 ## 具体## 方法的类型。

[nil-interface-values.go](ch4-methods/nil-interface-values/nil-interface-values.go)

```go
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## 14.空接口

指定了零个方法的接口值被称为 **空接口:**

	interface{}

空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

空接口被用来处理未知类型的值。例如，`fmt.Print` 可接受类型为 `interface{}` 的任意数量的参数。

[empty-interface.go](ch4-methods/empty-interface/empty-interface.go)

```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

## 15.类型断言

**类型断言**提供了访问接口值底层具体值的方式。

	t := i.(T)

该语句断言接口值 `i` 保存了具体类型 `T`，并将其底层类型为 `T` 的值赋予变量 `t`。

若 `i` 并未保存 `T` 类型的值，该语句就会触发一个panic。

为了 ## 判断## 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

	t, ok := i.(T)

若 `i` 保存了一个 `T`，那么 `t` 将会是其底层值，而 `ok` 为 `true`。

否则，`ok` 将为 `false` 而 `t` 将为 `T` 类型的零值，程序并不会产生panic。

请注意这种语法和读取一个映射时的相同之处。

[type-assertions.go](ch4-methods/type-assertions/type-assertions.go)

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}
```

## 16.类型选择

**类型选择**是一种按顺序从几个类型断言中选择分支的结构。

类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

	switch v := i.(type) {
	case T:
		// v 的类型为 T
	case S:
		// v 的类型为 S
	default:
		// 没有匹配，v 与 i 的类型相同
	}

类型选择中的声明与类型断言 `i.(T)` 的语法相同，只是具体类型 `T` 被替换成了关键字 `type`。

此选择语句判断接口值 `i` 保存的值类型是 `T` 还是 `S`。在 `T` 或 `S` 的情况下，变量 `v` 会分别按 `T` 或 `S` 类型保存 `i` 拥有的值。在默认（即没有匹配）的情况下，变量 `v` 与 `i`
的接口类型和值相同。

[type-switches.go](ch4-methods/type-switches/type-switches.go)

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

## 17.Stringer

[fmt](https://go-zh.org/pkg/fmt/) 包中定义的 [Stringer](https://go-zh.org/pkg/fmt/#Stringer) 是最普遍的接口之一。

	type Stringer interface {
		String() string
	}

`Stringer` 是一个可以用字符串描述自己的类型。`fmt` 包（还有很多包）都通过此接口来打印值。

[stringer.go](ch4-methods/stringer/stringer.go)

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
```

## 18.练习：Stringer

通过让 `IPAddr` 类型实现 `fmt.Stringer` 来打印点号分隔的地址。

例如，`IPAddr{1,`2,`3,`4}` 应当打印为 `"1.2.3.4"`。

[exercise-stringer.go](ch4-methods/exercise-stringer/exercise-stringer.go)

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: 给 IPAddr 添加一个 "String() string" 方法

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

## 19.错误

Go 程序使用 `error` 值来表示错误状态。

与 `fmt.Stringer` 类似，`error` 类型是一个内建接口：

	type error interface {
		Error() string
	}

（与 `fmt.Stringer` 类似，`fmt` 包在打印值时也会满足 `error`。）

通常函数会返回一个 `error` 值，调用它的代码应当判断这个错误是否等于 `nil` 来进行错误处理。

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)

`error` 为 nil 时表示成功；非 nil 的 `error` 表示失败。

[errors.go](ch4-methods/errors/errors.go)

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
```

## 20.练习：错误

从[之前的练习](ch2-flowcontrol/exercise-loops-and-functions/exercise-loops-and-functions.go)中复制 `Sqrt` 函数，修改它使其返回 `error` 值。

`Sqrt` 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

创建一个新的类型

	type ErrNegativeSqrt float64

并为其实现

	func (e ErrNegativeSqrt) Error() string

方法使其拥有 `error` 值，通过 `ErrNegativeSqrt(-2).Error()` 调用该方法应返回 `"cannot`Sqrt`negative`number:`-2"`。

**注意:** 在 `Error` 方法内调用 `fmt.Sprint(e)` 会让程序陷入死循环。可以通过先转换 `e`

来避免这个问题：`fmt.Sprint(float64(e))`。这是为什么呢？

修改 `Sqrt` 函数，使其接受一个负数时，返回 `ErrNegativeSqrt` 值。

[exercise-errors.go](ch4-methods/exercise-errors/exercise-errors.go)

```go
package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
```

## 21.Reader

`io` 包指定了 `io.Reader` 接口，它表示从数据流的末尾进行读取。

Go 标准库包含了该接口的[许多实现](https://go-zh.org/search?q=Read#Global) ，包括文件、网络连接、压缩和加密等等。

`io.Reader` 接口有一个 `Read` 方法：

	func (T) Read(b []byte) (n int, err error)

`Read` 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 `io.EOF` 错误。

示例代码创建了一个 [strings.Reader](https://go-zh.org/pkg/strings/#Reader) 并以每次 8 字节的速度读取它的输出。

[reader.go](ch4-methods/reader/reader.go)

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```

## 22.练习：Reader

实现一个 `Reader` 类型，它产生一个 ASCII 字符 `'A'` 的无限流。

[exercise-reader.go](ch4-methods/exercise-reader/exercise-reader.go)

```go
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法

func main() {
	reader.Validate(MyReader{})
}
```

## 23.练习：rot13Reader

有种常见的模式是一个 [io.Reader](https://go-zh.org/pkg/io/#Reader) 包装另一个 `io.Reader`，然后通过某种方式修改其数据流。

例如，[gzip.NewReader](https://go-zh.org/pkg/compress/gzip/#NewReader) 函数接受一个 `io.Reader`
（已压缩的数据流）并返回一个同样实现了 `io.Reader` 的 `## gzip.Reader`（解压后的数据流）。

编写一个实现了 `io.Reader` 并从另一个 `io.Reader` 中读取数据的 `rot13Reader`，通过应用 [rot13](http://en.wikipedia.org/wiki/ROT13 )
代换密码对数据流进行修改。

`rot13Reader` 类型已经提供。实现 `Read` 方法以满足 `io.Reader`。

[exercise-rot-reader.go](ch4-methods/exercise-rot-reader/exercise-rot-reader.go)

```go
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
```

## 24.图像

[`image`](https://go-zh.org/pkg/image/#Image) 包定义了 `Image` 接口：

	package image
	
	type Image interface {
		ColorModel() color.Model
		Bounds() Rectangle
		At(x, y int) color.Color
	}

**注意:**  `Bounds` 方法的返回值 `Rectangle` 实际上是一个 [image.Rectangle](https://go-zh.org/pkg/image/#Rectangle) ，它在 `image`
包中声明。

（请参阅[文档](https://go-zh.org/pkg/image/#Image) 了解全部信息。）

`color.Color` 和 `color.Model` 类型也是接口，但是通常因为直接使用预定义的实现 `image.RGBA` 和 `image.RGBAModel`
而被忽视了。这些接口和类型由 [image/color](https://go-zh.org/pkg/image/color/) 包定义。

[images.go](ch4-methods/images/images.go)

```go
package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
```

## 25.练习：图像

还记得之前编写的[图片生成器](ch3-moretypes/exercise-slices/exercise-slices.go) 吗？我们再来编写另外一个，不过这次它将会返回一个 `image.Image` 的实现而非一个数据切片。

定义你自己的 `Image` 类型，实现[必要的方法](https://go-zh.org/pkg/image/#Image) 并调用 `pic.ShowImage`。

`Bounds` 应当返回一个 `image.Rectangle` ，例如 `image.Rect(0,`0,`w,`h)`。

`ColorModel` 应当返回 `color.RGBAModel`。

`At` 应当返回一个颜色。上一个图片生成器的值 `v` 对应于此次的 `color.RGBA{v,`v,`255,`255}`。

[exercise-images.go](ch4-methods/exercise-images/exercise-images.go)

```go
package main

import "golang.org/x/tour/pic"

type Image struct{}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
```

## 26.恭喜！

你已经完成了本课程！

你可以返回[模块](list.md)列表看看接下来要学什么，或者继续[后面的课程](ch5-concurrency.md)。
