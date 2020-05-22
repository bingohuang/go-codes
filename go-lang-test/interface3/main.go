package main

import (
	"fmt"
)

type (
	Writer interface {
		Write([]byte) []byte
	}
	WriterImpl struct {
		Writer
	}
)

func (w WriterImpl) Write(bytes []byte) {
	writer := w.Writer
	writer.Write(bytes)
	return
}

func testWrite(i interface{}) {
	writer, ok := i.(Writer)
	if !ok || writer == nil {
		fmt.Println("fail")
		return
	}

	writer.Write(nil)
	fmt.Println("Success")
}

func main() {
	var p *WriterImpl
	testWrite(p)

	//testWrite(WriterImpl{})
}
