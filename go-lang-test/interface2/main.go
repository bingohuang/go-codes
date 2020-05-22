package main

import (
	"fmt"
	"os"
)

type (
	Reader interface {
		Read() []byte
	}
	ReaderImpl struct {
		itf *os.File
	}
)

func (r ReaderImpl) Read() []byte {
	return nil
}

func testRead(i interface{}) {
	reader, ok := i.(Reader)
	if !ok || reader == nil {
		fmt.Println("fail")
		return
	}

	reader.Read()
	fmt.Println("Success")
}

func main() {
	//var p *ReaderImpl
	//testRead(p)

	testRead(ReaderImpl{})
}
