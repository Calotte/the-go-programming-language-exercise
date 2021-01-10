package bytecounter

import "testing"

func TestByteCounter_Write(t *testing.T) {
	var c ByteCounter
	n, err := c.Write([]byte("hello"))
	want := 5
	if n != want {
		t.Fatalf("Write 'hello' len %d, %v, want match for %d", n, err, want)
	}
}

func TestWordCounter_Write(t *testing.T) {
	var w WordCounter
	s := "Hello World"
	n, err := w.Write([]byte(s))
	want := 2
	if n != want {
		t.Fatalf("Write '%s' len %d, %v, want match for %d", s, n, err, want)
	}
}

func TestLineCounter_Write(t *testing.T) {
	var l LineCounter
	s := "Hello\nWorld\nCalotte"
	n, err := l.Write([]byte(s))
	want := 3
	if n != want {
		t.Fatalf("Write '%s' len %d, %v, want match for %d", s, n, err, want)
	}
}
