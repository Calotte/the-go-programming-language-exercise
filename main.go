package main

import (
	"fmt"
	"the-go-programming-language-exercise/ch5/toposort"
)

func main() {
<<<<<<< HEAD
	for i, course := range toposort.TopoSort(toposort.Prereqs) {//aa
		fmt.Printf("%d:\t%s\n", i, course)//comment
=======
	for i, course := range toposort.TopoSort(toposort.Prereqs) {//
		fmt.Printf("%d:\t%s\n", i, course)
>>>>>>> f71d1053ce3b8b6ea3798ada828442699a3e0908
	}
}
