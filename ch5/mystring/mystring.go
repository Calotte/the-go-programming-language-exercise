package mystring

import "bytes"

func Join(sep string,a ...string)string{
	var buffer bytes.Buffer
	for i,s := range a{
		if i>0{
			buffer.WriteString(sep)
		}
		buffer.WriteString(s)
	}
	return buffer.String()
}
