package main

import (
	"crypto/tls"
	"log"
)

/**
$ go build echoclient.go
$ go run echoclient.go

output:
2017/02/16 23:56:41 client: connected to:  127.0.0.1:8000
2017/02/16 23:56:41 client: handshake:  true
2017/02/16 23:56:41 client: mutual:  true
2017/02/16 23:56:41 client: wrote "Hello\n" (6 bytes)
2017/02/16 23:56:41 client: read "Hello\n" (6 bytes)
2017/02/16 23:56:41 client: exiting
*/

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, err := tls.Dial("tcp", "127.0.0.1:8000", conf)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()
	log.Println("client: connected to: ", conn.RemoteAddr())

	state := conn.ConnectionState()
	log.Println("client: handshake: ", state.HandshakeComplete)
	log.Println("client: mutual: ", state.NegotiatedProtocolIsMutual)

	message := "Hello\n"
	n, err := conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("client: write: %s", err)
	}
	log.Printf("client: wrote %q (%d bytes)", message, n)

	reply := make([]byte, 256)
	n, err = conn.Read(reply)
	log.Printf("client: read %q (%d bytes)", string(reply[:n]), n)
	log.Print("client: exiting")
}
