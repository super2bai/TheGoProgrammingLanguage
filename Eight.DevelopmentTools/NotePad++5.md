### 8.5 Notepad++
Notepad++是Windows平台上最受欢迎的第三方文本编辑工具之一。相比另一个名头更响的工具UltraEdit，Notepad++的最大优势在于免费。可以将Notepad++简单配置以下，使其支持Go语言的语法高亮，并让开发者尽可能在不离开Notepad++的情况下即可进行开发Go语言程序的所需动作。
可以先从Notepad++的[官方网站](http://notepad-plus-plus.org/)下载该工具并安装，之后按下面的步骤配置。

#### 8.5.1 语法高亮
在Go语言的安装目录下，已经自带了针对Notepad++的语法高亮配置文件。可以在`/usr/local/go/misc/notepadplus`目录下找到这些配置文件。只需安装对应的READMD文档进行以下几个步骤的操作：
* 将`userDefineLang.xml`的内容合并到Notepad++配置目录下的`userDefineLang.xml`文件。如果安装目录下不存在这个文件，则直接复制该文件即可。Notepad++的配置目录通常位于`%HOME%\AppData\Roaming\Notepad++`。
* 将`go.xml`复制到安装目录的`plugins\APIs`目录下
* 重启Notepad++

`%HOME`是指HOME目录，如果不知道在哪里，在命令行中执行`echo %HOME`即可看到。

#### 8.5.2 编译环境
推荐使用Notepad++用户在安装另外两个Notepad++的插件：
* NppExec(支持自定义命令)
* Explorer(避免在Notepad++和资源管理器之间频繁切换)

在Notepad++中即可完成目录结构和文件的操作。Notepad++的插件安装非常简单，只需在插件对话框中找到这两个插件并选中即可。

**1.配置NppExec插件**

在安装好NppExec插件后，可以通过以下几个简单的步骤将NppExec配置为适合用于构建Go程序的环境：
* 点击菜单`Plugins`->`NppExec`，打开对话框，勾选:`Show Console Dialog`、`No internal messages`、`Save all files on execute`和`Follow $(CURRENT_DIRECTORE)`这4歌选项。
* 在Exec对话框中分别键入`go build`、`go clean & go install`和`go test`，并保存为`build`、`install`和`test`脚本。此时已经可以测试Go工程的build是否能够正常进行，**以下步骤为可选操作**
* 在`Advanced Options`中添加3条正对以上脚本的命令，分别为`Build current project`、`Install current project`和`Test current project`
* 点击菜单`Settings`->`Shortcut Mapper`->`Plugin Commands`为这3条命令分配快捷键，**如果快捷键冲突，则需要先清除这些快捷键的默认配置**
* 通过`Console Output Filters`对话框的`Highlight`选项卡美化程序运行结果消息。添加一下内容高亮规则:
	* 筛选规则：`*PASS*`;显示格式：蓝色粗体(`*PASS*`为填入到mask框中的内容，蓝色和粗体则通过在Blue中填入0xff和勾选B来完成)
	* 筛选规则：`%FILE%`;显示格式：红色下划线（这一条很有价值，因为可以让你双击消息定位到相应的代码行。可惜还不支持正则表达式，否则就真正强大了）
	* 筛选规则：`gotest:parse error:%FILE:%LINE%:*`;显示格式：红色
	* 筛选规则：`*FAIL*`；显示格式：红色粗体

**2.配置Explorer插件**

点击菜单`Plugins`->`Explorer`->`Explorer`打开目录树窗格，并按自己的喜好配置Explorer的显示选项即可。因为Go语言已经抛弃了专门的工程文件，因此管理工程就是管理目录结构，不再需要复杂的配置工具。Explorer插件就足以满足需求。


