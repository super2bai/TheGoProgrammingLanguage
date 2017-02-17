### 7.3 远程import支持
我们知道，如果要在Go语言中调用包，可以采用如下格式：
```go 
package main
import (
	"fmt"
)
```
其中`fmt`是我们导入的一个本地包。实际上，Go语言不仅允许导入本地包，还支持在语言级别调用远程的包。加入，有一个计算CRC32的包托管于`Github`，那么可以这样写：

```go
package main
import (
	"fmt"
	"github.com/myteam/exp/crc32"
)
```
然后，在执行`go get`之后，就会在`src`目录中看到`github`目录，其中包含`myteam/exp/crc32`目录。在`crc32`中，就是该包的所有源代码。也就是说，go工具会自动帮你获取位于远程的包源码，在随后的编译中，也会在pkg目录中生成对应的`.a`文件

所有魔术般的工作，其实都是go工具在完成。对于Go语言本身来讲,`github.com/myteam/exp/crc32`知识一个与`fmt`无异的本地路径而已。