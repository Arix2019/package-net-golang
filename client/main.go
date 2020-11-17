// client
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	//conn.Write([]byte(fmt.Sprint("Olar!!")))
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)

}
