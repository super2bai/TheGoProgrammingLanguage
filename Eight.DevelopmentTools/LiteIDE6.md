### 8.7 LiteIDE
LiteIDE是国内第一款，也是世界上第一款专门为Go语言开发的集成开发环境(IDE)，目前支持Windows、Linux、iOS三个平台。它的安装和使用都很简单方便，是初学者较好的选择，支持语法高亮、集成构建和代码调试。虽然与专业的IDE相比，LiteIDE需要在很多细节上继续打磨，但仍不失为开发Go语言程序的首选之一。
在部署上，只需要[下载安装包安装](http://code.google.com/p/golangide/downloads/list)，并配置好环境即可。安装过程非常简单，因此不再赘述。
最新发布的版本已经融入了Go 1的全部新特性，尤其是在工程管理上，与Go工具兼容，可以直接根据`GOPATH`导入工程。同时，也支持关键字自动完成。
x11版在IDE的环境配置上，是基于XML的。例如：想把代码关键字由粗体变为正常，需要通过菜单`Option`->`LiteEditor`->`ColorStyle Scheme`菜单来打开和编辑`default.xml`。

原文是:
```xml
<?xml version="1.0" encoding="UTF-8"?>
<style-scheme version="1.0" name="Default">
<!-- Empty scheme, relying entirely on the built-in defaults. -->
</style-scheme>
```
加入关键字定义后，该文件更新为下面的形式：
```xml
<?xml version="1.0" encoding="UTF-8"?>
<style-scheme version="1.0" name="Default">
<!-- Empty scheme, relying entirely on the built-in defaults. -->
<style name="Keyword" foreground="#0000cd" bold="false"/>
</style-scheme>
```
保存该文件即可看到效果。

在根据`GOPATH`完成工程导入之后，`GOPATH`中的工程会显示在IDE左边的"项目"窗口中。这里有个关键的步骤，那就是需要设置"当前选中的工程"，以让IDE环境能够识别需要编译和调试的工程。具体操作方法是双击工程名字或者右击工程目录，然后单击菜单中的"设置当前项目"。完成设置后，当前项目会以蓝色字体显示在"工程"对话框的顶部。

由于LiteIDE目前还不支持版本的自动更新功能，使用LiteIDE的开发者需要自省关注它的[官方主页](http://code.google.com/p/golangide/)以了解最新动态。