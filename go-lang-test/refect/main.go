package main

import (
	"fmt"
	"reflect"
)

func main() {
	//x := 2
	//d := reflect.ValueOf(&x).Elem()
	//d.Set(reflect.ValueOf(4))
	//fmt.Println(d)

	x := 2
	d := reflect.ValueOf(&x).Elem()
	//d.Set(4)
	d.SetFloat(4)
	fmt.Println(d)
}
