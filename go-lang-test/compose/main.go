package main

import "fmt"

type (
	ItfA interface{ FuncA() }
	ItfB interface{ FuncB() }
	ItfC interface{ FuncC() }
	ItfD interface{ FuncD() }

	AB      struct{}
	CD      struct{}
	Compose struct {
		ItfA
		CD
	}
)

func (AB) FuncA() {}
func (AB) FuncB() {}
func (CD) FuncC() {}
func (CD) FuncD() {}

func main() {
	var exist bool
	var composeItf ItfA = Compose{AB{}, CD{}}

	_, exist = composeItf.(ItfA)
	if exist {
		fmt.Println("Assert ItfA Success")
	}

	_, exist = composeItf.(ItfB)
	if exist {
		fmt.Println("Assert ItfB Success")
	}

	_, exist = composeItf.(ItfC)
	if exist {
		fmt.Println("Assert ItfC Success")
	}

	_, exist = composeItf.(ItfD)
	if exist {
		fmt.Println("Assert ItfD Success")
	}
}
