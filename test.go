package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup

func main() {
	c:=make(chan int)
	var a string
	wg.Add(1)
	go func() {
		defer wg.Done()
		a="hello world"
		<-c
		fmt.Println("hello go func")
		time.Sleep(2*time.Second)
	}()
	c<-0
	fmt.Println(a)
	wg.Wait()
}