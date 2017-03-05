package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"fmt"
)

/*
Go的第三方库很丰富，无论是对于关系型数据库和驱动还是非关系型的键值存储系统的介入
都有着很好的支持，而且还有丰富的Go语言Web开发框架以及用于Web开发的相关工具包
可以访问 http://godashboard.appspot.com/project(没有请求成功，嘤嘤嘤）
了解更多第三方库的详细信息
*/
const (
	ListDir      = 0x0001
	UPLOAD_DIR   = "./uploads"
	TEMPLATE_DIR = "./views"
)

/**
全局变量
存放所有模版内容
map类型的复合结构
map的key是字符串类型，即模版名字
值value是*template.Template类型
*/
var templates = make(map[string]*template.Template)

/**
7.模版缓存
即一次性预加载模版
可以在程序初始化的时候，将所有模版一次性加载到程序中
Go的包加载机制允许我们在init()函数中作这样的事情
init()会在main()函数之前执行

template.Must确保了模版不能解析成功时，一定会出发错误处理流程。
之所以这么作，是因为倘若模版不能成功加载，程序能做的唯一有一一的事情就是退出
*/

func init() {
	var fileInfoArr = []os.FileInfo{}
	fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)

	check(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		fmt.Println("templateName=" + templateName)
		ext := path.Ext(templateName)
		fmt.Println("ext=" + ext)
		if ext != ".html" {

			fmt.Println("wrong : != html" + ext)
			continue
		} else {
			fmt.Println("ext=" + ext)
		}

		fmt.Println(TEMPLATE_DIR)
		templatePath = TEMPLATE_DIR + "/" + templateName

		fmt.Println("Loading template:", templatePath)
		//template.ParseFiles()会读取指定模版的内容并返回一个*template.Template值
		t := template.Must(template.ParseFiles(templatePath))
		templates[templateName] = t
		fmt.Println(templates)
	}

}

/**
统一进行错误处理
*/
func check(err error) {
	if err != nil {
		panic(err)
	}
}

/**
渲染网页模版
使用Go标准库提供的html/templete包，可以让我们将HTML从业务逻辑中抽离出来形成独立的模版文件
这样业务逻辑程序只负责处理业务逻辑部分和提供模版需要的数据
模版文件负责数据要表现的具体形式
模版解析器将这些数据以定义好的模版规则结合模版文件进行渲染
最终将渲染后的结果一并输出，形成一个完整的网页
*/
func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) {
	tmpl += ".html"
	//Execute()会根据模版愈发来执行模版的渲染，并将渲染后的结果作为HTTP的返回数据输出
	err := templates[tmpl].Execute(w, locals)
	check(err)
}

/**
5.处理不存在的图片访问
理论上，只要是uploads/目录下有的图片，都能够访问到，但还是假设有意外情况
比如网页中传入的图片在uploads/没有对应的文件
不管是给出友好的提示错误还是返回404页面，都应该对这种情况作相应处理
*/
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	/**
	1.结合main()和uploadHandler方法，针对HTTP GET方式请求/upload路径
	程序将会往http.ResponseWriter类型的示例对象w中写入一段HTML文本
	即输出一个HTTP上传表单。
	如果我们使用浏览器访问这个地址，那么网页上将会一个可以上传文件的表单
	*/

	if r.Method == "GET" {
		renderHtml(w, "upload", nil)
	}
	/**
	2.如果是客户端发起的HTTP POST请求，那么首先从表单提交过来的字段寻找名为image的文件域并对其接值
	调用r.FormFile()方法返回3个值，各个值的类型分别是multipart.File、*multipart.FileHeader和error
	如果上传的图片接收不成功，则会返回一个HTTP服务端的内部错误给客户端
	如果上传的图片接收成功，则将该图片的内容复制到一个临时文件里
	如果临时文件创建失败，或者图片副本保存失败，都将出发服务端内部错误
	如果临时文件创建成功并且图片副本保存成功，即表示图片上传成功，就跳转到查看图片页面
	此外，还定义了两个defer语句，无论图片上传成功还是失败
	当uploadHandler()方法执行结束时，都会先关闭临时文件句柄，继而关闭图片上传到服务器文件流的句柄
	当图片上传成功后，即可在网页上查看这张图片，顺便确认图片是否真正上传到了服务器端
	*/
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		check(err)
		defer t.Close()
		_, err = io.Copy(t, f)
		check(err)
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

/**
3.在网页上显示图片
要在网页中显示图片，必须有一个可以访问到该图片的网址
在前面的示例代码中，图片上传成功后会跳转到/view?id=<ImageId>这样的网址
因此要对/view路径的访问映射到某个具体的业务逻辑处理方法

首先从客户端请求中对参数进行接值
r.FormValue("id")即可得到客户端传递的图片唯一ID
结合之前保存图片用的目录进行组装，即可得到文件在服务器上的存放路径

接着，调用http.ServeFile()方法将该路径下的文件从磁盘中读取并作为服务端的返回信息输出给客户端
同时，也将HTTP响应头输出格式预设为image类型
这是一种比较简单的示意写法，实际上应该严谨些，准确解析出文件的MimeType并将其作为Content-Type进行输出
具体可参考Go语言标准库中的http.DetectContentType()方法和mimn包提供的相关方法
*/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if ok := isExists(imagePath); !ok {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

/**
6.列出所有已上传图片
在网页上列出/uploads下存放的所有文件
这里只需列出可供访问的文件名称即可

遍历/uploads目录，得到所有文件并赋值到fileInfoArr变量里
fileInfoArr是一个数组，其中的每一个元素都是一个文件对象
然后遍历fileInfoArr数组并从中得到图片的名称
用于在后续的HTML片段中显示文件名和传入的参数内容
*/
func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	check(err)
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	renderHtml(w, "list", locals)
}

/**
巧妙闭包避免程序运行时出错崩溃
Go支持闭包。闭包可以是一个函数里返回的另一个匿名函数，
该匿名函数包含了定义在它外面的值
使用闭包，可以让网站的业务逻辑处理程序更安全的运行

接收一个业务逻辑处理函数作为参数，同时调用这个业务逻辑处理函数

HandlerFunc有两个参数：httpResponseWriter和*htt.Request
函数规格同photoweb的业务逻辑处理函数完全一致
事实上，正式要把业务逻辑处理函数作为参数传入到safeHandler()方法中
这样任何一个错误处理流程向上回溯的时候
都能对其进行拦截处理，从而也能避免程序停止运行

利用defer关键字和recover()方法中介panic的肆行
该业务逻辑函数执行完毕后，safeHandler()中defer指定的匿名函数会执行
倘若业务逻辑处理函数里面引发了panic，则调用recover对其进行检测
若为一般性错误，则输出HTTP 50X出错信息并记录日志，而程序将继续良好运行
*/
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e, ok := recover().(error); ok {
				http.Error(w, e.Error(), http.StatusInternalServerError)

				// 或者输出自定义的 50x 错误页面
				// w.WriteHeader(http.StatusInternalServerError)
				// renderHtml(w, "error", e.Error())

				// logging
				log.Println("WARN: panic fired in %v .panic - %v", fn, e)

			}
		}()
		fn(w, r)
	}
}

/**
业务逻辑都是动态请求，但若是针对静态资源(css、js)是没有业务逻辑处理的，值需要提供静态输出
net/http包提供的ServeFile()函数可以将服务器的一个文件内容读写到http.ResponseWriter
并返回给请求来源的*http.Request客户端
用闭包技巧结合ServeFile()方法，就可以实现业务逻辑的动态请求和静态资源的完全分离

如果使用外部Web服务器(比如Nginx等)，就没有必要使用Go编写的静态文件服务了
*/
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			fi, err := os.Stat(file)
			if err != nil || fi.IsDir() {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

/**
4.完成viewHandler()的业务逻辑后
将该方法注册到程序的mian()方法中，与/view路径访问形成映射关联
这样当客户端(浏览器)访问/view路径并传递id参数时
即可直接以HTTP形式看到图片的内容
在网页上，将会呈现一种可视化的图片
*/
func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)
	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/list", safeHandler(listHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	err := http.ListenAndServe(":9080", mux)
	if err != nil {
		fmt.Println("ListenAndServe: ", err.Error())
	}
}
