利用第二章所学的内容来实现一个完整的排序算法的比较程序。
从命令行指定输入的数据文件和输出的数据文件，并指定对应的排序算法。
用法：
sorter -i <int> -o <out> -a <qsort|bubblesort>

示例:
```bash
$ ./sorter -I in.dat -o out.dat -a qsort
The sorting process costs 10us to complete
```
------

该函数分为两类：
* 主程序：sorter.go
* 排序算法函数：qsort.go(实现快速排序)；bubblesort.go(实现冒泡排序)

主程序功能：
* 获取并解析命令行输入
* 从对应文件中读取输入数据
* 调用对应的排序函数
* 将排序的结果输出到对应的文件中
* 打印排序所话费时间的信息

------

输入文件是一个纯文本文件，每一行是一个需要被排序的数字。需要逐行从这个文件中读取内容，并解析为int类型的数据，再添加到一个int类型的数组切片中。

------
###自动编译并运行
>编辑sorter/build.sh,设置GOPATH为当前路径

```bash
./build.sh
./run.sh
```

### 手动编译

>1.需添加 sorter 当前路径到GOPATH中(export GOPATH=之前设置的GOPATH内容:sorter当前路径)，然后在terminal里执行:

```bash
go build algorithm/qsort
go build algorithm/bubblesort
go test algorithm/qsort
go test algorithm/bubblesort
go install algorithm/qsort
go install algorithm/bubblesort
```

>2.然后进入到bin文件夹,执行(如果在根路径下直接执行下面两行命令，生成文件会在跟路径下，暂不得解)

```bash 
go build sorter
go install sorter
```

>3.手动创建unsorted.dat文件，并逐行写进大小顺序错乱的数字（一个数字一行）

>4.执行

```bash
./sorter -i unsorted.dat -o sorted.dat -a qsort
./sorter -i unsorted.dat -o sorted.dat -a bubblesort
```




