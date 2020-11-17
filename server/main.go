// SERVER
// Implementação de uma requisição HTTP, utilizando o pacote 'net'.
package main

import (
	"fmt"
	"net"
)

var resp = "O servidor diz:\n\tOlar!!!\n\t  E\n\tAte' logo!\r\n"

type connHeader struct {
	contLen  string
	http     string
	contType string
}

var ch = connHeader{
	contLen:  "Content-Length: ",
	http:     "HTTP/1.1 200 OK\r\n",
	contType: "Content-Type: text/plain\r\n\r\n",
}

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Aguardando conexões...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go response(conn)
	}
}

func response(conn net.Conn) {
	conn.Write([]byte(ch.http))
	conn.Write([]byte(fmt.Sprint(ch.contLen, len(resp), "\r\n")))
	conn.Write([]byte(ch.contType))
	conn.Write([]byte(resp))
	conn.Close()
}
