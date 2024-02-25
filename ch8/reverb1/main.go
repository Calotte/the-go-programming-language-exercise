package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	if _, err := fmt.Fprintln(c, "\t", strings.ToUpper(shout)); err != nil {
		fmt.Println(err)
	}
	time.Sleep(delay)
	if _, err := fmt.Fprintln(c, "\t", shout); err != nil {
		fmt.Printf("%s, err: %v", shout, err)
	}
	time.Sleep(delay)
	if _, err := fmt.Fprintln(c, "\t", strings.ToLower(shout)); err != nil {
		fmt.Printf("%s, err: %v", strings.ToLower(shout), err)
	}
}
