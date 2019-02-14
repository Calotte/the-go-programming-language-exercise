package main

import "fmt"

func counter(out chan<-int){
	for x:=0;x<100;x++{
		out<-x
	}
	close(out)
}

func square(out chan<-int,in <-chan int){
	for x:= range in{
		out<-x*x
	}
	close(out)
}

func print(in <-chan int){
	for x:=range in{
		fmt.Println(x)
	}
}
func main(){
	natures :=make(chan int)
	squares :=make(chan int)
	go counter(natures)
	go square(squares,natures)
	print(squares)
}
