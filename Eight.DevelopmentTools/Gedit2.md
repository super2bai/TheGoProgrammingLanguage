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

#### 8.2.2 编译环境
在配置构建相关命令之前，需要确认regit是否已经安装了名为`External Tools`的插件。单击`View`->`Preference`菜单项，弹出选项对话框，该对话框的最后一个选项也就是`Plugins`。插件的安装比较简单，只要在插件列表中找到`External Tools`并确认该项已经被勾选即可。

接下来配置几个常用的工程构建命令：
* 构建当前工程(`Go Build`)
* 编译当前打开的Go文件(`Go Compile`)
* 运行单元测试(`Go Test`)
* 安装(`Go Install`)
要添加命令，单击`Tools`->`Manage External Toools...`菜单项，打开管理对话框，然后在该对话框中添加即可。

需要添加的命令主要如下表所示

命令 |  名称  |                                                      脚本内容 						|  保存  |输入|
----|-------|-------------------------------------------------------------------------------------|-------|----|
构建 | Build |#! /bin/bash<br/>echo "Building..."<br/>cd $GEDIT_CURRENT_DOCUMENT_DIR<br/>go build -v |所有文档|无|
运行 |  Run  |#! /bin/bash<br/>echo "Runing..."<br/>cd $GEDIT_CURRENT_DOCUMENT_DIR<br/>go run        |当前文档|当前文档|
测试 | Test  | #! /bin/bash<br/>echo "Testing..."<br/>cd $GEDIT_CURRENT_DOCUMENT_DIR<br/>go test     |所有文档|无|
安装 |Install|#! /bin/bash<br/>echo "Installing..."<br/>cd $GEDIT_CURRENT_DOCUMENT_DIR<br/>go Install|所有文档|无|

可以很容易看出来，每个命令的内容其实就是一个`shell`脚本，可以根据自己的需求进行任意的修改和扩充。
添加完命令后，可以在`Tool`->`External Tools`菜单中看到刚刚添加的所有命令。每次淡季菜单项来构建也不是非常方便，因此建议在添加命令时顺便设置一下快捷方式。