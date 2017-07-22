# 我的 Go 旅程

## Kotlin vs Go ?

前阵子 [Kevin][kevin] 跟我们分享了一波 [Kotlin][kotlin]。
我感觉 [Kotlin][kotlin] 很好玩，
于是想尝试一下。
但后来 [张总][tothegump] 又问我，
你既然想玩一下 [Kotlin][kotlin]，
那怎么不试一下 [Go][go] 呢？
我就去查了查 `Kotlin VS Go` 的这个话题，
发现 [知乎上有人这么说]：

> Kotlin 的大优势和大劣势都是完全兼容 Java。
> 而 Go 最大优势就是对多线程的优化。

我又去看了看 [Slant上面的对比][slant-kotlin-go]。
嗯，感觉其实两个语言都很有前途啊！
然而 :wink: 我选择 ~~Python~~ Go...

![go-name][go-name]

值得一提的是，
`GoLang` 只是网站名，
有的时候 Google 搜 Go 不好搜，
也会用 GoLang 来搜索。

但 **[Go][go] 语言没有别的名字，本身就叫 `Go`**


## Start my journey

好了，既然选择了 [Go][go]，
那我们首先从语言本身开始走起。

* 安装：因为[我是 Windows 党][windows-setup]，
所以用 [Choco 安装][choco] 或者[去官网下载 msi 装一下][go-dl]都行。

* IDE：因为我是 IDE 党...
而且 [JetBrains 也出了 Go 的 IDE Gogland][https://www.jetbrains.com/go/]
那自然就是选择它了。

* 语法学习：认识新语言的语法，
我首先就会去 [Learn X In Y Minutes][xy-go] 上看看……
然后再[稍微看了看官方的 Go Tour][go-tour]。

![wtf-tabs][tabs]

> 注: 这里十分蛋疼地发现 Go 是推崇 Tabs 的...
> 我思索了一下，决定~~还是用 Space~~入乡随俗，使用 tabs。

百看不如一练，
直接上手用吧。


## Hi, Lirian

首先我们新建一个Project，
等等，
问题来了，
**Go 语言的 Workspace 一般是什么样子？**

[文档里面是这么说的][go-workspace]：

> 一般来说，项目目录下面会有三个文件夹：src, pkg, bin。
>
> 其中 src 是放源文件的，pkg 是放各种~~内裤~~类库、包的，bin是放可执行文件的。
>
> 而项目一般是放在 home 目录下的。
> 假如你要放在其它目录，
> 你就要去设定一下 GOPATH 这个环境变量了。

噢，这里引入了一个新的词，
叫 `GOPATH`。
再看看文档，
发现这个概念跟 [Java 的 classpath][classpath]、
[Python 的 sys.path][sys-path] 很类似。

在我们用 `go build` 命令编译项目的时候，
在 `$GOPATH/src` 下面的源文件们就会被翻出来然后编译。

这里有几条命令可以跑跑看：

```
> go env
set GOROOT=C:\CodeEnv\Go  # Go语言的安装路径
set GOPATH=C:\Code\github\LKI\go  # 本项目的路径
...

> go env GOPATH  # 只查看 GOPATH 这个环境变量
C:\Code\github\LKI\go

> go help gopath  # 简单地看一看关于 GOPATH 的说明
The Go path is bla bla...
bla bla...
bla bla...

> go help build  # 理论上这个也可以跑
usage: go build [bla bla] bla bla [packages]

> go build hi  # 跑跑看


> hi
Hi, Lirian
```

在边玩边学中，
我们第一个程序就跑成功了。
我们可以试试各种各样的语法，
比如 `defer`, `error` 等。

接下来我们尝试写一个 [Low DB][lowdb] 那样的 json/in-memory database吧。
取个名，就叫…


## Fate DB

TODO: 施工中...


[kevin]: http://www.heyongjian.com/
[kotlin]: http://kotlinlang.org/
[tothegump]: https://github.com/tothegump
[go]: https://golang.org/
[zhihu-kotlin-go]: https://www.zhihu.com/question/60064789
[slant-kotlin-go]: https://www.slant.co/versus/126/1543/~golang_vs_kotlin
[go-name]: /doc/images/go.jpg
[windows-setup]: http://www.liriansu.com/windows-dev-env
[choco]: https://chocolatey.org/
[go-dl]: https://golang.org/dl/
[gogland]: https://www.jetbrains.com/go/
[xy-go]: https://learnxinyminutes.com/docs/go/
[go-tour]: https://tour.golang.org/welcome/
[tabs]: /doc/images/tabs_wtf.jpg
[go-workspace]: https://golang.org/doc/code.html#Workspaces
[classpath]: https://docs.oracle.com/javase/8/docs/technotes/tools/windows/classpath.html
[sys-path]: https://docs.python.org/3/library/sys.html
[lowdb]: https://github.com/typicode/lowdb
