### 8.2 gedit
如果在Linux下习惯用gedit，那么可以照此来配置一个"goedit".

gedit是绝大部分Linux发行版自带且默认的文本编辑工具(比如Ubuntu上直接被成为TextEditor)，因此，绝大多数情况下，
只要你在使用Linux，就已经在使用gedit了，不需要单独安装。
接下来我们介绍如何将gedit设置为一个基本的Go语言开发环境。

#### 8.2.1 语法高亮
一般支持自定义语法高亮的文本饿编辑器都是通过一个语法定义文件来设定语法高亮规则的，gedit也是如此。
Go语言社区有人贡献了[可用于gedit的Go语言语法高亮文件](http://go-lang.cat-v.org/text-editors/gedit/go.lang)，
下载后，该文件应该放置到目录`/usr/share/gtksourceview-2.0/language-spece`下。不过如果用的是Ubuntu比较新的版本，
比如v11.01，那么可能会发现gedit默认已经支持Go语言的语法高亮。读者可以在gedit中查看菜单View-HighlightMode-Sources里是否包含名为"Go"的菜单项。