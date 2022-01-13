# 欢迎！

学习使用本指南：包括如何在不同的课程间切换以及运行代码。

Go 作者组编写，Go-zh 小组翻译。
https://go-zh.org

## 1.Hello, 世界

欢迎来到 [Go 编程语言](https://go-zh.org/) 指南。

点击页面左上角的 [Go 指南](list.md) 可以访问本指南的模块列表。

你可以随时点击页面右上角的菜单来浏览内容列表。

本指南由一系列幻灯片贯穿始终，其中有一些练习需要你来完成。

你可以

- 按`上一页` 或 `PageUp` 键跳转到上一页，
- 按`下一页` 或 `PageDown` 键跳转到下一页。

该指南可以进行交互。点击`运行`按钮（或按 `Shift+Enter`）可以在远程服务器上 你的电脑上 编译并执行程序。运行结果会显示在代码下面。

本指南中的示例程序展示了 Go 的各个方面，它们可以成为你积累经验的开始。

编辑程序并再次执行它。

注意当你点击`格式化`（快捷键： `Ctrl-Enter`）时， 编辑器中的文本会被 [gofmt](https://go-zh.org/cmd/gofmt/) 工具格式化。 你可以点击`语法`开启或关闭语法高亮。

如果你准备好了，请点击页面底部的`右箭头`或按 `PageDown` 键继续。

[hello.go](ch0-welcome/hello/hello.go)

## 2.Go 本地化

本指南也有其它语言的版本：

- [Brazilian Portuguese — Português do Brasil](https://go-tour-br.appspot.com/)
- [Catalan — Català](https://go-tour-ca.appspot.com/)
- [Simplified Chinese — 中文（简体）](https://tour.go-zh.org/)
- [Traditional Chinese — 中文（繁體）](https://go-tour-zh-tw.appspot.com/)
- [Czech — Česky](https://go-tour-cz.appspot.com/)
- [English](https://tour.golang.org/)
- [French — Français](https://go-tour-fr.appspot.com/)
- [German — Deutsch](https://go-tour-de.appspot.com/)
- [Hebrew — עִבְרִית](https://go-tour-he.appspot.com/)
- [Indonesian — Bahasa Indonesia](https://go-tour-id2.appspot.com/)
- [Italian — Italiano](https://go-tour-ita.appspot.com/)
- [Japanese — 日本語](https://go-tour-jp.appspot.com/)
- [Korean — 한국어](https://go-tour-kr.appspot.com/)
- [Romanian — Română](https://go-tour-ro.appspot.com/)
- [Russian - Русский](https://go-tour-ru-ru.appspot.com/)
- [Spanish — Español](https://gotour-es.appspot.com/)
- [Thai - ภาษาไทย](https://go-tour-th.appspot.com/)
- [Turkish - Türkçe](https://go-tour-turkish.appspot.com/)
- [Ukrainian — Українська](https://go-tour-ua.appspot.com/)
- [Uzbek — Ўзбекча](https://go-tour-uz.appspot.com/)

点击`下一页`按钮或者按 `PageDown` 键继续。

## 3.Go 离线版教程

本指南也可作为独立的程序使用，这样你无需访问互联网就能运行它。独立的 tour 更快，它会在你自己的机器上构建并运行代码示例。要在本地安装并运行此教程的中文版，请在命令行执行：

```bash
go get -u github.com/Go-zh/tour tour
```

该程序会打开一个浏览器并显示你本地版本的教程。当然，你也可以继续在线学习本教程。

## 4.Go 练习场

本指南构建在 [Go 练习场](https://play.golang.org/) 上，它是一个运行在[golang.org](https://golang.org/) 服务器上的一个 Web 服务。 该服务会接收一个 Go
程序，然后在沙盒中编译、链接并运行它，最后返回输出。在练习场中运行的程序有一些限制：

- 练习场中的时间始于 2009-11-10 23:00:00 UTC（此日期的含义留给读者自己去发现）。 赋予程序确定的输出能让缓存程序更加容易。
- 程序的执行时间、CPU 和内存的使用同样也有限制，并且程序不能访问外部网络中的主机。 练习场使用最新发布的 Go 的稳定版本。参阅 [Go 练习场的内部机制](https://blog.go-zh.org/playground )
  了解更多信息。

[sandbox.go](ch0-welcome/sandbox/sandbox.go)

## 恭喜！

你已经完成了本指南的第一部分！

你可以返回[模块](list.md)列表看看接下来要学什么，或者继续[后面的课程](ch1-basics.md)。
