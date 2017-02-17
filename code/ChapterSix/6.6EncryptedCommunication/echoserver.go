package main

import (
	"crypto/rand"
	"crypto/tls"
	"io"
	"log"
	"net"
	"time"
)

/*
先运行server
再运行client
生成pem需要修改/etc/ssl/openssl.cnf
对应配置已上传
$ go build echoserver.go
$ go run echoserver.go

output:
2017/02/16 23:56:38 server: listen
2017/02/16 23:56:38 server: conn: waiting
2017/02/16 23:56:41 server: accepted from 127.0.0.1:60218
2017/02/16 23:56:41 server: conn: waiting
2017/02/16 23:56:41 server: conn : waiting
2017/02/16 23:56:41 server: conn: echo "Hello\n"
2017/02/16 23:56:41 server: conn: wrote 6 bytes
2017/02/16 23:56:41 server: conn : waiting
2017/02/16 23:56:41 server: conn: closed
*/
func main() {
	cert, err := tls.LoadX509KeyPair("../../cert/cert.pem", "../../cert/key.pem")
	if err != nil {
		log.Fatalf("server: loadkeys : %s", err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	config.Time = time.Now
	config.Rand = rand.Reader

	service := ":8000"

	listener, err := tls.Listen("tcp", service, config)
	if err != nil {
		log.Fatalf("server: listen: %s", err)
	}
	defer listener.Close()

	log.Print("server: listen")
	for {
		log.Print("server: conn: waiting")
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server: accept: %s", err)
			continue
		}
		log.Printf("server: accepted from %s ", conn.RemoteAddr())

		go handleClient(conn)

	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 512)
	for {
		log.Print("server: conn : waiting")
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("server: conn: read: %s", err)
			}
			break
		}
		//		log.Println("conn Read : " + n)
		log.Printf("server: conn: echo %q\n", string(buf[:n]))
		n, err = conn.Write(buf[:n])
		log.Printf("server: conn: wrote %d bytes", n)

		if err != nil {
			log.Printf("server: write: %s", err)
			break
		}
	}
	log.Println("server: conn: closed")

}
