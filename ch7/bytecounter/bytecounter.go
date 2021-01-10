package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func counter(s string, split bufio.SplitFunc) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(split)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, nil
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	s := string(p)
	count, err := counter(s, bufio.ScanWords)
	*w += WordCounter(count)
	return count, err
}

func (l *LineCounter) Write(p []byte) (int, error) {
	s := string(p)
	count, err := counter(s, bufio.ScanLines)
	*l += LineCounter(count)
	return count, err
}

func main() {
	var w LineCounter
	s := "a\nb\nworld"
	_, err := fmt.Fprintf(&w, "hello %s", s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(w)
}
