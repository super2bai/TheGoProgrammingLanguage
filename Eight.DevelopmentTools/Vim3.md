### 8.3 Vim
Go语言安装包中已经包含了对Vim的环境支持。要将Vim配置为适合作为Go语言的开发环境，我们只需要按`$GOROOT/misc/vim`中的说明文档做以下设置即可。
创建一个`shell`脚本`govim.sh`，内容如下：
```bash
$ mkdir -p $HOME/.vim/ftdetect
$ mkdir -p $HOME/.vim/syntax
$ mkdir -p $HOME/.vim/autoload/go
$ ln -s $GOTOOT/misc/vim/ftdetect/gofiletype.vim $HOME/.vim/ftdetect
$ ln -s $GOTOOT/misc/vim/syntax/go.vim $HOME/.vim/syntax
$ ln -s $GOTOOT/misc/vim/autoload/go/complete.vim $HOME/.vim/autoload/go
$ echo "syntax on" >> $HOME/.vimrc
```
在执行该脚本之前，先确认`GOROOT`环境变量是否正确设置并已经起作用，具体代码如下：
```bash
$ echo $GOROOT
/usr/local/go
```
如果上面这个命令的输出为空，则表示`GOROOT`尚未正确设置，请保证`GOROOT`环境变量正确设置后再执行上面的`govim.sh`脚本。
现在可以执行这个脚本了，该脚本只需要执行一次。执行成功的话，在`$HOME`目录下将会创建一个`.vim`目录。之后再用Vim打开一个`go`文件，应该就可以看到针对Go语言的语法高亮效果了。
Vim还可以配合`gocode`支持输入提示功能。接下来简单配置以下。
首先获取`gocode`:
```bash 
$ go get -u github.com/nsf/gocode
```
这个命令会下载`gocode`相应内容到Go的安装目录(比如`usr/local/go`)，因此需要保证有目录的写权限。然后开始配置`gocode`
```bash
$ cd /usr/local/go/src/pkg/github.com/nsf/gocode/
$ cd vim
$ ./update.bash
```
配置就是这么简单。现在使用以下Vim的语法提示效果。用Vim创建一个新的Go文件(比如命名为`auto.go`),输入以下内容：
```go
package main
import "fmt"
func main() {
	fmt.Print
}
```
请将光标停在`fmt.Print`后面，然后按组合键`Ctrl+X+O`，会看到`fmt`包里的3个以`Print`开头的全局函数都被列了出来:`Print`、`Printf`和`Println`。之后就可以用上下方向键选取，按回车键即可完成输入，非常方便。
`gocode`其实是一个独立的提供输入提示的服务器程序，并非专为Vim打造。比如Emacs也可以很容易添加给予`gocode`的Go语言输入提示功能。可以查看`gocode`的Github主页上的提示。