package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		_, _ = io.Copy(os.Stdout, conn)
		if tcpConn, ok := conn.(*net.TCPConn); ok {
			_ = tcpConn.CloseRead()
			fmt.Println("Read conn closed")
		}
		log.Println("done")
		done <- struct{}{}
	}()
	str := "Hello\n"
	r := strings.NewReader(str)
	mustCopy(conn, r)
	if tcpConn, ok := conn.(*net.TCPConn); ok {
		_ = tcpConn.CloseWrite()
		fmt.Println("Write conn closed")
	} else {
		log.Println("convert failed")
		_ = conn.Close()
	}
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	n, err := io.Copy(dst, src)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("copy %d data\n", n)
}
