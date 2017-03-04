### 8.4 Eclipse
Eclipse是一个成熟的IDE平台，目前已经可以支持大部分流行的语言，包括Java、C++等。Goclipse是Eclipse的插件，用于支持Golang。从整体上看，安装Goclipse插件的Eclipse是目前最优秀的Go语言开发环境，可以实现语法高亮、成员联想、断点调试，基本上满足了所有的需求。
接下来一步步配置Eclipse，将其配置为适合Go语言开发的环境。
* 安装JDK 1.6及以上版本。在目前流行的Linux发行版中，都会预装OpenJDK,虽然功能与Oracle的官方JDK基本一致，但是建议先删除OpenJDK,具体操作方法如下（此操作在安装官方JDK之前进行）：
``` bash
rpm -qa | grep java
rpm -e --nodeps java-1.6.0-openjdk-1.6.0.0-1.7.b09.e15
```
在Windows平台上，不需要此步骤。只简单的安装官方JDK即可。
* 安装Eclipse3.6及以上版本。无论是在Linux还是Windows平台上，一般只需要解压到一个指定的位置即可，不需要特别的配置
* 安装Go编译器，并配置好`GOROOT`、`GOBIN`等环境变量
* 打开Eclipse，点击`Help`->`Install New Software`菜单，打开安装软件对话框()Eclipse版本不同，菜单位置和名称可能也略有差异，但是功能没有区别)。
* 在打开的安装软件对话框的`Work with`文本框中，输入以下URL:https://goclipse.googlecode.com/svn/trunk/goclipse-update-site ，并按回车
* 根据Eclipse的提示，单击`Next`按钮即可。此过程需要一定时间的等待， 如果中途出错，可以多次重试，直到成功为止。

在整个过程中，会因为网络不稳定或者操作系统版本的缘故，下载缓慢或者失败，只要重复上述步骤即可。
* 重启Eclipse，并通过菜单项`Window`->`Preferences`->`Go`打开Go语言的配置选项框，配置Go编译器的路径和GDB的路径

配置完成后，查看执行效果（编辑界面及调试界面）。
因为Go编译器生成的调试信息为DWARFv3格式，因此需要确认所安装的GDB版本必须高于V7.1。