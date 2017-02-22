**本小结代码未经测试**

### 7.7 跨平台开发
跨平台**泛指同一段程序可在不同硬件架构和操作系统的设备上运行**，这些设备包括服务器、个人电脑和各种移动设备等。

个人电脑和服务器通常为`80x86`架构，而移动设备则以`ARM`架构为主。

根据CPU的寻址能力，CPU又分为不同的位数：
* 8位
* 16位
* 32位
* 64位
* ...

位数越高，寻址能力就越强。

程序的执行过程其实就是操作系统读取可执行文件的内容，依次调用相应CPU指令的过程。不同的操作系统所支持的可执行文件格式也各不相同，比如Windows支持PE格式，Linux支持ELF。因此，Windows上的可执行文件无法直接直接在Linux上运行。

正因为有了以上这些区别，我们才有了跨平台开发这个话题。

#### 7.7.1 交叉编译
之前生成的`xx.a`之所以放到`linux_amd64目录下，是根据Go编译器决定的。如果当前的编译目标为AMD64架构的64位Linux,对应的安装位置就是windows_386。

鉴于Google对Linux的偏爱，目前Go语言对Linux平台的支持最佳。Mac OS X因为底层也是*nix架构，因此运行Go也没有明显障碍。但Go语言对Windows平台的支持就比较欠缺了，需要通过MingGW间接支持，自然性能不会很好，且开发过程中时常会遇到一些奇怪的问题。

目前而言，Go对64位的x86处理器架构的支持最为成熟(即AMD64),已经可以支持32位的x86和ARM架构，暂时还不支持MIPS。

此外，Go编译器支持交叉编译。如果要在一台安装了64位Linux操作系统的AMD64电脑上执行一段Go代码，就必须用能够生成ELF文件格式的Go编译器进行编译和链接。

Go当前的交叉编译能力如下所示：
+ 在Linux下，可以生成以下目标格式
	- x86 ELF
	- AMD64 ELF
	- ARM ELF
	- X86 PE
	- AMD64 PE
+ 在Windows下，可以生成以下目标格式：
	- x86 PE
	- AMD64 PE

可以通过设置`GOOS`和`GOARCH`两个环境变量来制定交叉编译的目标格式。

下表位当前支持的情况说明，其中darwin对应于MAC OS X。
| GOOS      | GOARCH        |    说明   |
| --------- |:-------------:| -----:   |
| darwin    | 386           | MAC OS X |
| darwin    | amd64         | MAC OS X |
| freebsd   | 386           |          |
| freebsd   | amd64         |          |
| linux     | 386           |          |
| linux     | amd64         |          |
| linux     | arm           |尚未完全支持|
| windows   | amd64         |          |
| windows   | 386           |尚未完全支持|

下面给出在Linux平台下构建Windows32位PE文件的详细步骤：
+ 获取Go源代码
+ 构建本机编译器环境，具体代码如下
	- $ cd  GOROOT路径 /src
	-  ./make.bash
+ 构建跨平台的编译器和链接器，具体代码如下
```bash 
$ cat ~/bin/buildcmd
#!/bin/sh
set -e
for arch in 8 6;do
	for amd in a c g l; do
		go tool dist install -v cmd/$arch$cmd
	done
done
exit 0	
```
+ 构建Windows版本的标准命令工具和库，如下：
```bash
$ cat ~/bin/buildcmd
#!/bin/sh
if [ -z "$1" ]; then
	echo 'GOOS is not specified' 1>&2
	exit 2
else
	export GOOS=$1
	if [ "$GOOS" = "windows" ]; then
		export CGO_ENABLED=0
	fi
fi
shift
if [ -n "$1" ]; then
	export GOARCH=$1
fi
cd $GOROOT/src
go tool dist install -v pkg/runtime
go install -v -a std
```
+ 然后执行下面这段脚本以准备好Windows交叉编译的环境：
```bash
$ ~/bin/buildpkg windows 386
```
+ 在Linux上生成Windows x86的PE文件，具体代码如下
```bash
$ cat hello.go
package main

import "fmt"

func main(){
	fmt.Printf("Hello\n")
}
$ GOOS=windows GOARCH=386 go build -o hello.exe hello.go
```
对于跨平台部署来说，经常会用到交叉编译，因此不用觉得这种功能是多此一举。

#### 7.7.2 Android支持
Android手机一般使用ARM的CPU，并且由于Android使用了Linux内核，属于复合Go语言当前完整支持的架构，因此在Android手机上可以运行Go程序。

首先要定制出能够生成对应目标二进制文件的Go工具链。在编译Go源代码之前，可以作如下设置：

```bash
$ export GOARCH=ARM
$ export GOOS=linux
$ ./all.bash
```

一切顺利的话，会生成5g和5l，其中5g是编译器，5l是链接器。假设生成的目标二进制文件是5.out，接下来使用adb调试器将5.out导入Android虚拟机或者真机中。

```bash
adb push 5.out /data/local/tmp/
adb shell
cd /data/local/tmp
./5.out
```

此时就可以看到执行结果了。鉴于Android开发部在本书的讨论范围，因此这里不做更近一步的解释，有兴趣的人可以参考相关资料。Android和Go都是Google推出的产品，相信两者之间的配合会越来越默契。