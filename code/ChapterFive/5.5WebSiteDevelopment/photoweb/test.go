package main

//import (
//	"fmt"
//	"html/template"
//	"io/ioutil"
//	"log"
//	"net/http"
//	//	"os"
//	"strings"
//)

//const (
//	ListDir      = 0x0001
//	UPLOAD_DIR   = "./uploads"
//	TEMPLATE_DIR = "./views"
//)

//var templates = make(map[string]*template.Template)

//func init() {

//	fileInfoArr, _ := ioutil.ReadDir(TEMPLATE_DIR)
//	fmt.Println("111")
//	fmt.Println(fileInfoArr[0])

//	fmt.Println(fileInfoArr[1])
//	//check(err)
//	//var templateName, templatePath string

//	for i, v := range fileInfoArr {
//		fmt.Println(i)

//		fmt.Println("444")
//		fmt.Println(v.Name())

//		//		log.Fatal(fileInfo)
//		//		templateName = fileInfo.Name()
//		//		log.Fatal("templateName=" + templateName)
//		//		ext := path.Ext(templateName)
//		//		log.Fatal("ext=" + ext)
//		//		if ext != ".html" {

//		//			log.Fatal("wrong : != html" + ext)
//		//			continue
//		//		} else {
//		//			log.Fatal("ext=" + ext)
//		//		}

//		//		log.Fatal(TEMPLATE_DIR)
//		//		templatePath = TEMPLATE_DIR + "/" + templateName

//		//		log.Fatal("Loading template:", templatePath)
//		//		//template.ParseFiles()会读取指定模版的内容并返回一个*template.Template值
//		//		t := template.Must(template.ParseFiles(templatePath))
//		//		templates[templateName] = t
//		//		log.Fatal(templates)
//	}

//}

//func sayhelloName(w http.ResponseWriter, r *http.Request) {

//	// 解析参数, 默认是不会解析的
//	r.ParseForm()

//	// 这些信息是输出到服务器端的打印信息
//	fmt.Println("request map:", r.Form)
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])

//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ";"))
//	}

//	// 这个写入到w的信息是输出到客户端的
//	fmt.Fprintf(w, "Hello gerryyang!\n")
//}

//func main() {

//	mux := http.NewServeMux()
//	// 设置访问的路由
//	mux.HandleFunc("/", sayhelloName)

//	// 设置监听的端口
//	err := http.ListenAndServe(":9090", mux)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", mux)
//	}
//}
