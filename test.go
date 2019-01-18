package main

import (
	"fmt"
	"the-go-programming-language-exercise/ch5/toposort"
)

func main() {
	for i, course := range toposort.TopoSort(toposort.Prereqs) {//
		fmt.Printf("%d:\t%s\n", i, course)//
	}
}
