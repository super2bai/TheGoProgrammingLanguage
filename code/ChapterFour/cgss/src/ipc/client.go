package ipc

import (
	"encoding/json"
)

/*
IpcClient的关键函数就是Call()了，这个函数回将调用信息封装成一个JSON格式的字符串发送到对应的channel，并等待获取返回
*/
type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{method, params}

	var b []byte
	b, err = json.Marshal(req)

	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <-client.conn // 等待返回值

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1

	return
}

func (client *IpcClient) close() {
	client.conn <- "CLOSE"
}
