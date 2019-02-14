package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"text/scanner"
)
func Marshal(v interface{})([]byte,error){
	var buf bytes.Buffer
	if err:=encode(&buf,reflect.ValueOf(v));err!=nil{
		return nil,err
	}
	return buf.Bytes(),nil
}

func Unmarshal(data []byte,out interface{})(err error){
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {
		if x:=recover();x!=nil{
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

