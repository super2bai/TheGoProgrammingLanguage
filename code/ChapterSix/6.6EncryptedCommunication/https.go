package main

import (
	"fmt"
	"net/http"
)

const SERVER_PORT = 8090
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "hello"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	//可以看到，使用http.ListenAndServeTLS()这个方法，这表明它是执行在TLS层上的HTTP协议。
	//运行该服务器后，我们可以在浏览器中访问localhost:8080并查看访问效果
	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "../../cert/cert.pem", "../../cert/key.pem", nil)
	//如果并不需要支持HTTPS，只需要把该方法替换为
	//http.ListenAndServe(fmt.Sprintf(":%d", SERVER_PORT), nil)

}
