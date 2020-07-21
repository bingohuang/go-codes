package main

import (
	"context"
	"log"
)

type customStruct struct {
}

func returnStruct() interface{} {
	var p *customStruct
	if p == nil {
		log.Printf("p is nil")
		return nil // 直接返回nil，而不要把nil赋值给interface{}
	}
	return p
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	p := returnStruct()

	if p == nil {
		log.Printf("is nil")
	} else {
		log.Printf("not nil")
	}
	_ = p

	cancel()
	_ = ctx

}
