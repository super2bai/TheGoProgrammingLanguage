package ipc

import (
	"encoding/json"
	"fmt"
)

/*
可以看出，用server借口确定了之后所要实现的业务服务器的统一接口。
因为IPC框架已经解决了“网络层”的通信问题(这里的网络层用channel代替了)
业务服务器的使用者需要定义支持的指令，然后进行实现即可。
*/
type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handler(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(c chan string) {
		for {
			request := <-c
			//关闭该链接
			if request == "CLOSE" {
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
				return
			}

			resp := server.Handler(req.Method, req.Params)
			b, err := json.Marshal(resp)
			c <- string(b) //返回结果
		}
		fmt.Println("Session closed.")
	}(session)

	fmt.Println("A new session has been created successfully.")
	return session
}
