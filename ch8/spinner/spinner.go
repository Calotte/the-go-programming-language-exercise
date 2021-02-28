package main

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration) {
	for {
		for _, ch := range `_\|/` {
			fmt.Printf("\r%c", ch)
			time.Sleep(delay)
		}
	}
}

func Fib(a int) int {
	if a < 2 {
		return a
	}
	return Fib(a-2) + Fib(a-1)
}

func main() {
	go Spinner(100 * time.Microsecond)
	const n = 45
	fibN := Fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
