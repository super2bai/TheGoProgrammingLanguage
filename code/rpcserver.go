package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

/**
注册服务对象并开启该RPC服务
*/
//	arith := new(Arith)
//	rpc.Register(arith)
//	rpc.HandleHTTP()
//	l, e := net.Listen("tcp", ":1234")
//	if e != nil {
//		log.Fatal("listen error:", e)
//	}
//	go http.Serve(l, nil)

/**
此时，RPC服务端注册了一个Arith类型的对象及其公开方法Arith.Multiply()和Arith.Divide()供RPC客户端调用。
RPC在调用服务端提供的方法之前，必须先与RPC服务端建立连接
如下列代码所示
*/
//	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
//	if err != nil {
//		log.Fatal("dialing:", err)
//	}

/**
在建立连接之后，RPC客户端就可以调用服务端提供的方法
下面是同步调用程序顺序执行的方式
*/

//	args := &server.Args{7, 8}
//	var reply int
//	err = client.Call("Arith.Multiply", args, &reply)
//	if err != nil {
//		log.Fatal("arith error:", err)
//	}
//	fmt.Printf("Arith:%d*%d=%d", args.A, args.B, reply)

/**
下面是异步方式进行调用
*/

//	quotient := new(Quotient)
//	divCall := client.Go("Arith.Divide", args, &quotient, nil)
//	replyCall := <-divCall.Done
