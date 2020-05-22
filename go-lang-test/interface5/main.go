package main

import "fmt"

type Person interface {
	GetID() uint32
}

type Student struct {
	id uint32
}

func (s *Student) GetID() uint32 {
	return s.id
}

func GetPerson() Person {
	var s *Student
	return s
}

func main() {
	if GetPerson() == nil {
		fmt.Println("hi")
	} else {
		fmt.Println("hello")
		GetPerson().GetID()
	}
}
